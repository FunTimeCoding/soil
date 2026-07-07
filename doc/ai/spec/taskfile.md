# Taskfile Spec

Standard project automation using [Task](https://taskfile.dev). The taskfile defines the local development pipeline and integrates with lefthook for git hooks.

## Standard Tasks

```yaml
version: '3'
tasks:
  default: {cmds: [{task: lint}, {task: test}, {task: build}], silent: true}
  pre-push: {cmds: [{task: lint}, {task: test}], silent: true}
  lint: ...
  test: ...
  build: ...
  check: ...
  generate: ...
  update: ...
  tool: ...
```

### Pipeline Tasks

| Task       | Steps                 | Purpose                                   |
|------------|-----------------------|-------------------------------------------|
| `default`  | lint -> test -> build | Full local pipeline (`task` with no args) |
| `pre-push` | lint -> test          | Git pre-push hook via lefthook            |

### Individual Tasks

| Task | Command | Purpose |
|------|---------|---------|
| `lint` | `golint --fix` + `golangci-lint run --build-tags local` + `gofix ./...` + `goanalyze ./...` + `goaudit .` | Lint with auto-fix, static analysis, AST fix and check, then compliance audit |
| `test` | `gotestsum --format dots -- --tags local ./...` | Run tests (including local-tagged) with minimal output |
| `build` | `go run cmd/gobuild/main.go` | Cross-compile all binaries via gobuild |
| `check` | `gosec -fmt=json ...` | Security scan |
| `generate` | `oapi-codegen --config config.yaml ...` per service | Regenerate OpenAPI clients and servers (see `generated-api.md`) |
| `update` | `goupdate` with pinned downgrades | Dependency update with known-bad version pins |
| `tool` | installs gotestsum, golint, goanalyze, gofix, goaudit, gobuild, golangci-lint, oapi-codegen | Dev tooling bootstrap |

## Lefthook Integration

```yaml
pre-push: {jobs: [{run: task pre-push}]}
```

The `pre-push` hook runs lint + test before every push. Configured in `.lefthook.yaml`.

## GitHub Actions

Minimal CI pipeline in `.github/workflows/build.yml`:

```yaml
name: Build
on: {push: {branches: [main]}, pull_request: {branches: [main]}}
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - {name: Repository, uses: actions/checkout@v4}
      - {name: Go, uses: actions/setup-go@v5, with: {go-version-file: go.mod}}
      - {name: Dependency, run: 'go mod tidy'}
      - {name: Build, run: 'go build -json ./...'}
      - {name: Test, run: 'go test -json ./...'}
```

Go version is pinned via `go.mod`, not hardcoded in the workflow.

## Conventions

- All tasks use `silent: true`
- Multi-line commands use `|` block scalar
- `build` task delegates to `gobuild` (see `build.md`), not raw `go build`
- `tool` task is the single source of truth for dev dependencies
- `update` task pins known-incompatible dependency versions with `--downgrade`
