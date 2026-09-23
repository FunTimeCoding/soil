# Lint

## Invocation

```
task lint
```

Runs the repository's lint pipeline — golint, golangci-lint, gofix,
goanalyze and goaudit, in that order.
Each tool's first line names the repository it scanned; after that
the pipeline is silent on success - any further output is a failure.
Lint has multiple stages - clearing the first may reveal new
issues from the next.

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

**Fixing lint issues**: separate obvious fixes from judgment calls.
Discuss ambiguous ones before changing. See the `/soil:lint` skill
for the full posture on what not to do.

**After a large refactor**: run `task lint` to catch naming drift,
unused imports, spacing violations. Run `goaudit` to catch structural
drift.
