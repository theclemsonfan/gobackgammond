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

**`AMBER — STATIC_BASELINE_ESTABLISHED / EXECUTABLE_BASELINE_PENDING`**

The repository is small and structurally understandable, with pinned modules, an existing golden test, and explicit build/test/vet commands. It is materially more complex than Monte-Ball because requests cross a live HTTP boundary, serialized game state, AI selection, templates, and SVG output. Health cannot be promoted to green until a Go-capable environment runs the existing checks. Missing Go on this controller is an environment-provisioning partial blocker, not evidence of an application defect.

## Gap analysis

| Priority | Evidence-backed gap | Safe next evidence | Boundary |
|---|---|---|---|
| 1 | No executable baseline has run in the current retrofit environment. | Run version/env, test, vet, and build on this exact branch. | Do not change dependencies to fit the host. |
| 2 | Handler and route behavior has no direct automated coverage. | Characterize root, new-game, malformed-token, and SVG status/content contracts. | Preserve observable behavior. |
| 3 | Server lifecycle uses the default mux and a blocking `ListenAndServe`; README explicitly says it is not productionized. | Document test seams and lifecycle risks before proposing changes. | Productionization is owner-reserved. |
| 4 | Randomness is process-global and initialized in `main`, while handler behavior depends on it. | Use the existing seed contract to design deterministic tests. | No redesign or modernization in baseline work. |
| 5 | The golden SVG test is broad but failure output is coarse; coverage for adapter methods and error paths is unknown. | Run coverage and inspect gaps after baseline succeeds. | Add only dependency-free, behavior-preserving tests. |
| 6 | Historical Go/module compatibility is unspecified by a `go` directive. | Record the actual successful toolchain version. | Adding/changing the directive is a separate decision. |

## Production protection

The retrofit is isolated on `project-os/retrofit-pilot-2`, based on exact `master` SHA `2ab7266c52dbbccee113d7614e2af868eb9aabaa`. GitHub branch-protection rules were reported disabled, so operational protection is the explicit no-merge/no-deploy/no-production-change rule plus draft review. Application source remains byte-for-byte unchanged from the adoption base.
