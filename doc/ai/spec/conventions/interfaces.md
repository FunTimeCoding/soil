# Interfaces

- **Generic interfaces** (no domain type imports) - define in
  `pkg/face/`. These are like `io.Reader`: small, shared, used by
  many consumers. Examples: `Worker`, `Notifier`, `Clock`.
- **Domain-specific interfaces** (reference types like
  `*alert.Alert`) - define at the consumer. Domain types often
  already import `face`, so putting the interface in `face` would
  create an import cycle. Example: `AlertSource` in
  `worker/worker.go`.
- **Mock convention** - hand-rolled structs in `mock_*` sub-packages
  near the real implementation (e.g., `mock_notifier/`,
  `mock_client/`). One file per method. Methods to manipulate state
  for testing (e.g., `Add()`, `Remove()`).
