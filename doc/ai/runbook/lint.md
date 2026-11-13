# Lint

## Invocation

```
task lint
```

Runs golangci-lint, then `gocheck --fix`: golint, goaudit, gofix and
goanalyze as stages of one run, each under its own header line naming the
repository. Every stage reports even when an earlier one found something, and
the run fails if any did. golint's and gofix's fixes are written once at the
end, and not at all if a file they touch changed meanwhile - the refusal names
each such file. A package that does not load is a finding, and gofix and
goanalyze skip that run. golangci-lint takes the repository's custom build tags
from `goreach --tags`, the set gocheck loads with, so a file behind any tag is
linted - a hand-written tag list drifts, and a directory whose files all carry
a missing tag fails the run. After the headers the pipeline is silent on
success - any further output is a failure. gocheck skips a package it checked
clean before at an unchanged key - its files, the APIs it uses, the packages it
imports, the workspace's interfaces and the constants in reach of it - so a run
loads only what moved; `gocheck --fresh` checks everything again. golint also
names every installed binary older than its repository's latest tag;
`gobuild stale` installs them all.

gocheck asks gosourced to run the check when `GOSOURCE_TOKEN` is set and the
daemon serves the repository, and runs in its own process otherwise - the same
code, so the same findings. gosourced runs the lint code it started with: a
run says so in one line on stderr when that code changed since, and
`gocheck --local` runs the code on disk.

`task lint-check` runs the same without writing: golangci-lint without
`--fix`, and gocheck printing one diff of what golint and gofix would change.
goanalyze checks the code as it would be after those fixes, with its lines
translated back to the files on disk. A green check means `task lint` has
nothing to change. Use it to look; fix with `task lint`. `REACH=own` narrows it
to what the current session changed, `REACH=<session>` to what a named one did:
`goreach` reads the session's transcript - the files its edit tools wrote and
the ones gosourced reports - keeps those git still sees changed, and prints the
directories they reach, each changed package and every package importing it,
read from import lines so code that does not type-check still counts. goaudit
stays whole-repository. A reach with nothing changed ends the task early. A
reach takes in whole packages, so another session's changed files ride along,
and their findings come with yours.
Bash writes join the reach where the command names its files
literally - `mv` and `git mv` (both sides), `cp` (the destination),
`rm`, `git rm`, `rmdir`, `sed -i`, `tee`, `touch` and redirects -
resolved from the directory the transcript recorded and any `cd`
before them. A glob operand (`rm tmp/*.md`) takes in the files git sees
changed that it matches, and counts as globbed. `goreach` counts the
rest apart: operands built from variables or loops, fixer runs (their
own run is their check) and scripts.
`go test $(goreach)` tests the same reach.

Run the full pipeline once per completed code scope - it covers
the whole repository and takes minutes. For doc-only edits, run
`golint <path>` on the changed files instead - the header says
which repository and scope it selected, a path that doesn't exist
exits non-zero, and the whole repository is still loaded so the
markdown pointer checks stay exact. `--root` points it at another
repository from wherever you stand, and `GOLINT_CONFIGURATION` (or
`--configuration`) names an optional YAML with repository-private
vocabulary such as container registries. Comment-only and other
lint-inert edits need no run at all. `golint --census` prints the
markdown references the pointer check could not resolve, sorted by
reason with a count per reason; `--verbose` is the per-file trace.

Markdown front matter anchors the pointer check: `base:` names the
package a document's short paths and `route:` spans resolve
against, and `hosts:` lists the hosts its `http` and `https`
locators may name, and `commands:` lists slash-led spans that are
neither skills nor paths (an ACL path, a znode, a harness command a
document describes). Each takes a scalar or a list. A locator
against an undeclared host fails; `localhost` and `127.0.0.1` need
no declaration, and other schemes are only tallied in the census.

## Relevant specs

Read these before fixing lint issues or adding analyzers, relative
to the plugin root:

- `doc/ai/spec/conventions/` - formatting, file structure,
  constants, and import rules that lint enforces
- `doc/ai/spec/naming.md` - variable names, error
  progression, package names, constants
- `doc/ai/spec/error-handling/` - PanicOnError
  default, when to return errors, MCP exception

## Common situations

### Fixing lint issues

Separate obvious fixes from judgment calls.
Discuss ambiguous ones before changing, grouped by the question they
raise rather than by file or count. See the `/soil:lint` skill for
the full posture on what not to do.

### A check refuses new code

Change the code to conform, even when
that means restructuring or moving it to another repository. Widening
what the check accepts is a design decision raised as a choice before
anything is written, never a change riding inside a feature. An exemption that
holds names a boundary with a mechanical cause, the way a sibling check
already draws it; an allowlist for awkward cases is a blind spot.
Moving vocabulary out of the check's reach so it stops firing is the
same blind spot - when constants collide, add more constants with the
same value and different names (`conventions/constants.md`).

### Extracting a constant

`string_constant` matches by value, and a
constant reaches every file under the directory owning its
`<path>/constant/` package - one directly under `pkg` reaches the
whole repository. Grep the value inside that subtree before extracting. The
findings converge over two or three rounds, and they show in
goanalyze while a scoped golint reports nothing, so check with
goanalyze between rounds. Findings in files you never opened, inside
that subtree, are yours.

### What the fixers and analyzers do to new code

- gofix removes an import alias where nothing collides and renames
  `err*` helpers; a file named after the old identifier then fails
  goanalyze, so rename the file with it
- golint sorts imports. A body left with an empty import block becomes
  `import ()` or a file that no longer parses, and a half-merged `var`
  block is left for gofix to settle - grep for `^import ()` after it
  runs on generated files
- goanalyze keeps `<path>/constant/` free of behaviour and each package
  to one receiver type: a value check goes in the service over constant
  slices. A test value equal to any constant's value collides under
  `string_constant`, and goaudit allows `<path>/constant/` only directly
  under `pkg/<domain>` or `pkg/tool/<name>`, so a test package's values
  stay literals
- Against another repository, golint, gofix and goanalyze take
  `--root` and goaudit takes the root as an argument. `goaudit --daemon`
  (or `daemon: true`) also arms `missing_daemon`, listing the daemons
  not yet on `<path>/daemon/`

### Changing a check

Fix every finding a changed or new analyzer
raises before installing its binary, because `task lint` runs the
installed binary - new source beside an old binary shows you the
findings and nobody else. Install only at zero in every repository
the check reaches, including one that reaches it through a `replace`
directive. A new analyzer goes the same way: build it, clear it,
then register it.

Teach a check only what the live tree needs: before extending it for a
suspected miss, show the miss exists with a break or a count, because an
extension that moves no finding is code nothing needs. And when tightening
a check floods findings of one shape, read what the loose rule was standing
in for before reverting - the flood is often a gap the old rule hid, not a
fault in the new one.

### After a large refactor

Run `task lint` to catch naming drift,
unused imports, spacing violations. Run `goaudit` to catch structural
drift.

After a repo-wide auto-fix, run `go vet ./...` or the tests: `go build ./...`
skips test files, so a fix that breaks one builds green.
