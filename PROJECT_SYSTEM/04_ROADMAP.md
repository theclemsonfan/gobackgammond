# Roadmap

## R0 — Protected adoption

Isolated branch exists and is persisted; `master` remains unchanged and unmerged. Draft PR publication is still pending.

## R1 — Reproducible baseline

Static baseline and `AMBER — STATIC_BASELINE_ESTABLISHED / EXECUTABLE_BASELINE_PENDING` health classification are recorded. Run the Makefile-equivalent Go build/test/vet commands in a Go-capable cloud environment and record exact evidence without modernization.

## R2 — Behavior-preserving validation

Characterize HTTP routes, deterministic seed behavior, handler inputs/state flow, SVG rendering, and server lifecycle. Add only small reversible tests where gaps are proven.

## R3 — Autonomous continuity

Validate five-task continuation, owner-question deferral, Overwatch classification, duplicate suppression, persistence verification, stale-packet reconciliation, and exact handoff.

## R4 — Owner review

Reconcile live branch evidence. Merge, deployment, server productionization, and dependency/toolchain modernization remain separate owner decisions.
