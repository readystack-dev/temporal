# Temporal v1.32.2 — ReadyStack release `v1.32.2-CE-ubuntu24.04-r5`

Generated 2026-11-20 from this release's build records and strategy. Statements that come from a single record name it in brackets; the others come from this release's strategy and build log.

## Release

- **Image:** `readystack/temporal:v1.32.2-CE-ubuntu24.04-r5` (linux/amd64). Index digest `sha256:c581b89648e0e4bd05c42904884ba3f833517770c2f2580f77bc769d44390cc3` [`docs/provenance/LOCKED_DIGEST.txt`].
- **Image manifest digest:** `sha256:584e42db12a15563fb531476753d6cafc0b1f939c68aa2bc11b3e13694b37ce4` [`TARGET_DIGEST.txt`].
- **Upstream:** Temporal v1.32.2 built from source at `v1.32.2` of https://github.com/temporalio/temporal — upstream notes: https://github.com/temporalio/temporal/releases/tag/v1.32.2 [`docs/provenance/BUILD_RECIPE.json`].
- **Base image:** none (`FROM scratch`): a chiselled Ubuntu 24.04 runtime, the files Canonical's chisel cut from Ubuntu's own packages (`vendored/chisel/slices.txt`, `vendored/chisel/packages.tsv`); runtime user `temporal`.
- **Build:** `source-rebuild-go`, compiled offline with `--network=none`.
- **Supersedes:** ReadyStack release `v1.32.2-CE-ubuntu24.04-r4`.

## What changed in this release

- Upstream version unchanged (v1.32.2); ReadyStack rebuild.
- Minimum dependency versions ReadyStack enforces on top of upstream (raised only where upstream's own version was lower); shipped versions from the SBOM [`docs/sbom/SBOM.json`]:
  - `golang.org/x/crypto` raised from upstream's `v0.57.0` to at least `v0.58.0` (shipped: `v0.58.0`)
  - `google.golang.org/grpc` at least `v1.83.2`: not raised, upstream already pins `v1.83.2` (shipped: `v1.83.2`)
  - `golang.org/x/net` raised from upstream's `v0.60.0` to at least `v0.62.0` (shipped: `v0.62.0`)
  - `github.com/klauspost/compress` raised from upstream's `v1.18.7` to at least `v1.20.2` (shipped: `v1.20.2`)
  - `github.com/go-jose/go-jose/v4` at least `v4.1.6`: not raised, upstream already pins `v4.1.6` (shipped: `v4.1.6`)
  - `github.com/apache/thrift` raised from upstream's `v0.24.0` to at least `v0.25.0` (shipped: `v0.25.0`)
- Temporal v1.32.2 is unchanged from the release this supersedes (v1.32.2-CE-ubuntu24.04-r4); this release rebuilds it. Changes: the image is rebuilt on the current ubuntu:24.04 image (sha256:9c7f200c07e3…), which its maintainers updated after v1.32.2-CE-ubuntu24.04-r4 was built on sha256:2198e4c65076…. Settings, ports and data volumes are those of v1.32.2-CE-ubuntu24.04-r4: replace the container with this image on the same volumes.

## Security status at release (measured 2026-11-20)

- **Docker Scout:** 0 fixable critical/high vulnerabilities; runs as `temporal` (non-root); SBOM + provenance attestations present (SPDX SBOM + SLSA v1) [`docs/provenance/SCOUT_VERDICT.json`].
- **OSV cross-check:** no fixable critical/high vulnerability beyond Docker Scout's findings.
- **Signature:** cosign signature verified by ReadyStack's release gate before publishing; check it yourself with `cosign verify --key https://readystack.dev/keys/cosign.pub readystack/temporal:v1.32.2-CE-ubuntu24.04-r5` (signatures are recorded in the public Sigstore transparency log).
- **ReadyStack release check at release:** passed (ReadyStack's own test, not a Docker Scout grade). Vulnerabilities are disclosed every day and the image never changes, so it is re-checked daily, beside Docker Scout's own dated grade, at https://readystack.dev/queue.
- **Before exposing this service:** Temporal's authorizer and claim mapper are off by default: any client that can reach the frontend ports (7233 gRPC, 7243 HTTP) has full control of every namespace and workflow. Keep the ports on a private network, or put mTLS or an authorizing proxy in front of them.
- **Before exposing this service:** With the JWT authorizer turned on, the internal service ports 7236 and 7246 still answer without a token to any container that can reach them (for example on the same Docker network); publish only the frontend ports and keep the container on a network with nothing you do not trust.

## Reproducibility

Rebuilds offline from this archive to the identical image digest `sha256:584e42db12a15563fb531476753d6cafc0b1f939c68aa2bc11b3e13694b37ce4`: run `./BUILD_OFFLINE.sh`, then `./VERIFY_OFFLINE.sh` (add `--online` to confirm the Docker Hub tag holds the same digest). Before publishing, ReadyStack rebuilt the kit with `./BUILD_OFFLINE.sh` and compared the result with the published digest [`docs/provenance/REPRO_VERIFICATION.json`], then ran `./BUILD_OFFLINE.sh` and `./VERIFY_OFFLINE.sh` from this archive (without `--online`).

## Validation

- 5 of 5 use cases passed (5 authored) against the running image [`docs/validation/test-results.json`].

## Known issues

- **high** (upstream): With Temporal's JWT authorizer on, the internal service ports 7236/7246 still answer without a token (https://github.com/temporalio/temporal/pull/3706) [`RSD-temporal-1bde533fc0`] (carried from `v1.32.0-CE-ubuntu24.04-r1`, recorded there as `RSD-temporal-71f1c69845`)

## Defects fixed

- 28 problems found before release and fixed before it shipped (28 by review); each is listed in `docs/defects/defects.json`.

## Verify this archive

```
sha256sum -c CHECKSUMS.sha256
./BUILD_OFFLINE.sh && ./VERIFY_OFFLINE.sh --online
cosign verify --key https://readystack.dev/keys/cosign.pub readystack/temporal:v1.32.2-CE-ubuntu24.04-r5
```

ReadyStack records every signature in the public Sigstore transparency log (Rekor), so the command needs no extra flags.

