# Build Spec

`gobuild` cross-compiles Go binaries. The version stamp is Go's own build
information - nothing is injected.

## Usage

```
gobuild goalertlog          # build one binary
gobuild                     # build all binaries in cmd/ (except example/)
gobuild --copy-to-bin       # also install matching-architecture binary to ~/bin
gobuild --linux-amd64       # build only linux-amd64
gobuild --native            # enable CGO
```

## What It Does

For each target architecture, `gobuild` runs:

```
go build -tags timetzdata -o tmp/<name>/<os>-<arch>/<name> ./cmd/<name>
```

It builds the package, never the file: a file-path build
(`cmd/<name>/main.go`) records no module and no version control state, while
a package build stamps both. `pkg/stamp` reads them at runtime through
`runtime/debug`:

- **Version** - the module version Go derives from the tags: `v0.11.170` on a
  tag, a pseudo-version such as `v0.11.171-0.20261003213813-fd20dc25a022`
  between tags, `+dirty` appended for uncommitted changes. `--version` prints
  a pseudo-version as the tag it follows (`v0.11.170 (untagged)`); the JSON
  report keeps Go's value.
- **GitHash** - `vcs.revision`, shortened
- **CommitDate** - `vcs.time`, the commit's time rather than the build's
- **Module** - the main module path
- **Dirty** - `vcs.modified`

A `go install`ed binary carries the same stamp, minus the version control
fields when installed from the module proxy. A build without a git checkout
fails unless `-buildvcs=false` is passed; CI checks out with full history, so
tags and revision are present.

## Install Semantics

Local tool installation goes through `gobuild --copy-to-bin <name>` -
never `go install`. Procfiles use `go run`.

`--copy-to-bin` installs via `system.ReplaceFile` - write to a
temporary file beside the destination, then rename over it. The
destination always gets a fresh inode: in-place overwrite makes
macOS kill the next exec of a previously executed binary
(stale kernel signature cache), and truncating a running
executable on Linux fails with "text file busy". Rename has
neither problem - running processes keep the old inode.

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

## Entry Point

See `entrypoint.md` for the `Main()` convention. `main` passes nothing; the
stamp comes from `stamp.New()`.

`gobuild` locates the entry point via `build.GuessMainPath(name)`, which looks
for `cmd/<name>/main.go`. Override with `--main` flag. Either way the build
target is the file's package (`build.Package`).

## Packages

```
cmd/gobuild/main.go              # entry point
pkg/tool/gobuild/main.go         # Main(): flags, dispatch
pkg/build/
  go.go                          # Go(): runs go build on the package
  architectures.go               # Architectures(): iterates selected targets
  guess_main_path.go             # GuessMainPath(): cmd/<name>/main.go
  package.go                     # Package(): main file to its package
  option/
    build.go                     # Build option struct
    new.go                       # constructor
```
