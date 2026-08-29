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

