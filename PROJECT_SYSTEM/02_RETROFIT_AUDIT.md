# Retrofit Audit and Health Classification

## Inspection boundary

This audit is derived from read-only inspection of `README.md`, `Makefile`, `go.mod`, `go.sum`, `main.go`, `handlers/handlers.go`, `svg/svg.go`, `svg/svg_test.go`, and Git history at adoption base `2ab7266c52dbbccee113d7614e2af868eb9aabaa`. No application code, dependency, default branch, deployment, or production state was changed.

## Baseline architecture

- `main.go` parses required `-port` and optional `-seed`, seeds the process-global random source, registers three routes on Go's default HTTP mux, and starts a blocking HTTP server.
- `handlers` owns request parsing, board-state serialization/compression, AI continuation ordering, HTML rendering, and error responses for `/`, `/game`, and `/game.svg`.
- local `svg` adapts the external SVGo writer to the board renderer.
- dependencies are pinned in Go modules; the historical module intentionally has no `go` directive.
- the Makefile exposes format, test, coverage, benchmark, vet, build, run, server, documentation, and clean targets.
- the only current automated check is a deterministic golden SVG rendering test.

## Health classification

**`AMBER — EXECUTABLE_BASELINE_GREEN / REMEDIATION_IN_PROGRESS`**

The executable baseline is green on the isolated retrofit branch. GitHub Actions run `33281569224` verified the unchanged application at commit `68919e3ef4a8f9db8c0e73d133941dc81831eee4` with Go `1.27.0` on Linux/amd64: test, vet, build, process startup, `/`, and `/game` all passed. Run `33281651519` then verified the pinned toolchain and dependency-free handler characterization tests at commit `2190711eafa31b6d4c4f16f034002caf02a1f930`.

Overall health remains amber because the draft review boundary is not yet published, coverage is still narrow, server lifecycle productionization is explicitly out of scope, and branch protection is procedural. No application defect was demonstrated by this batch.

## Gap analysis

| Priority | Evidence-backed gap | Safe next evidence | Boundary |
|---|---|---|---|
| 1 | The draft PR review/control surface is not yet published. | Create it as draft with explicit no-merge/no-deploy guardrails. | Never merge or enable auto-merge. |
| 2 | Handler coverage is still intentionally narrow after adding root, token, turn-selection, new-game, and invalid-state contracts. | Measure coverage and add only high-value behavior-preserving cases. | Preserve observable behavior and dependencies. |
| 3 | Server lifecycle uses the default mux and a blocking `ListenAndServe`; README explicitly says it is not productionized. | Document test seams and lifecycle risks before proposing changes. | Productionization is owner-reserved. |
| 4 | Randomness is process-global and initialized in `main`, while handler behavior depends on it. | Use the existing seed contract to design deterministic tests. | No redesign or modernization in baseline work. |
| 5 | The golden SVG test is broad but failure output is coarse; adapter-method coverage is unknown. | Run coverage and inspect gaps. | Add only dependency-free, behavior-preserving tests. |
| 6 | Historical Go/module compatibility is unspecified by a `go` directive. | Keep CI pinned to the verified Go `1.27.0` baseline. | Adding/changing the module directive is a separate owner decision. |

## Production protection

The retrofit is isolated on `project-os/retrofit-pilot-2`, based on exact `master` SHA `2ab7266c52dbbccee113d7614e2af868eb9aabaa`. GitHub branch-protection rules were reported disabled, so operational protection is the explicit no-merge/no-deploy/no-production-change rule plus draft review. Application source remains byte-for-byte unchanged from the adoption base.
