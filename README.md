[![ReadyStack](https://readystack.dev/logo.svg)](https://readystack.dev)

**Stable, Open-source Releases, Automated** · [readystack.dev](https://readystack.dev)

[![Docker Hub](https://img.shields.io/badge/Docker%20Hub-readystack%2Ftemporal-2496ED?logo=docker&logoColor=white)](https://hub.docker.com/r/readystack/temporal) [![Docker pulls](https://img.shields.io/docker/pulls/readystack/temporal)](https://hub.docker.com/r/readystack/temporal) [![Grade at release](https://img.shields.io/badge/grade%20at%20release-A%20%2894%25%29-2ea44f)](https://readystack.dev/products/temporal) [![Signed: cosign](https://img.shields.io/badge/signed-cosign-2ea44f)](https://readystack.dev/keys/cosign.pub) ![SBOM: SPDX](https://img.shields.io/badge/SBOM-SPDX-2ea44f) ![Provenance: SLSA v1](https://img.shields.io/badge/provenance-SLSA%20v1-2ea44f) ![Offline rebuild: verified](https://img.shields.io/badge/offline%20rebuild-verified-2ea44f)

# Temporal: ReadyStack build kit

The offline build kit for the Docker image `readystack/temporal:v1.32.2-CE-ubuntu24.04-r5`. It holds what the image is built from (upstream source, base images, toolchain, pinned dependencies) and the documents, so the image can be rebuilt with no network access (`docker build --network=none`).

- Release: [v1.32.2-CE-ubuntu24.04-r5](https://github.com/readystack-dev/temporal/releases/tag/v1.32.2-CE-ubuntu24.04-r5)
- Image: [`docker.io/readystack/temporal:v1.32.2-CE-ubuntu24.04-r5`](https://hub.docker.com/r/readystack/temporal/tags), manifest digest `sha256:584e42db12a15563fb531476753d6cafc0b1f939c68aa2bc11b3e13694b37ce4`
- Product page: https://readystack.dev/products/temporal
- Source and build recipe: https://readystack.dev/products/temporal/recipe
- Build trail: https://readystack.dev/products/temporal/trail/v1.32.2-CE-ubuntu24.04-r5
- Defects: this repository's Issues (defects recorded from 25 November 2026 on) and https://readystack.dev/defects.html?product=temporal

## The image

Single container running the temporal-server binary in the foreground, with all four services (frontend, history, matching, worker) in one process and a self-contained pure-Go SQLite store. This is a durable-execution backend, not an application with a UI — clients connect with a Temporal SDK over gRPC on 7233 and run workflows against it.

## Who uses Temporal and why

Developers who build long-running, multi-step processes that must finish even when services crash.

Temporal is a durable execution platform. Its server runs units of application logic called Workflows so that they resume where they left off after crashes or outages, and it retries failed operations automatically.

A typical use: An online shop's order process charges a card, reserves stock, books a courier and emails the customer, and any step can fail. Written as a Temporal Workflow, the process retries a failed step, carries on after a server restart without losing its place, and can wait days for a delivery confirmation.

From the upstream project's own description: https://github.com/temporalio/temporal, https://docs.temporal.io/.

## This release

- Docker Scout policy grade at release: **A** (94%), graded 2026-11-20 with `docker scout policy` (CLI 1.26.0, docker-default configuration).
  - Docker Scout counts 15 packages with copyleft licences (GPL, LGPL, MPL), all of them operating-system packages (it counts only packages whose licence it can read; the SBOM lists every component).
- Signed with ReadyStack's release key (fingerprint `04daa92344a52dbc`): https://readystack.dev/keys/cosign.pub.
- Runs as a non-root user (`temporal`).
- Attestations on the image: an SPDX SBOM and SLSA provenance v1.
- Rebuilt from this kit with network egress blocked on 2026-11-20; the rebuilt image digest is the published one.
- Base: Ubuntu 24.04.

The current grade, measured again every day, is on https://readystack.dev/products/temporal.

Temporal is distributed under its upstream licence (MIT); the kit contains upstream's source under that licence, and each vendored component under its own.

## Download

From the [v1.32.2-CE-ubuntu24.04-r5 release](https://github.com/readystack-dev/temporal/releases/tag/v1.32.2-CE-ubuntu24.04-r5): `ReadyStack-temporal-v1.32.2-CE-ubuntu24.04-r5.tar.gz` (506.8 MB), SHA-256 `00af72f97e99a0905ce1d4b8eaa17db0d4aaaa0ba206c014ea0ecc0daf90d269`.

```sh
sha256sum --check --ignore-missing SHA256SUMS
```

## Verify and rebuild

The rebuild statement is signed with the ReadyStack release key. Check the signature before you rely on it:

```sh
cosign verify-blob --key https://readystack.dev/keys/cosign.pub \
  --bundle ReadyStack-temporal-v1.32.2-CE-ubuntu24.04-r5.rebuild.json.sigstore ReadyStack-temporal-v1.32.2-CE-ubuntu24.04-r5.rebuild.json
```

Unpack with `-p` so file modes are kept, then follow the kit's own README.md and docs/REBUILD.md:

```sh
tar -xpzf ReadyStack-temporal-v1.32.2-CE-ubuntu24.04-r5.tar.gz
cd v1.32.2-CE-ubuntu24.04-r5
less README.md docs/REBUILD.md
```

The build trail (`ReadyStack-temporal-v1.32.2-CE-ubuntu24.04-r5.trail.json`, signed the same way, and its Markdown copy) records how this release was built and checked.

## This repository

`recipe/` holds the kit's recipe files (Dockerfiles, build scripts, pins) as published at https://readystack.dev/products/temporal/recipe. The complete kit is the release download. Each release is a tag here. How to report a vulnerability: [SECURITY.md](SECURITY.md).

## About ReadyStack

**Stable, Open-source Releases, Automated.** ReadyStack releases popular open-source apps every week with every available critical and high security fix applied, plus full documentation, reference data, tests and a build recipe that recreates each release bit-for-bit. Each release is rebuilt from its pinned inputs with the network off, checked against its published digest, signed, and scanned and graded again every day. The images and build kits are free, and ReadyStack is soon to be a nonprofit.

[readystack.dev](https://readystack.dev) · [All products](https://readystack.dev/products.html) · [Build kits on GitHub](https://github.com/readystack-dev) · [Defect log](https://readystack.dev/defects.html) · contact@readystack.dev

— ReadyStack · https://readystack.dev
