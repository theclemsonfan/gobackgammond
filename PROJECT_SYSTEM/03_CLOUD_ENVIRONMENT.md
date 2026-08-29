# Cloud Environment Preflight

- Language/runtime: Go; historical repository has no `go` directive in `go.mod`.
- Package manager: Go modules with committed `go.mod` and `go.sum`.
- Direct dependencies: `github.com/ajstarks/svgo` 2018 pseudo-version; `github.com/chandler37/gobackgammon v0.1.7`.
- Build/test tools: `go`, `make`, network access to the configured Go module proxy/cache.
- Environment variables/secrets: none evidenced for build/test; server uses command-line `-port` and optional `-seed`.
- Platform: ordinary Go-supported host capable of binding localhost for manual smoke testing.

## Required preflight

```bash
go version
go env GOMOD GOPROXY
go test ./...
go vet ./...
go build .
```

## Current controller evidence

The repository was cloned read-only at exact SHA `2ab7266c52dbbccee113d7614e2af868eb9aabaa`, but this controller host has no `go` executable. Executable validation is therefore `ENVIRONMENT_PROVISIONING_PARTIAL_BLOCKER`, not an application defect. Do not change `go.mod`, dependency pins, or source to accommodate this host. Static analysis and Project OS work continue.

