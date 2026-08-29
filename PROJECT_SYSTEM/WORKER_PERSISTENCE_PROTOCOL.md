# Worker Persistence Protocol

Every bounded worker ends in exactly one state: `PERSISTED` (intended branch and exact contents independently verified), `PERSISTENCE_PACKET` (complete controller-applicable packet emitted), `NO_CHANGE` (legitimate no-change outcome), or `PERSISTENCE_BLOCKED` (direct persistence and a complete packet both impossible). A local commit alone is not `PERSISTED`.

A packet identifies repository, intended branch, starting SHA, commit message, actual checks, and each UTF-8 text action/path/previous blob SHA/complete new content. It never contains secret values.

The controller validates authorization and current head, rejects destructive/secret/deployment/production/material-cost/out-of-scope actions, applies only safe non-conflicting branch changes, and re-reads the resulting SHA and exact content.

For stale packets, compare intervening changed paths. Apply an independent create only if the path is still absent. Never replace or delete a path changed since packet start. Regenerate state/backlog/handoff from live evidence and record each action as applied, reconciled, or rejected.

