# GENERATED L-009 hermetic build (--network=none, SBOM+SLSA-attestation-producing).
# Pinned bases from BASE_IMAGE_DIGESTS.lock; vendored Go tree + debs from Dockerfile.vendor.
FROM golang:1.26.11-bookworm@sha256:dc40a13b61304647ed7e5d581d246b86413dfc29d63c6d018310feb93cd98fc0 AS build_temporal-server
WORKDIR /src
COPY vendored/temporal-src/ .
ENV CGO_ENABLED=0 GOTOOLCHAIN=local GOWORK=off GOFLAGS="-mod=vendor -buildvcs=false" GOPROXY=off SOURCE_DATE_EPOCH=1795164612
RUN go build -trimpath -tags "disable_grpc_modules" -ldflags='-s -w' -o /temporal-server ./cmd/server && /temporal-server --version && \
    { readelf -lW /temporal-server | sed -n 's/.*Requesting program interpreter: \(.*\)]$/\1/p'; readelf -dW /temporal-server | sed -n 's/.*Shared library: \[\(.*\)\]$/\1/p'; } > /temporal-server.needs
FROM golang:1.26.11-bookworm@sha256:dc40a13b61304647ed7e5d581d246b86413dfc29d63c6d018310feb93cd98fc0 AS build_temporal-sql-tool
WORKDIR /src
COPY vendored/temporal-src/ .
ENV CGO_ENABLED=0 GOTOOLCHAIN=local GOWORK=off GOFLAGS="-mod=vendor -buildvcs=false" GOPROXY=off SOURCE_DATE_EPOCH=1795164612
RUN go build -trimpath -tags "disable_grpc_modules" -ldflags='-s -w' -o /temporal-sql-tool ./cmd/tools/sql && /temporal-sql-tool --version && \
    { readelf -lW /temporal-sql-tool | sed -n 's/.*Requesting program interpreter: \(.*\)]$/\1/p'; readelf -dW /temporal-sql-tool | sed -n 's/.*Shared library: \[\(.*\)\]$/\1/p'; } > /temporal-sql-tool.needs
FROM golang:1.26.11-bookworm@sha256:dc40a13b61304647ed7e5d581d246b86413dfc29d63c6d018310feb93cd98fc0 AS build_tdbg
WORKDIR /src
COPY vendored/temporal-src/ .
ENV CGO_ENABLED=0 GOTOOLCHAIN=local GOWORK=off GOFLAGS="-mod=vendor -buildvcs=false" GOPROXY=off SOURCE_DATE_EPOCH=1795164612
RUN go build -trimpath -tags "disable_grpc_modules" -ldflags='-s -w' -o /tdbg ./cmd/tools/tdbg && /tdbg --version && \
    { readelf -lW /tdbg | sed -n 's/.*Requesting program interpreter: \(.*\)]$/\1/p'; readelf -dW /tdbg | sed -n 's/.*Shared library: \[\(.*\)\]$/\1/p'; } > /tdbg.needs
# rs-probe: the HEALTHCHECK helper (standard library only, from the kit's rs-probe.go)
FROM golang:1.26.11-bookworm@sha256:dc40a13b61304647ed7e5d581d246b86413dfc29d63c6d018310feb93cd98fc0 AS build_rs_probe
WORKDIR /probe
COPY rs-probe.go main.go
COPY rs-probe.mod go.mod
ENV CGO_ENABLED=0 GOTOOLCHAIN=local GOFLAGS="-buildvcs=false" GOPROXY=off SOURCE_DATE_EPOCH=1795164612
RUN go build -trimpath -ldflags='-s -w' -o /rs-probe .
# the chiselled rootfs (vendored/chisel, cut by Dockerfile.vendor), the service user, the binaries
FROM ubuntu:24.04@sha256:9c7f200c07e3d4a03e7249d2c16bf66a51120fb876aa2242cf31451412543289 AS rootfs
ARG SOURCE_DATE_EPOCH=1795164612
ADD vendored/chisel/rootfs.tar /rootfs/
RUN printf 'temporal:x:1000:1000::/etc/temporal:/bin/sh\n' >> /rootfs/etc/passwd && printf 'temporal:x:1000:\n' >> /rootfs/etc/group && \
    mkdir -p /rootfs/etc/temporal /rootfs/var/lib/temporal && chown -R 1000:1000 /rootfs/etc/temporal /rootfs/var/lib/temporal
