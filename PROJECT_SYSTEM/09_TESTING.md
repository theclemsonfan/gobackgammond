# Testing and Evidence

## Repository-derived checks

The Makefile defines `go test ./...`, `go vet ./...`, `go build .`, coverage, and benchmarks. `svg/svg_test.go` is the existing automated test file.

## Cloud baseline — unchanged application

- Source: GitHub Actions run `33281569224`, job `99177497343`.
- Timestamp: 2026-08-29 23:41–23:42 UTC.
- Branch/SHA: `project-os/retrofit-pilot-2` at `68919e3ef4a8f9db8c0e73d133941dc81831eee4`.
- Environment: Ubuntu 24.04 runner; `go version go1.27.0 linux/amd64`; `GOPROXY=https://proxy.golang.org,direct`.
- `go test ./...`: pass; root and handlers reported no test files; `svg` passed.
- `go vet ./...`: pass.
- `go build -o "$RUNNER_TEMP/gobackgammond" .`: pass; no artifact published.
- Startup smoke: pass on localhost port 18080 with seed 37.
- Root route smoke: pass; expected Backgammon title observed.
- New-game route smoke: pass; expected Backgammon Game title observed.

## Cloud remediation verification

- Source: GitHub Actions run `33281651519`, job `99177711365`.
- Timestamp: 2026-08-29 23:44 UTC.
- Branch/SHA: `project-os/retrofit-pilot-2` at `2190711eafa31b6d4c4f16f034002caf02a1f930`.
- Environment: pinned `go version go1.27.0 linux/amd64`.
- `go test ./...`: pass; handlers `0.009s`, SVG cached, root package has no test files.
- `go vet ./...`: pass.
- `go build`: pass.
- Startup, root route, and new-game route smoke checks: pass.
- New test contracts: root response, token missing/single/empty/multiple/malformed cases, turn-selection query behavior, deterministic new-game response, invalid game state rejection for HTML and SVG.

Do not misstate static inspection as executable or CI evidence. Future results must record command, host, timestamp, branch SHA, output/result, and whether the evidence is local, cloud-run, CI, or independently read from GitHub.

## Static audit evidence

- Production application source differs from adoption base: no; only workflow, test, and Project OS files changed.
- Architecture, baseline contract, health classification, gaps, and production boundary are recorded in `02_RETROFIT_AUDIT.md`.
- Executable claims above are derived from independently read GitHub Actions job steps and logs, not static inspection.

For handoff writes, the pre-handoff SHA anchors the state used to compose the file. The commit containing a handoff cannot embed and verify its own SHA; post-write branch SHA and exact content require independent observation.
