// rs-probe: the HEALTHCHECK helper of ReadyStack's slim runtime images, which carry no curl.
//
//	rs-probe URL         exit 0 when GET URL answers 2xx or 3xx (redirects not followed), 1 otherwise
//	rs-probe -body URL   the same, and print the response body
//	rs-probe -code URL   print the HTTP status code; exit 0 on any response, 1 (printing 000) on none
//	  -k                 do not verify the server's TLS certificate (curl -k)
//	  -cert F -key F     present this client certificate and key (curl --cert F --key F)
//	  -X METHOD          send METHOD instead of GET (curl -X)
//	  -T FILE            send FILE as the request body, with PUT unless -X names another method (curl -T)
//	  -sigv4 aws:amz:REGION:SERVICE -user ACCESS_KEY:SECRET_KEY
//	                     sign the request with AWS Signature Version 4 (curl --aws-sigv4 ... --user ...): the
//	                     host, x-amz-content-sha256 (the body's SHA-256) and x-amz-date headers are signed
//	rs-probe -tcp HOST:PORT
//	                     send standard input to HOST:PORT, print what it answers within 2 s, exit 0 once connected
//	                     (curl telnet://HOST:PORT, as a protocol check of a server that does not speak HTTP)
//
// The verdicts of `curl -fsS -o /dev/null URL`, `curl -fsS URL` and `curl -s -o /dev/null -w '%{http_code}' URL`.
// Standard library only, built with CGO_ENABLED=0 by the kit's pinned Go toolchain: no library, no package.
package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"time"
)

func usage() {
	fmt.Fprintln(os.Stderr, "usage: rs-probe [-k] [-cert FILE -key FILE] [-X METHOD] [-T FILE]\n"+
		"                [-sigv4 aws:amz:REGION:SERVICE -user KEY:SECRET] [-body|-code] http(s)://host:port/path\n"+
		"       rs-probe -tcp host:port < request")
	os.Exit(2)
}

func main() {
	var mode, target, cert, key, tcpAddr, method, upload, sigv4, user string
	insecure := false
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		a := args[i]
		switch {
		case a == "-body" || a == "-code":
			mode = a
		case a == "-k":
			insecure = true
		case (a == "-cert" || a == "-key" || a == "-tcp" || a == "-X" || a == "-T" || a == "-sigv4" || a == "-user") && i+1 < len(args):
			i++
			switch a {
			case "-cert":
				cert = args[i]
			case "-key":
				key = args[i]
			case "-X":
				method = args[i]
			case "-T":
				upload = args[i]
			case "-sigv4":
				sigv4 = args[i]
			case "-user":
				user = args[i]
			default:
				tcpAddr = args[i]
			}
		case target == "" && a != "" && a[0] != '-':
			target = a
		default:
			usage()
		}
	}
	if tcpAddr != "" {
		if target != "" || mode != "" || method != "" || upload != "" || sigv4 != "" || user != "" {
			usage()
		}
		tcpProbe(tcpAddr)
		return
	}
	if target == "" || (cert == "") != (key == "") || (sigv4 == "") != (user == "") {
		usage()
	}
	var body []byte
	if upload != "" {
		b, err := os.ReadFile(upload)
		if err != nil {
			fmt.Fprintln(os.Stderr, "rs-probe:", err)
			os.Exit(1)
		}
		body = b
		if method == "" {
			method = "PUT"
		}
	}
	if method == "" {
		method = "GET"
	}
	cfg := &tls.Config{InsecureSkipVerify: insecure}
	if cert != "" {
		pair, err := tls.LoadX509KeyPair(cert, key)
		if err != nil {
			fmt.Fprintln(os.Stderr, "rs-probe:", err)
			os.Exit(1)
		}
		cfg.Certificates = []tls.Certificate{pair}
	}
	c := &http.Client{
		Timeout:       10 * time.Second,
		Transport:     &http.Transport{Proxy: http.ProxyFromEnvironment, TLSClientConfig: cfg},
		CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse },
	}
	req, err := http.NewRequest(method, target, bytes.NewReader(body))
	if err == nil && sigv4 != "" {
		err = signV4(req, body, sigv4, user, time.Now().UTC())
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "rs-probe:", err)
		os.Exit(2)
	}
	r, err := c.Do(req)
	if err != nil {
		if mode == "-code" {
			fmt.Print("000")
		}
		fmt.Fprintln(os.Stderr, "rs-probe:", err)
		os.Exit(1)
	}
	if mode == "-body" {
		io.Copy(os.Stdout, r.Body)
	}
	r.Body.Close()
	if mode == "-code" {
		fmt.Print(r.StatusCode)
		os.Exit(0)
	}
	if r.StatusCode >= 200 && r.StatusCode < 400 {
		os.Exit(0)
	}
	fmt.Fprintln(os.Stderr, "rs-probe: HTTP", r.StatusCode)
	os.Exit(1)
}

