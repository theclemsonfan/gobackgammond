# Retrofit Control Plane

## CI evidence gate

Executable claims require a durable workflow or another explicitly identified execution host. Record the exact branch SHA, event, run and job IDs, runner OS, language/tool versions, commands, conclusions, coverage when measured, smoke scope, and whether evidence was read independently from the provider. A green job name alone is insufficient. Failed experimental assertions remain part of the evidence trail and must be classified as product defects, test defects, or unresolved.

The current baseline uses Ubuntu 24.04, Go `1.27.0`, `GOTOOLCHAIN=local`, `GOFLAGS=-mod=readonly`, the committed module checksums, and `govulncheck v1.7.0`. The GitHub Actions dependencies still use mutable major tags; changing them to commit SHAs is a separate supply-chain policy decision.

## Draft PR creation fallback

Attempt the supported API/connector path first. If GitHub returns `403 Resource not accessible by integration`, record the exact failure and do not broaden permissions, copy credentials, or retry destructively. If an already authenticated browser session is available, prepare the same base, head, title, body, and explicit draft/no-merge/no-deploy state in the GitHub UI. Creating the draft PR is an external publication action and requires the applicable action-time confirmation. After creation, independently verify PR number, draft/open/unmerged state, base/head refs and SHAs, mergeability, and unchanged default branch.

The durable connector fix is owner/developer-managed GitHub App `Pull requests: Read and write` scope followed by installation approval. A repository owner cannot add an undeclared app permission locally.

## Branch-protection boundary

Read the live protection state; never infer protection from branch names or PR status. When rules are absent, document that draft/no-merge/no-deploy guardrails are procedural only. Do not enable rules, required checks, merge queues, or permission changes without owner authorization. Keep the exact default-branch SHA in every persistence verification.

## Persistence and handoff verification

Before writing the handoff, record the observed branch head. After commit and push, independently re-read the remote branch head, PR head/base and state, default-branch SHA, successful checks, and exact changed-file content. The handoff commit cannot contain or verify its own resulting SHA; resolve it from live path history or later controller evidence. End in exactly one persistence terminal state.
