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
5. `R2-010` — converted the amber analysis into this prioritized split backlog and reconciled testing/handoff evidence. Draft PR publication remains the final external control-surface action for the requested batch.

## Project OS / process gaps

| Priority | ID | Gap | Evidence / next safe action | Decision boundary |
|---|---|---|---|---|
| P0 | `PROC-001` | Draft PR review surface is not yet published. | Create draft PR from `project-os/retrofit-pilot-2` to `master`; preserve explicit no-merge/no-deploy language. | Browser submission needs action-time safety confirmation; merge remains owner-reserved. |
| P1 | `PROC-002` | `master` and retrofit branches are not protected by GitHub rules. | Keep draft/no-merge procedural guardrails; separately propose branch protection. | Enabling repository rules changes governance and requires owner approval. |
| P1 | `PROC-003` | Action dependencies use major-version tags, not immutable SHAs. | Propose SHA pinning with update procedure. | Supply-chain policy choice; do not silently freeze/update actions. |
| P2 | `PROC-004` | Controller host cannot run Go locally. | Use the verified cloud workflow; do not retrofit the controller. | No blocker while GitHub Actions is available. |
| P2 | `PROC-005` | No documented Go-version support policy exists. | Treat Go `1.27.0` as the observed CI baseline only. | Adding a `go` directive or support matrix is owner-reserved modernization. |

## Application defects

No application defect is currently demonstrated. The unchanged application passed tests, vet, build, startup, `/`, and `/game` on Go `1.27.0`; new characterization tests also passed.

## Application risks / unproven gaps

| Priority | ID | Risk or gap | Next safe evidence | Decision boundary |
|---|---|---|---|---|
| P1 | `APP-RISK-001` | The default mux and blocking server lifecycle limit isolated lifecycle testing and graceful shutdown. | Document seams and failure modes; do not refactor yet. | Productionization is owner-reserved. |
| P1 | `APP-RISK-002` | Randomness is process-global, so test isolation depends on serial seeded execution. | Add deterministic cases only when they do not require concurrency or redesign. | Randomness redesign is product/architecture work. |
| P1 | `APP-RISK-003` | Route coverage remains narrow, especially valid serialized-state and victory flows. | Measure coverage and add reversible characterization cases. | Preserve current observable behavior. |
| P2 | `APP-RISK-004` | SVG golden-test failure output is coarse and adapter coverage is unknown. | Record coverage, then improve test diagnostics without changing rendering. | No dependency changes. |

Move repository-changing work to completed only after `PERSISTED`; a partial blocker never stops unrelated ready work.