func hmacSHA256(k []byte, s string) []byte {
	m := hmac.New(sha256.New, k)
	m.Write([]byte(s))
	return m.Sum(nil)
}

// awsEscape: RFC 3986 percent-encoding of everything but the unreserved characters, as SigV4's canonical query wants.
func awsEscape(s string) string {
	return strings.ReplaceAll(url.QueryEscape(s), "+", "%20")
}

// signV4 signs req with AWS Signature Version 4, the way curl --aws-sigv4 aws:amz:REGION:SERVICE --user KEY:SECRET
// does for S3: SignedHeaders host;x-amz-content-sha256;x-amz-date.
func signV4(req *http.Request, body []byte, provider, user string, now time.Time) error {
	p := strings.Split(provider, ":")
	if len(p) != 4 || p[0] != "aws" || p[1] != "amz" || p[2] == "" || p[3] == "" {
		return fmt.Errorf("-sigv4 wants aws:amz:REGION:SERVICE (curl --aws-sigv4's form), not %q", provider)
	}
	region, service := p[2], p[3]
	keyID, secret, ok := strings.Cut(user, ":")
	if !ok || keyID == "" {
		return fmt.Errorf("-user wants ACCESS_KEY:SECRET_KEY")
	}
	amzDate := now.Format("20060102T150405Z")
	date := amzDate[:8]
	sum := sha256.Sum256(body)
	payload := hex.EncodeToString(sum[:])
	req.Header.Set("X-Amz-Date", amzDate)
	req.Header.Set("X-Amz-Content-Sha256", payload)
	path := req.URL.EscapedPath()
	if path == "" {
		path = "/"
	}
	q := req.URL.Query()
	keys := make([]string, 0, len(q))
	for k := range q {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var pairs []string
	for _, k := range keys {
		vs := append([]string(nil), q[k]...)
		sort.Strings(vs)
		for _, v := range vs {
			pairs = append(pairs, awsEscape(k)+"="+awsEscape(v))
		}
	}
	const signed = "host;x-amz-content-sha256;x-amz-date"
	canon := req.Method + "\n" + path + "\n" + strings.Join(pairs, "&") + "\n" +
		"host:" + req.URL.Host + "\n" + "x-amz-content-sha256:" + payload + "\n" + "x-amz-date:" + amzDate + "\n\n" +
		signed + "\n" + payload
	scope := date + "/" + region + "/" + service + "/aws4_request"
	h := sha256.Sum256([]byte(canon))
	sts := "AWS4-HMAC-SHA256\n" + amzDate + "\n" + scope + "\n" + hex.EncodeToString(h[:])
	k := hmacSHA256([]byte("AWS4"+secret), date)
	k = hmacSHA256(k, region)
	k = hmacSHA256(k, service)
	k = hmacSHA256(k, "aws4_request")
	req.Header.Set("Authorization", "AWS4-HMAC-SHA256 Credential="+keyID+"/"+scope+", SignedHeaders="+signed+
		", Signature="+hex.EncodeToString(hmacSHA256(k, sts)))
	return nil
}

func tcpProbe(addr string) {
	in, _ := io.ReadAll(os.Stdin)
	conn, err := net.DialTimeout("tcp", addr, 5*time.Second)
	if err != nil {
		fmt.Fprintln(os.Stderr, "rs-probe:", err)
		os.Exit(1)
	}
	defer conn.Close()
	conn.Write(in)
	conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	io.Copy(os.Stdout, conn)
}
