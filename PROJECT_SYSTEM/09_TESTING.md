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

## Expanded characterization and coverage

- Failed evidence retained: push run `33283733265`, job `99183176018`, commit `7b492996b7d8aebf7ad173e9ad7cda9fd6f2e188`, failed in `go test ./...` because a proposed test incorrectly assumed the documented nearly-finished game must produce immediate victory after one seeded turn. Vet, build, and smoke were skipped after the test failure. No application defect was demonstrated.
- Correction: commit `79a53610e5707f333e27e164d26981c6e0a2878b` changed only that test contract from immediate-victory assertion to valid take-turn outcome characterization.
- Successful evidence: push run `33283773288`, job `99183283407`, and PR run `33283775027`, 2026-08-30 00:38 UTC, Go `1.27.0` Linux/amd64 on Ubuntu 24.04.
- `go test ./...`: pass; handlers `0.009s`, SVG cached, root package has no tests.
- Coverage: handlers 67.5%, SVG 100.0%, root package 0.0%, total statements 64.2%; profile stayed in runner temporary storage and no artifact was published.
- `go vet ./...`: pass.
- `go build -o "$RUNNER_TEMP/gobackgammond" .`: pass.
- Startup on localhost port 18080 with seed 37, root route, and new-game route: pass.
- Added contracts: valid compressed state renders game HTML, valid compressed state renders an SVG response with the expected media type, and a seeded take-turn request returns either the continued-game or victory template permitted by current behavior.

## Hardened reproducibility, risk coverage, and security evidence

- Source: GitHub Actions push run `33349587401`, job `99360182171`, and PR run `33349589876`.
- Timestamp: 2026-08-31 02:05–02:06 UTC.
- Branch/SHA: `project-os/retrofit-pilot-2` at `69eab3c9852e647373c8757af7b8f3ac8e881494`.
- Environment: explicit Ubuntu 24.04; Go `1.27.0`; `GOTOOLCHAIN=local`; `GOFLAGS=-mod=readonly`; `GOPROXY=https://proxy.golang.org,direct`.
- `go test ./...`: pass. Added handler-level malformed/multiple-token routing checks, uncompressed serialization compatibility for HTML and SVG, AI-ranked continuation and compressed serialization round trips, multiple turn parameters, and dynamic template escaping.
- Coverage: handlers 70.2%, SVG 100.0%, root package 0.0%, total statements 66.4%. `smartBoards` reached 93.8%; token, turn selection, root, SVG handler, and template constructors reached 100%. The remaining `GameHandler` seam is primarily deterministic victory/error construction; `main` remains covered by process smoke rather than unit tests.
- `go vet ./...`: pass.
- `go run golang.org/x/vuln/cmd/govulncheck@v1.7.0 ./...`: pass; `No vulnerabilities found.` This is a known-vulnerability database result, not proof that the historical dependencies are maintained or risk-free.
- `go build -o "${RUNNER_TEMP}/gobackgammond" .`: pass; no artifact published.
- Startup/smoke: pass on localhost port 18080 with seed 37; root HTML, new-game HTML, valid SVG rendering, and invalid-state HTTP 400/error text all matched.
- Independently read live state after the run: PR #1 remained open, draft, mergeable/clean, and unmerged; base `master` remained `2ab7266c52dbbccee113d7614e2af868eb9aabaa`; head was `69eab3c9852e647373c8757af7b8f3ac8e881494`; both branches reported protection disabled.
- Documentation persistence verification: push run `33350002209` and PR run `33350004490` passed at `384b7f497695476a32b4b9fbe7fe7902f1c953eb`. Local and remote `PROJECT_SYSTEM/` blob SHAs matched exactly; PR #1 remained open/draft/clean/unmerged, `master` remained unchanged, and both branches still reported protection disabled.

Do not misstate static inspection as executable or CI evidence. Future results must record command, host, timestamp, branch SHA, output/result, and whether the evidence is local, cloud-run, CI, or independently read from GitHub.

## Static audit evidence

- Production application source differs from adoption base: no; only workflow, test, and Project OS files changed.
- Architecture, baseline contract, health classification, gaps, and production boundary are recorded in `02_RETROFIT_AUDIT.md`.
- Executable claims above are derived from independently read GitHub Actions job steps and logs, not static inspection.

For handoff writes, the pre-handoff SHA anchors the state used to compose the file. The commit containing a handoff cannot embed and verify its own SHA; post-write branch SHA and exact content require independent observation.
