# Build Spec

`gobuild` cross-compiles Go binaries. The version stamp is Go's own build
information - nothing is injected.

## Usage

```
gobuild build                        # every cmd/ program except example/
gobuild build goalertlog golint      # the named programs
gobuild build goalertlog --linux-amd64
gobuild build gosed --module github.com/funtimecoding/soil@v0.11.190
gobuild install goalertlog           # host architecture into ~/bin
gobuild install gosed --module github.com/funtimecoding/soil
gobuild stale                        # rebuild every stale binary on the PATH
```

Every mistake fails with a non-zero exit: no verb, an unknown verb or flag, a
name with a comma or a slash, a name without `cmd/<name>/main.go`, `--module`
without names, `--main` or `--output` with more than one name, and a failed
`go build`.

| Flag | Verbs | Meaning |
|------|-------|---------|
| `--linux-amd64`, `--darwin-arm64`, `--darwin-amd64` | build | Narrow the targets; none means all three |
| `--module path[@version]` | build, install | Build from a module instead of the work directory; no version means the latest tag |
| `--native` | build, install | Enable CGO |
| `--tags` | build, install | Build tags |
| `--main` | build | Main file, with exactly one name and no `--module` |
| `--output` | build | Output path, with exactly one name and one target |

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

## Remote Modules

`--module` builds a command of a module that is not the work directory, the
way a consumer builds soil's tools without vendoring them. The version
resolves through `go list -m`, `@latest` when none is given, the highest
released tag. `go install` refuses cross-compilation into `GOBIN`, so
`build.Module` writes a throwaway module in a temporary directory, requires
the command at the resolved version there, and runs the same `go build` per
target with that directory as its working directory. The binary is stamped as
the remote module at its version, with no version control fields, and lands
in the same output layout as a local build.

## Install Semantics

Local tool installation goes through `gobuild install <name>` -
never `go install`. Procfiles use `go run`.

`install` places the binary via `system.ReplaceFile` - write to a
temporary file beside the destination, then rename over it. The
destination always gets a fresh inode: in-place overwrite makes
macOS kill the next exec of a previously executed binary
(stale kernel signature cache), and truncating a running
executable on Linux fails with "text file busy". Rename has
neither problem - running processes keep the old inode.

`stale` installs what golint reports as stale: every binary on the
PATH whose `cmd/<name>` lives in a repository beside the current one and whose
version is below that repository's latest tag, or carries none
(`installed.Outdated`, the same list golint's finding reads). Each builds
natively from its repository's working tree, dirty or not - a working-tree
build stamps a pseudo-version past the tag, so a rerun skips what already went
through. The first build that fails stops the run.

## Target Architectures

By default, `build` builds all three; `install` builds the host's. Pass flags
to `build` to select specific ones:

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
cmd/gobuild/main.go               # entry point
pkg/tool/gobuild/
  main.go                         # Main(): cobra root, fails without a verb
  build_command.go                # build verb, target flags
  install_command.go              # install verb, host architecture
  stale_command.go                # stale verb
  validate.go                     # Validate(): the failing rules
  run.go                          # Run(): local or remote per name
  rebuild_stale.go                # rebuildStale(): installs installed.Outdated
  option/                         # command line option
pkg/build/
  go.go                           # Go(): runs go build on the package
  architectures.go                # Architectures(): iterates the targets
  module.go                       # Module(): remote build per target
  module_origin.go                # ModuleOrigin(): resolves the version
  throwaway.go                    # the throwaway module a remote build runs in
  guess_main_path.go              # GuessMainPath(): cmd/<name>/main.go
  package.go                      # Package(): main file to its package
  option/                         # build option
```
