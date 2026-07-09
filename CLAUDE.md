# Claude Code

## Specs

Design and coding specs live in `doc/ai/spec/`:

- `conventions.md` - coding style, error handling, naming, structure
- `naming.md` - banned identifier segments, replacement patterns, type/field naming rules
- `entrypoint.md` - linker variables, Main(), sentry setup (shared by all cmd/ programs)
- `service-tool.md` - long-running service tool pattern (Run, lifecycle wiring, routes)
- `check-tool.md` - CLI check tool pattern (fetch, filter, format, print)
- `lifecycle.md` - lifecycle manager for servers and workers
- `entity-wrapper.md` - entity wrapper pattern
- `console-status.md` - fluent status line builder
- `database.md` - service persistence (storage paths, sqlite/postgres openers, lite-as-default selection)
- `testing.md` - integration testing patterns (mocks, lifecycle HTTP, store)
- `build.md` - gobuild cross-compilation and linker variable convention
- `taskfile.md` - task runner, git hooks, CI pipeline
- `locator.md` - fluent URL builder (`pkg/web/locator`)
- `generated-api.md` - OpenAPI codegen pattern (generated/, client/, server/ structure)
- `error-handling/` - error handling strategies (README for principle + strategy, leaves: mcp, rest, external-api, infrastructure)
- `pillars.md` - unified service wiring: reporter, logger, recovery, telemetry
- `stdlib-wrappers.md` - PanicOnError wrappers for filesystem, JSON, and time stdlib calls
- `example.md` - example code conventions (pkg/<domain>/example/, cmd/example/)
- `fixture.md` - test fixture patterns (fixture/ at repo root, helpers, constants)
- `model-context.md` - MCP tool exposure pattern (model_context/ subpackage, HTTP transport)
- `multi-instance.md` - multi-backend services (inventory, service layer, instance resolution)
- `package-design.md` - one struct with receivers per package rule (when to extract sub-packages)
- `package-taxonomy.md` - package roles, dependency direction, and promotion criteria for growing service tools

Read the relevant spec before working in that area.

## Project Structure

- `cmd/` - service entry points
- `pkg/` - library and service packages
- `doc/ai/spec/` - design and coding specs
- `.claude-plugin/` - Claude Code plugin manifest and marketplace catalog
- `skills/` - plugin skills, invoked as `/soil:<name>` from consuming repositories
- `doc/ai/runbook/` - operational runbooks backing the skills

## Claude Code plugin

This repository is a Claude Code plugin named `soil` and its own
marketplace. General-purpose skills (session lifecycle, lint, commit
messages) live here; skills specific to VirtualMinds infrastructure
stay in the internal repositories. Skills reference runbooks via
`${CLAUDE_PLUGIN_ROOT}` so paths resolve from any consuming repository.

- `/plugin marketplace add <path-or-url-to-this-repo>` registers it
- `/plugin install soil@soil` installs the plugin
- `/reload-plugins` picks up local edits during development
