# Temporal: ReadyStack build kit

The offline build kit for the Docker image `readystack/temporal:v1.32.2-CE-ubuntu24.04-r5`. It holds what the image is built from (upstream source, base images, toolchain, pinned dependencies) and the documents, so the image can be rebuilt with no network access (`docker build --network=none`).

- Release: [v1.32.2-CE-ubuntu24.04-r5](https://github.com/readystack-dev/temporal/releases/tag/v1.32.2-CE-ubuntu24.04-r5)
- Image: [`docker.io/readystack/temporal:v1.32.2-CE-ubuntu24.04-r5`](https://hub.docker.com/r/readystack/temporal/tags), manifest digest `sha256:584e42db12a15563fb531476753d6cafc0b1f939c68aa2bc11b3e13694b37ce4`
- Product page: https://readystack.dev/products/temporal
- Defects: this repository's Issues (defects recorded from 25 November 2026 on) and https://readystack.dev/defects.html?product=temporal

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
