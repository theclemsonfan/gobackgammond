# Roadmap

## R0 — Protected adoption

Isolated branch exists and is persisted; draft PR #1 is open, draft, and unmerged; `master` remains unchanged.

## R1 — Reproducible baseline

Complete. The unchanged application passed test, vet, build, startup, and route smoke checks on the verified Go `1.27.0` cloud baseline without modernization.

## R2 — Behavior-preserving validation

In progress. Root, token parsing, turn-selection, new-game, invalid-state, valid compressed-state, SVG route, take-turn, and existing SVG rendering contracts are covered by dependency-free tests. Coverage is measured at 67.5% for handlers, 100.0% for SVG, and 64.2% repository-wide. Deterministic victory coverage and owner-reserved lifecycle changes remain.

## R3 — Autonomous continuity

Validate five-task continuation, owner-question deferral, Overwatch classification, duplicate suppression, persistence verification, stale-packet reconciliation, and exact handoff.

## R4 — Owner review

Reconcile live branch evidence. Merge, deployment, server productionization, and dependency/toolchain modernization remain separate owner decisions.
