# Build Spec

`gobuild` cross-compiles Go binaries with version metadata injected via linker flags.

## Usage

```
gobuild goalertlog          # build one binary
gobuild                     # build all binaries in cmd/ (except example/)
gobuild --copy-to-bin       # also copy matching-architecture binary to ~/bin
gobuild --linux-amd64       # build only linux-amd64
gobuild --native            # enable CGO
```

## What It Does

For each target architecture, `gobuild` runs:

```
go build -ldflags "-X main.Version=v0.10.294 -X main.GitHash=144f841a -X main.BuildDate=2026-02-20T13:36:08+01:00" -o tmp/<name>/<os>-<arch>/<name> cmd/<name>/main.go
```

- **Version** - latest git tag (`git describe --tags --abbrev=0`)
- **GitHash** - short commit hash (`git rev-parse --short HEAD`)
- **BuildDate** - current time in RFC3339

## Target Architectures

By default, all three are built. Pass flags to select specific ones:

| Flag             | Target                       |
|------------------|------------------------------|
| `--linux-amd64`  | Linux AMD64                  |
| `--darwin-arm64` | Darwin ARM64 (Apple Silicon) |
| `--darwin-amd64` | Darwin AMD64 (Intel Mac)     |

## Output Layout

```
tmp/<name>/
  linux-amd64/<name>
  darwin-arm64/<name>
  darwin-amd64/<name>
```

## Entry Point Contract

See `entrypoint.md` for the linker variable and `Main()` convention. `gobuild` injects `Version`, `GitHash`, and `BuildDate` at compile time.

## Main Path Resolution

`gobuild` locates the entry point via `build.GuessMainPath(name)`, which looks for `cmd/<name>/main.go`. Override with `--main` flag.

## Packages

```
cmd/gobuild/main.go              # entry point
pkg/tool/gobuild/main.go         # Main(): flags, dispatch
pkg/build/
  go.go                          # Go(): runs go build with ldflags
  architectures.go               # Architectures(): iterates selected targets
  git_tag.go                     # GitTag(): latest git tag
  git_hash.go                    # GitHash(): short commit hash
  date.go                        # Date(): RFC3339 now
  guess_main_path.go             # GuessMainPath(): cmd/<name>/main.go
  option/
    build.go                     # Build option struct
    new.go                       # constructor
```