COPY --from=build_temporal-server /temporal-server /rootfs/usr/bin/temporal-server
COPY --from=build_temporal-server /temporal-server.needs /tmp/needs/temporal-server
COPY --from=build_temporal-sql-tool /temporal-sql-tool /rootfs/usr/bin/temporal-sql-tool
COPY --from=build_temporal-sql-tool /temporal-sql-tool.needs /tmp/needs/temporal-sql-tool
COPY --from=build_tdbg /tdbg /rootfs/usr/bin/tdbg
COPY --from=build_tdbg /tdbg.needs /tmp/needs/tdbg
COPY --from=build_rs_probe /rs-probe /rootfs/usr/local/bin/rs-probe
COPY entrypoint.sh /rootfs/usr/local/bin/entrypoint.sh
COPY dynamicconfig.yaml /rootfs/etc/temporal/config/dynamicconfig/docker.yaml
RUN chmod 755 /rootfs/usr/bin/temporal-server /rootfs/usr/bin/temporal-sql-tool /rootfs/usr/bin/tdbg /rootfs/usr/local/bin/rs-probe && \
    chown 1000:1000 /rootfs/usr/local/bin/entrypoint.sh && \
    chown 1000:1000 /rootfs/etc/temporal/config/dynamicconfig/docker.yaml && \
    chmod 0755 /rootfs/usr/local/bin/entrypoint.sh && \
    chmod 0644 /rootfs/etc/temporal/config/dynamicconfig/docker.yaml
RUN --mount=type=bind,source=vendored/chisel,target=/chisel set -eu; \
    for n in /tmp/needs/*; do while read -r lib; do [ -n "$lib" ] || continue; \
      case "$lib" in /*) [ -e "/rootfs$lib" ] ;; *) [ -n "$(find /rootfs/usr/lib /rootfs/usr/lib64 -name "$lib" 2>/dev/null | head -1)" ] ;; esac || \
      { echo "$(basename "$n") needs $lib, which the chiselled rootfs does not have: add its slice to runtime_base.slices" >&2; exit 1; }; \
    done < "$n"; done; rm -rf /tmp/needs; \
    [ -e /rootfs/etc/os-release ] || { echo "no /etc/os-release: base-files_release-info is missing" >&2; exit 1; }; \
    [ "$(grep -c '^Package: ' /rootfs/var/lib/dpkg/status)" = "$(wc -l < /chisel/packages.tsv)" ] || \
      { echo "/var/lib/dpkg/status does not list every package of vendored/chisel/packages.tsv" >&2; exit 1; }; \
    rm -rf /rootfs/var/log/dpkg.log /rootfs/var/log/alternatives.log /rootfs/var/log/apt /rootfs/var/cache/ldconfig/aux-cache /rootfs/var/cache/fontconfig
FROM scratch
ARG SOURCE_DATE_EPOCH=1795164612
COPY --from=rootfs /rootfs/ /
ENV PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin DB=sqlite SQLITE_SETUP=true SQLITE_MODE=rwc NUM_HISTORY_SHARDS=4 BIND_ON_IP=0.0.0.0 TEMPORAL_ADDRESS=127.0.0.1:7233 DBNAME=/var/lib/temporal/temporal.db VISIBILITY_DBNAME=/var/lib/temporal/temporal_visibility.db
LABEL org.opencontainers.image.title="Temporal"
LABEL org.opencontainers.image.version="v1.32.2-CE-ubuntu24.04-r5"
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.source="https://github.com/temporalio/temporal"
LABEL org.opencontainers.image.vendor="ReadyStack.dev"
LABEL readystack.product="temporal"
LABEL readystack.edition="CE"
LABEL readystack.build_method="source-rebuild-go"
LABEL org.opencontainers.image.description="Temporal runs your long-lived business processes as code that survives crashes, restarts and downtime. Workflows keep their place automatically, so retries, timeouts and recovery stop being your application's problem."
LABEL org.opencontainers.image.revision="6d085822273f4336796e6959c73ab1d14a23eb77"
HEALTHCHECK --interval=5s --timeout=30s --start-period=40s \
  CMD /usr/local/bin/rs-probe http://127.0.0.1:7243/api/v1/namespaces || exit 1
USER temporal
EXPOSE 7233 7243
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
