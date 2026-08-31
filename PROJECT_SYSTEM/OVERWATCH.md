# Overwatch

Use live PR/base/head SHA, durable Project OS files, backlog readiness, Owner Inbox blocking scope, active worker/controller cycles, previous nudge keys, and persistence state. Conversation summaries alone are not evidence.

Evaluate in order: `OWNER_ACTION_REQUIRED` when an owner/safety boundary blocks every safe stream; `APPROVED_IDLE` when authorized scope is durably complete; `HEALTHY_ACTIVE` when work or recent durable progress has an exact next action; otherwise `UNAUTHORIZED_IDLE` when safe work remains without a valid global blocker.

Use duplicate key `repository + PR/ref + head SHA + exact next-action fingerprint`. Suppress equivalent unresolved nudges, active matching work, and requests whose contents are already durable. A task summary or task-local commit never proves persistence. Record pre/post SHA and exact file evidence after every relaunch or controller recovery.

Current classification: `APPROVED_IDLE`. The fourth-batch code and documentation are persisted and independently verified; master Project OS `2.1.0-master` is packaged and its regression suite passes 9/9. Draft PR #1 is open, draft, cleanly mergeable, and unmerged; `master` remains unchanged and both branches remain procedurally guarded rather than protected. Duplicate key: `theclemsonfan/gobackgammond|PR-1|384b7f497695476a32b4b9fbe7fe7902f1c953eb|pilot-approved-idle-master-2.1.0`.

Continuation maturity is `MONITOR_ONLY` → `TRIGGER_AVAILABLE` → `DIRECT_PERSISTENCE_UNVERIFIED` → `CONTROLLER_PERSISTENCE_AVAILABLE` → `AUTO_NUDGE_AVAILABLE`; do not skip levels based on summaries.
