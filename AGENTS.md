# AI Project OS Agent Instructions

This repository uses AI Project OS `2.0.0-master` on the isolated retrofit branch `project-os/retrofit-pilot-2`. GitHub branch-protection rules are not currently enabled, so safety depends on the draft-PR/no-merge guardrail.

Read every file in `PROJECT_SYSTEM/` before changing application code. Preflight the Go environment before executable claims. Missing tools are environment problems first; do not silently update Go modules or application behavior to fit the host.

Work toward the durable milestone goal in batches of five safe outcomes when five are ready. Defer non-global owner questions to `06_OWNER_INBOX.md` with a recommended default and continue independent work.

Do not merge, deploy, alter production, expose secrets, force-push, delete important data, incur material cost, or modernize dependencies without explicit authorization. A task-local commit is not persistence proof. End with exactly one of `PERSISTED`, `PERSISTENCE_PACKET`, `NO_CHANGE`, or `PERSISTENCE_BLOCKED`, independently verified against the intended GitHub branch, and update `12_SESSION_HANDOFF.md`.
