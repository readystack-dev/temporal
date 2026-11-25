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

`recipe/` holds the kit's recipe files (Dockerfiles, build scripts, pins) as published at https://readystack.dev/products/temporal/recipe. The complete kit is the release download. Each release is a tag here.

— ReadyStack · https://readystack.dev
