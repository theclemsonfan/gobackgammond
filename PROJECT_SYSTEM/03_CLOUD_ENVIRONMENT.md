# Cloud Environment Baseline

- Language/runtime: Go; historical repository has no `go` directive in `go.mod`.
- Package manager: Go modules with committed `go.mod` and `go.sum`.
- Direct dependencies: `github.com/ajstarks/svgo` 2018 pseudo-version; `github.com/chandler37/gobackgammon v0.1.7`.
- Build/test tools: `go`, `make`, network access to the configured Go module proxy/cache.
- Environment variables/secrets: none evidenced for build/test; server uses command-line `-port` and optional `-seed`.
- Platform: GitHub-hosted Ubuntu runner capable of binding localhost for smoke testing.

## Reproducible cloud path

`.github/workflows/retrofit-go-baseline.yml` is branch-scoped to `project-os/retrofit-pilot-2`, pull requests targeting `master`, and explicit manual dispatch. It has `contents: read`, a ten-minute timeout, no deployment, no artifact publication, and no production credentials.

The verified baseline is:

- GitHub Actions runner image: `ubuntu-24.04`, image version `20260823.283.1`.
- Go: `go1.27.0 linux/amd64`, now pinned exactly in the workflow.
- Module: `/home/runner/work/gobackgammond/gobackgammond/go.mod`.
- Proxy: `https://proxy.golang.org,direct`.
- Actions: `actions/checkout@v7` and `actions/setup-go@v7`.
- First executable baseline: run `33281569224`, commit `68919e3ef4a8f9db8c0e73d133941dc81831eee4`, success.
- Remediation verification: run `33281651519`, commit `2190711eafa31b6d4c4f16f034002caf02a1f930`, success.

The controller host still has no local Go executable, but that is no longer a validation blocker because the cloud path is durable and verified. Do not add a `go` directive, change dependency pins, or modify application behavior merely to match this environment.
