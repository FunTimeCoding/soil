# Interfaces

- **Generic interfaces** (no domain type imports) - define in
  `pkg/face/`. These are like `io.Reader`: small, shared, used by
  many consumers. Examples: `Worker`, `Notifier`, `Clock`.
- **Domain-specific interfaces** (reference types like
  `*alert.Alert`) - define in `<domain>/face/` inside their domain,
  never in the package that consumes them. Domain types often already
  import `pkg/face`, so putting the interface there would create an
  import cycle. Examples: `AlertSource` in
  `go:pkg/prometheus/face`, `GitLabClient` in `go:pkg/gitlab/face`,
  and a daemon's client faces in `<path>/face/` (`service-tool.md`).
- **Upstream provider faces** - one `<domain>/face/` per domain, each
  interface named for its provider: `JiraClient` and `ConfluenceClient`
  in `go:pkg/atlassian/face`, `PostgresClient` in `go:pkg/postgres/face`. A
  provider package keeps no face of its own; its mock stays beside the
  client (`<path>/mock_client/`). A face split into parts keeps the parts'
  names in a one-provider face and prefixes them with the provider in a
  shared one. A daemon's client keeps `face.Client` in `<path>/face/`
  (`generated-api.md`). A face has no escape to the client's internals
  (`Basic()`, `Nested()`): each thing a consumer reached through one
  becomes a method on the client.
- **Mock convention** - hand-rolled structs in `mock_*` sub-packages
  near the real implementation (e.g., `<path>/mock_notifier/`,
  `<path>/mock_client/`). One file per method. Methods to manipulate state
  for testing (e.g., `Add()`, `Remove()`).
