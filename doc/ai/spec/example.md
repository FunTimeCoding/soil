# Example Spec

Conventions for example code that demonstrates library usage.

## Location

Example functions live in `pkg/<domain>/example/` subpackages. Entry points that run them
live in `cmd/example/<name>/main.go`.

```
pkg/gitlab/example/runner.go        → func Runner()
cmd/example/gitlab/main.go          → calls example.Runner()
```

## Single-Function Examples

One exported function per file. File name matches function name in snake_case.

```go
// pkg/gitlab/example/runner.go
package example

func Runner() {
    g := gitlab.NewEnvironment()
    // ...
}
```

## Multi-Function Examples

When an example needs private helper functions, use a subdirectory. The exported entry
point and each helper get their own file.

```
pkg/kubernetes/example/node_check/
    node_check.go           → func NodeCheck()
    print_events.go         → func printEvents()
    filter_pods.go          → func filterPods()
```

Entry points are exported. Helpers are private.

## cmd/example/ Entry Points

Each `cmd/example/<name>/main.go` is a **single file**. No helper files, no splits.
`go run main.go` only compiles the one file passed to it, so all code must be reachable
from that file via imports.

```go
package main

import "github.com/funtimecoding/soil/pkg/gitlab/example"

func main() {
    example.Runner()
}
```

### Disabled Scenarios

The live example runs at the top level. Disabled examples go in a single `if false`
block at the bottom.

```go
func main() {
    example.Runner()

    if false {
        example.MergeRequest()
        example.Project()
    }
}
```

## What Not to Do

- Do not split `cmd/example/` into multiple `.go` files - breaks `go run main.go`
- Do not put private helpers in `cmd/example/` - move them to `pkg/<domain>/example/`
- Do not put complex logic directly in `cmd/example/<name>/main.go` - extract to a `pkg` example function
