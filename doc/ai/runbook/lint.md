# Lint

## Invocation

```
task lint
```

Runs the repository's lint pipeline — golint and golangci-lint
everywhere, gofix and goanalyze where the taskfile includes them.
Must be completely silent on success - any output is a failure.
Lint has multiple stages - clearing the first may reveal new
issues from the next.

## Relevant specs

Read these before fixing lint issues or adding analyzers, relative
to the plugin root:

- `doc/ai/spec/conventions.md` - naming, formatting,
  import ordering, file structure rules that lint enforces
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
