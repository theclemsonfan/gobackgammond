# Roadmap

## R0 — Protected adoption

Isolated branch exists and is persisted; draft PR #1 is open, draft, and unmerged; `master` remains unchanged.

## R1 — Reproducible baseline

Complete. The unchanged application passed test, vet, build, startup, and route smoke checks on the verified Go `1.27.0` cloud baseline without modernization.

## R2 — Behavior-preserving validation

Complete for this safe batch. Root, routing errors, token parsing, turn-selection, new-game, invalid-state, compressed and uncompressed state, AI ordering, serialization round trips, template escaping, SVG, take-turn, and existing rendering contracts are covered. Coverage is 70.2% for handlers, 100.0% for SVG, and 66.4% repository-wide. Deterministic victory coverage remains a future evidence task; lifecycle changes remain owner-reserved.

## R3 — Autonomous continuity

Complete for the second-pilot lessons. CI evidence, browser fallback, branch-protection limitations, persistence verification, stale-packet reconciliation, and exact handoff are durable.

## R4 — Owner review

Reconcile live branch evidence. Merge, deployment, server productionization, and dependency/toolchain modernization remain separate owner decisions.

## R5 — First routine maintenance use

In progress on the isolated branch under stable Project OS `2.1.0`. The bounded milestone adds deterministic victory-route evidence derived from the pinned dependency's supported board model, refreshes lifecycle/maintenance risks, reruns the full branch CI contract, and persists an independently verifiable handoff. It does not change production source, dependencies, `master`, or PR draft status.
