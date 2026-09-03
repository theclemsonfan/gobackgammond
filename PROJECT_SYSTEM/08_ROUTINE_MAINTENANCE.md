# Routine Maintenance Register

This register turns the completed retrofit pilot into routine use of stable AI Project OS `2.1.0`. It separates evidence work that may proceed on the isolated branch from changes that need a future plan or explicit owner approval.

| Area | Current evidence | Next safe maintenance | Boundary |
|---|---|---|---|
| Victory route | A completed board can be represented using the pinned dependency's serialization and `TakeTurn` contract. | Characterize the HTTP victory result with a deterministic test and retain CI output. | Do not alter scoring, gameplay, templates, or dependencies. |
| HTTP lifecycle | `main` uses global flags, the default mux, process-global RNG, and blocking `http.ListenAndServe`; smoke tests exercise the process externally. | Keep process smoke coverage and document failure seams. | Dedicated mux/server, timeouts, shutdown, limits, or RNG injection are architecture/product changes. |
| Dependencies | Direct versions and checksums are historical; pinned `govulncheck v1.7.0` found no reachable known vulnerability at the last verified head. | Re-run the pinned scanner with each maintenance batch and preserve exact results. | Upgrades, a `go` directive, or a support matrix require owner approval. |
| CI supply chain | Go and scanner versions are explicit; GitHub Actions references use major tags. | Record resolved runs, runner, tools, and commands. | Pinning action SHAs and selecting an update policy are governance decisions. |
| Branch safety | PR #1 is draft/unmerged and `master` is unchanged; branch rules are absent. | Re-read state before and after every batch and preserve no-merge/no-deploy language. | Enabling protection, required checks, or merge policy changes governance. |

## Lifecycle failure seams

- Route registration is process-global, so repeated in-process `main` execution is not an isolated test seam.
- `ListenAndServe` blocks and the application has no graceful-shutdown path; current evidence must come from a bounded subprocess smoke test.
- Startup errors are returned to the standard logger, while readiness is inferred by retrying localhost requests.
- Process-global random seeding makes handler behavior deterministic only when tests remain serial and explicitly seed before relevant calls.
- No timeout, connection-limit, or structured operational logging contract is present; green local smoke evidence does not establish production readiness.

## Routine cadence

For each authorized batch: verify live control state, run the exact behavior-preserving CI contract, make only reversible isolated-branch changes, classify failures without rewriting the product contract, persist exact evidence, and independently read back branch/PR/default-branch state. Merge, deploy, productionization, modernization, governance changes, and Internet exposure remain owner-reserved.
