# Testing and Evidence

## Repository-derived checks

The Makefile defines `go test ./...`, `go vet ./...`, `go build .`, coverage, and benchmarks. `svg/svg_test.go` is the existing automated test file.

## Controller preflight

- Base SHA verified: `2ab7266c52dbbccee113d7614e2af868eb9aabaa`.
- Read-only clone: pass.
- `go version`: blocked; `go` executable absent.
- `go test ./...`, `go vet ./...`, `go build .`: not run because preflight failed.
- Classification: environment-provisioning partial blocker.

Do not misstate static inspection as executable or CI evidence. Future results must record command, host, timestamp, branch SHA, output/result, and whether the evidence is local, cloud-run, CI, or independently read from GitHub.

## Static audit evidence

- Application source differs from adoption base: no.
- Project OS branch locally matches its recorded remote-tracking ref at `b860de7c339d0e97e85827a75723db1cba09a391` before this audit write.
- Architecture, baseline contract, health classification, gaps, and production boundary are recorded in `02_RETROFIT_AUDIT.md`.
- No executable Go claim was made.

For handoff writes, the pre-handoff SHA anchors the state used to compose the file. The commit containing a handoff cannot embed and verify its own SHA; post-write branch SHA and exact content require independent observation.
