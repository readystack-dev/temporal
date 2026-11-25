# Provenance — Temporal (`v1.32.2-CE-ubuntu24.04-r5`)

Every item below is copied from a record in this archive; the record is named in brackets.

## Source

- Upstream repository: https://github.com/temporalio/temporal
- Revision built: `v1.32.2` (commit `6d085822273f4336796e6959c73ab1d14a23eb77`)
- Source tree: `96b41e631f563f71248b899d3d4aca1d1507334e5e489c7be7a793f7bb6c39e6` over 3792 files [`docs/provenance/SOURCE_TREE_HASH.json`]
- Community Edition only: enterprise code is removed or refused before the build (CE audit at acquire).

## Build

- Method: `source-rebuild-go`; base image `scratch + a chiselled Ubuntu 24.04 runtime (vendored/chisel/slices.txt, packages.tsv)`
- Compiled with `--network=none` from vendored source, dependencies and base images; BuildKit pinned to `moby/buildkit:v0.31.2` (shipped in `toolchain/`).
- SOURCE_DATE_EPOCH: `1795164612` [`SOURCE_DATE_EPOCH.txt`]
- Frozen inputs: [`docs/provenance/FREEZE_MANIFEST.json`], build recipe: [`docs/provenance/BUILD_RECIPE.json`].

## Output

- Published image: `readystack/temporal:v1.32.2-CE-ubuntu24.04-r5` — index digest `sha256:c581b89648e0e4bd05c42904884ba3f833517770c2f2580f77bc769d44390cc3` [`docs/provenance/LOCKED_DIGEST.txt`]
- linux/amd64 image digest `sha256:584e42db12a15563fb531476753d6cafc0b1f939c68aa2bc11b3e13694b37ce4` — reproduced from this archive before publishing [`docs/provenance/REPRO_VERIFICATION.json`].

## Attestations and signature

- SBOM + SLSA provenance attached to the image: attached: an SPDX SBOM and SLSA v1 provenance [`docs/provenance/SCOUT_VERDICT.json`]
- cosign signature by the ReadyStack release key: verified before publishing. Check it: `cosign verify --key https://readystack.dev/keys/cosign.pub readystack/temporal:v1.32.2-CE-ubuntu24.04-r5`

## Pinned inputs, and why

- 31 pinned inputs, each with where it came from, why this one, how it was checked and when it was decided: `docs/PINS.md` [`docs/provenance/PINS.json`] (upstream source 2, base and build images 3, image digests 2, dependency versions readystack sets 6, pins in the kit's dockerfiles 11, readystack's build engine 5, dependencies upstream chose 2).

## Build trail

- How this release passed through the ReadyStack harness, record by record (every chain, drive, stage, review round, defect and check, each naming the record it is copied from): `ReadyStack-temporal-v1.32.2-CE-ubuntu24.04-r5.trail.json`, published beside this archive and signed with the release key (`ReadyStack-temporal-v1.32.2-CE-ubuntu24.04-r5.trail.json.sigstore`), and readable at https://readystack.dev/products/temporal/trail/v1.32.2-CE-ubuntu24.04-r5.

