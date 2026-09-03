# Overwatch

Use live PR/base/head SHA, durable Project OS files, backlog readiness, Owner Inbox blocking scope, active worker/controller cycles, previous nudge keys, and persistence state. Conversation summaries alone are not evidence.

Evaluate in order: `OWNER_ACTION_REQUIRED` when an owner/safety boundary blocks every safe stream; `APPROVED_IDLE` when authorized scope is durably complete; `HEALTHY_ACTIVE` when work or recent durable progress has an exact next action; otherwise `UNAUTHORIZED_IDLE` when safe work remains without a valid global blocker.

Use duplicate key `repository + PR/ref + head SHA + exact next-action fingerprint`. Suppress equivalent unresolved nudges, active matching work, and requests whose contents are already durable. A task summary or task-local commit never proves persistence. Record pre/post SHA and exact file evidence after every relaunch or controller recovery.

Current classification: `APPROVED_IDLE`, subject to controller post-handoff persistence readback. The first routine-use batch under stable Project OS `2.1.0` added deterministic victory evidence and a refreshed maintenance register on the isolated branch. Push run `33819014529` / job `100857382602` and PR run `33819017762` / job `100857392347` passed at `6b7312b439cf2df291910f38e0e4cf8ecd9f89fd`. Draft PR #1 remained open, draft, and unmerged; `master` remained `2ab7266c52dbbccee113d7614e2af868eb9aabaa`. Duplicate key: `theclemsonfan/gobackgammond|PR-1|6b7312b439cf2df291910f38e0e4cf8ecd9f89fd|routine-use-complete-2.1.0`.

Continuation maturity is `MONITOR_ONLY` → `TRIGGER_AVAILABLE` → `DIRECT_PERSISTENCE_UNVERIFIED` → `CONTROLLER_PERSISTENCE_AVAILABLE` → `AUTO_NUDGE_AVAILABLE`; do not skip levels based on summaries.
