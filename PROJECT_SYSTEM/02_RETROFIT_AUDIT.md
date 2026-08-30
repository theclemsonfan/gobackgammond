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

**`AMBER — EXECUTABLE_BASELINE_GREEN / DRAFT_REVIEW_ACTIVE`**

The executable baseline is green on the isolated retrofit branch. GitHub Actions run `33281569224` verified the unchanged application at commit `68919e3ef4a8f9db8c0e73d133941dc81831eee4` with Go `1.27.0` on Linux/amd64: test, vet, build, process startup, `/`, and `/game` all passed. Run `33281651519` then verified the pinned toolchain and dependency-free handler characterization tests at commit `2190711eafa31b6d4c4f16f034002caf02a1f930`.

Run `33283773288` verified the expanded characterization at commit `79a53610e5707f333e27e164d26981c6e0a2878b`: test, coverage, vet, build, startup, `/`, and `/game` passed. Handler statement coverage is 67.5%, SVG coverage is 100.0%, and repository-wide statement coverage is 64.2%.

Overall health remains amber because branch protection is procedural, the historical dependencies and module contract have no declared support policy, and server lifecycle productionization is explicitly out of scope. Draft PR #1 is open, draft, and unmerged. No application defect was demonstrated by this batch.

## Current PR diff classification

### Project OS / process gaps

- GitHub branch-protection rules remain disabled; safety depends on the draft/no-merge guardrail.
- Actions use major-version tags rather than immutable commit SHAs.
- The controller host lacks Go; the branch-scoped cloud workflow is the verified execution path.
- The repository has no documented Go support matrix or module `go` directive.

### Actual application defects

None demonstrated. One proposed victory assertion failed in run `33283733265`; inspection showed that the test assumed an immediate victory the application never promises. The test was corrected to characterize the valid take-turn outcome without changing application behavior.

### Legacy / dependency risks

- Direct dependencies are pinned to `github.com/ajstarks/svgo` at a 2018 pseudo-version and `github.com/chandler37/gobackgammon v0.1.7`; compatibility and vulnerability posture were not modernized or independently certified.
- The historical module omits a `go` directive, so the verified Go `1.27.0` baseline is evidence, not a declared minimum-support policy.
- Process-global randomness and default-mux/blocking-server lifecycle constrain test isolation and graceful shutdown.
- These are risks or architectural constraints, not proven defects, and remediation requires owner policy or product decisions.

## Gap analysis

| Priority | Evidence-backed gap | Safe next evidence | Boundary |
|---|---|---|---|
| 1 | GitHub branch protection is disabled. | Keep draft/no-merge guardrails; separately propose rules. | Governance change requires owner approval. |
| 2 | Handler coverage is 67.5% after adding root, token, turn-selection, new-game, invalid-state, valid-state, SVG, and take-turn contracts. | Add only high-value behavior-preserving cases for still-uncovered error/victory seams. | Preserve observable behavior and dependencies. |
| 3 | Server lifecycle uses the default mux and a blocking `ListenAndServe`; README explicitly says it is not productionized. | Document test seams and lifecycle risks before proposing changes. | Productionization is owner-reserved. |
| 4 | Randomness is process-global and initialized in `main`, while handler behavior depends on it. | Use the existing seed contract to design deterministic tests. | No redesign or modernization in baseline work. |
| 5 | The golden SVG test is broad but failure output is coarse; adapter-method coverage is unknown. | Run coverage and inspect gaps. | Add only dependency-free, behavior-preserving tests. |
| 6 | Historical Go/module compatibility is unspecified by a `go` directive. | Keep CI pinned to the verified Go `1.27.0` baseline. | Adding/changing the module directive is a separate owner decision. |

## Production protection

The retrofit is isolated on `project-os/retrofit-pilot-2`, based on exact `master` SHA `2ab7266c52dbbccee113d7614e2af868eb9aabaa`. GitHub branch-protection rules were reported disabled, so operational protection is the explicit no-merge/no-deploy/no-production-change rule plus draft PR #1. Production application source remains byte-for-byte unchanged from the adoption base; only test code was added under `handlers/`.
