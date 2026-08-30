# Backlog

## Completed first five safe outcomes

1. `R2-001` — selected and fingerprinted the multi-component HTTP/handler/SVG pilot at exact base SHA.
2. `R2-002` — created the isolated retrofit branch without changing `master` or production.
3. `R2-003` — installed the repository-specific environment manifest and scoped missing Go as a partial blocker.
4. `R2-004` — recorded the baseline contract, existing test surface, build commands, dependency pins, and safety boundaries.
5. `R2-005` — completed the static baseline, health classification, evidence-backed gap analysis, Overwatch/persistence state, and durable handoff.

## Completed second five-task batch outcomes

1. `R2-006` — established and verified the branch-only GitHub Actions Go environment.
2. `R2-007` — ran the unchanged executable baseline: test, vet, build, startup, root, and new-game smoke checks passed.
3. `R2-008` — pinned the successfully observed Go `1.27.0` toolchain without changing `go.mod` or dependencies.
4. `R2-009` — added dependency-free characterization for root HTML, token parsing, turn selection, new-game HTML, and invalid game state handling; cloud verification passed.
5. `R2-010` — converted the amber analysis into this prioritized split backlog and reconciled testing/handoff evidence. At that batch boundary, draft PR publication was the remaining external control-surface action; it is now completed under `R2-011`.

## Completed third five-task batch outcomes

1. `R2-011` — independently verified draft PR #1: open, draft, unmerged, base `master` at `2ab7266c52dbbccee113d7614e2af868eb9aabaa`.
2. `R2-012` — reconfirmed the exact Go `1.27.0` Linux/amd64 baseline on the branch-scoped Ubuntu 24.04 workflow.
3. `R2-013` — ran test, coverage, vet, build, startup, `/`, and `/game` checks at `79a53610e5707f333e27e164d26981c6e0a2878b`; all passed after correcting one invalid test assumption.
4. `R2-014` — audited the PR diff and separated Project OS/process gaps, actual application defects, and legacy/dependency risks; no production-source change or demonstrated application defect was found.
5. `R2-015` — added valid compressed-state, SVG response, and take-turn characterization plus durable coverage reporting; no application behavior or dependency changed.

## Project OS / process gaps

| Priority | ID | Gap | Evidence / next safe action | Decision boundary |
|---|---|---|---|---|
| P0 | `PROC-001` | Draft PR review surface was missing. | Completed: PR #1 is open, draft, and unmerged with no-merge/no-deploy language. | Merge remains owner-reserved. |
| P1 | `PROC-002` | `master` and retrofit branches are not protected by GitHub rules. | Keep draft/no-merge procedural guardrails; separately propose branch protection. | Enabling repository rules changes governance and requires owner approval. |
| P1 | `PROC-003` | Action dependencies use major-version tags, not immutable SHAs. | Propose SHA pinning with update procedure. | Supply-chain policy choice; do not silently freeze/update actions. |
| P2 | `PROC-004` | Controller host cannot run Go locally. | Use the verified cloud workflow; do not retrofit the controller. | No blocker while GitHub Actions is available. |
| P2 | `PROC-005` | No documented Go-version support policy exists. | Treat Go `1.27.0` as the observed CI baseline only. | Adding a `go` directive or support matrix is owner-reserved modernization. |

## Application defects

No application defect is currently demonstrated. The unchanged production application passed tests, coverage, vet, build, startup, `/`, and `/game` on Go `1.27.0`; expanded characterization tests also passed. Run `33283733265` exposed an invalid proposed victory-test assumption, which was corrected without changing application code.

## Application risks / unproven gaps

| Priority | ID | Risk or gap | Next safe evidence | Decision boundary |
|---|---|---|---|---|
| P1 | `APP-RISK-001` | The default mux and blocking server lifecycle limit isolated lifecycle testing and graceful shutdown. | Document seams and failure modes; do not refactor yet. | Productionization is owner-reserved. |
| P1 | `APP-RISK-002` | Randomness is process-global, so test isolation depends on serial seeded execution. | Add deterministic cases only when they do not require concurrency or redesign. | Randomness redesign is product/architecture work. |
| P1 | `APP-RISK-003` | Valid serialized-state and take-turn paths are covered, but deterministic victory coverage is still absent. | Find or construct a dependency-supported terminal state before adding a reversible contract. | Preserve current observable behavior. |
| P2 | `APP-RISK-004` | SVG golden-test failure output is coarse and adapter coverage is unknown. | Record coverage, then improve test diagnostics without changing rendering. | No dependency changes. |

## Legacy / dependency risks

| Priority | ID | Risk | Next safe evidence | Decision boundary |
|---|---|---|---|---|
| P1 | `LEGACY-001` | Direct dependencies are historically pinned; their current vulnerability/support posture is unverified. | Run read-only advisory analysis when an approved tool is available. | Do not update dependencies without owner approval. |
| P1 | `LEGACY-002` | `go.mod` has no `go` directive or support policy. | Continue exact Go `1.27.0` evidence. | Module-policy change is owner-reserved. |
| P1 | `LEGACY-003` | Default mux, blocking server lifecycle, and global RNG reduce isolation and graceful-shutdown options. | Preserve documented seams and failure modes. | Refactoring is product/architecture work. |

Move repository-changing work to completed only after `PERSISTED`; a partial blocker never stops unrelated ready work.
