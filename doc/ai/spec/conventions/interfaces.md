# Interfaces

- **Generic interfaces** (no domain type imports) - define in
  `pkg/face/`. These are like `io.Reader`: small, shared, used by
  many consumers. Examples: `Worker`, `Notifier`, `Clock`.
- **Domain-specific interfaces** (reference types like
  `*alert.Alert`) - define in `<domain>/face/` inside their domain,
  never in the package that consumes them. Domain types often already
  import `pkg/face`, so putting the interface there would create an
  import cycle. Examples: `AlertSource` in
  `go:pkg/prometheus/alertmanager/face`, `Forge` in `go:pkg/gitlab/face`,
  and a daemon's client faces in `<path>/face/` (`service-tool.md`).
- **Mock convention** - hand-rolled structs in `mock_*` sub-packages
  near the real implementation (e.g., `<path>/mock_notifier/`,
  `<path>/mock_client/`). One file per method. Methods to manipulate state
  for testing (e.g., `Add()`, `Remove()`).
