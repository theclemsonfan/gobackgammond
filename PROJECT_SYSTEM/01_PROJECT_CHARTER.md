# Project Charter

## Verified purpose and behavior

`gobackgammond` is an HTTP user interface and AI player built on `gobackgammon` and `svgo`. `main.go` accepts `-port` and an optional deterministic `-seed`, then registers `/`, `/game`, and `/game.svg`. The `handlers` package owns game flow and the `svg` package renders boards. Existing behavior is the baseline contract until a defect is demonstrated.

## Authorized retrofit scope

- Project OS control-plane files on `project-os/retrofit-pilot-2`.
- Read-only inspection and reproducible build/test/runtime evidence.
- Later small, reversible tests that do not change application behavior.

No merge, deployment, production change, dependency modernization, secret change, destructive action, or material cost is authorized.

## Success criteria

A fresh worker can resume from repository state, the Go environment is explicit, existing checks are reproducible or precisely blocked, five-task batches survive partial blockers, and branch persistence/next action are independently verifiable.

