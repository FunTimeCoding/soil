# Model Context Protocol Integration

Pattern for exposing MCP tools from a service daemon, mounted on the same HTTP server as the REST API.

## When to Use

Add a `model_context/` subpackage when a daemon needs to expose MCP tools alongside its REST API. Tools typically mirror the read-oriented operations already available via REST endpoints.

## Package Structure

```
pkg/tool/go<tool>d/
├── constant/
│   └── constant.go            # Tool name + tool-specific parameter constants
├── convert/
│   └── <type>.go              # One file per converter function
├── model_context/
│   ├── server.go              # Server struct (holds MCPServer + dependencies)
│   ├── new.go                 # New(deps, r face.Reporter) *Server
│   ├── mount.go               # Mount(g *guard.Mux) - HTTP transport setup
│   ├── register.go            # register() - wires all tools to the MCPServer
│   ├── capture_fail.go        # captureFail(e error, message string)
│   └── <tool_name>.go         # One file per tool function (get_alerts.go, etc.)
└── server/
```

`model_context/` is a sibling of `server/`. Both implement the same
domain operations - `server/` over REST, `model_context/` over MCP.

`model_context/` is the standard package name. Do not use `tool/` or
`toolset/`.

## Convert Package

When domain objects are complex (nested fields, changelogs, computed
properties), use a `convert/` package to produce slim response types:

```
pkg/tool/go<tool>d/convert/
├── jira_issue.go              # JiraIssue(*issue.Issue) *generated.JiraIssue
├── jira_issues.go             # JiraIssues([]*issue.Issue) []*generated.JiraIssue
└── ...
```

Converter functions are exported and shared by both `server/` and
`model_context/`. This prevents MCP tools from leaking raw structs
into the LLM context.

For simple flat responses (3-5 fields), inline `map[string]any`
construction in the handler is acceptable. The convert package adds
value when types are nested, shared between REST and MCP, or when
filtering out sensitive fields.

## Service Layer

When a daemon has compound operations (fetch-then-merge edits,
name resolution, multi-step deletes), extract a `service/` package
that both `model_context/` and `server/` call through:

```
pkg/tool/go<tool>d/
├── service/
│   ├── service.go              # holds the upstream client
│   ├── edit_link.go            # fetch existing, merge, update
│   ├── search.go               # cross-entity search
│   └── resolve_list.go         # name-or-ID resolution
├── model_context/              # MCP handlers call service
└── server/                     # REST handlers call service
```

The service accepts domain language (names, not IDs) and resolves
internally. Both surfaces stay thin - they parse input, call the
service, format the response. The service owns the business logic.

Use a service layer when:
- An operation requires multiple upstream calls (fetch + modify)
- Name-to-ID resolution is needed by both MCP and REST
- The upstream API has quirks that callers shouldn't know about
  (e.g. PATCH requires fields that aren't changing)

Don't add a service layer for simple CRUD - direct client calls
from handlers are fine until compound logic appears.

## Constants

Tool name and tool-specific parameter constants live in `pkg/tool/go<tool>d/constant/constant.go`:

```go
const (
    GetAlerts    = "get_alerts"
    GetStatus    = "get_status"
)

const (
    IncludeResolved = "include_resolved"
)
```

Reusable parameter names shared across multiple MCP tools (`query`, `limit`, `key`, `body`, `identifier`, `title`, `message`) live in `pkg/generative/constant/model_context.go` and are referenced as `constant.ParameterQuery`, `constant.ParameterLimit`, etc.

## Server Struct

`server.go`:
```go
type Server struct {
    server   *server.MCPServer
    store    *store.Store    // omit if no store
    reporter face.Reporter
    worker   *worker.Worker  // omit if no worker
}
```

`new.go` - create the MCPServer via the factory, register tools, return:
```go
func New(s *store.Store, r face.Reporter, t face.Recorder, w *worker.Worker, version string) *Server {
    result := &Server{
        server: mark.New(constant.Identity, version).WithRecorder(t).Server(),
        store:    s,
        reporter: r,
        worker:   w,
    }
    result.register()

    return result
}
```

The factory (`mark/server`) handles tool capabilities, instructions,
and baseline telemetry hooks. `WithRecorder(t)` registers an
AfterCallTool hook that records every MCP tool call as a baseline
telemetry event. Import as
`mark "github.com/funtimecoding/soil/pkg/generative/mark/server"`.
```

`mount.go` - wires the MCP server onto the guard (the shared
`pkg/generative/model_context/server` package registers every MCP
transport route token-checked through it):
```go
func (s *Server) Mount(g *guard.Mux) {
    server.New(s.server).Setup(g)
}
```

## Registering Tools

`register.go` - one `AddTool` call per tool, function in its own file.
Tool names are constants from the service's `constant/` package.
Parameter names use string literals for clarity or shared constants
from the `Parameter*` family in `pkg/generative/constant/` when the name is
reusable across services (query, limit, identifier, name, etc.):

```go
func (s *Server) register() {
    s.server.AddTool(
        mcp.NewTool(
            constant.ListLinks,
            mcp.WithDescription("List links."),
            mcp.WithNumber(
                "collection",
                mcp.Description("Collection ID to filter by."),
            ),
            mcp.WithNumber(
                "limit",
                mcp.Description("Maximum number of results to return."),
            ),
        ),
        mcp.NewTypedToolHandler(s.ListLinks),
    )
}
```

## Typed Tool Handlers

The preferred pattern uses `mcp.NewTypedToolHandler` with an
`argument/` struct. The struct's JSON tags must match the parameter
names in the tool definition:

```
pkg/tool/go<tool>d/model_context/
├── argument/
│   ├── list_links.go          # type ListLinks struct
│   ├── search_links.go        # type SearchLinks struct
│   └── add_link.go            # type AddLink struct
```

```go
package argument

type ListLinks struct {
    Collection float64 `json:"collection"`
    Limit      float64 `json:"limit"`
    Offset     float64 `json:"offset"`
}
```

Handler signature with typed arguments:
```go
func (s *Server) ListLinks(
    _ context.Context,
    _ mcp.CallToolRequest,
    a argument.ListLinks,
) (*mcp.CallToolResult, error) {
    links, e := s.client.Links(int32(a.Collection))

    if e != nil {
        return s.captureFail(e, "failed to list links")
    }

    return response.SuccessAny(map[string]any{"links": result})
}
```

MCP numbers arrive as `float64`. Cast to `int32` or `int` at the
call site. Tools without parameters use `struct{}` as the argument type.

## Legacy Request Accessors

Older services use the untyped pattern with request
accessors and a validator:

```go
func (s *Server) ChangeField(
    _ context.Context,
    r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
    v := validator.New(&r)
    name := v.RequireString(constant.Name.Key)

    if v.HasErrors() {
        return response.Fail("%s", v.Error())
    }
    // ...
}
```

This works but the typed handler pattern is preferred for new tools.

## Response Conventions

Tool function files use `response.Fail` for input validation,
`captureFail` for infrastructure errors, `response.SuccessAny` for
JSON results, and `response.Success` for plain text results:

```go
func (s *Server) getAlerts(
    _ context.Context,
    _ mcp.CallToolRequest,
    a argument.GetAlerts,
) (*mcp.CallToolResult, error) {
    result, e := s.store.ByName(a.Name, int(a.Limit))

    if e != nil {
        return s.captureFail(e, "get alerts failed")
    }

    return response.SuccessAny(convert.Alerts(result))
}
```

`capture_fail.go`:
```go
func (s *Server) captureFail(
    e error,
    message string,
) (*mcp.CallToolResult, error) {
    return response.CaptureFail(s.reporter, e, message)
}
```

- Required params: `r.RequireString(...)` / `r.RequireFloat(...)` - return `response.Fail(...)` on failure
- Optional params: `r.GetString(constant.Key, "")` / `r.GetFloat(constant.Key, 0)` / `r.GetBool(constant.Key, false)`
- JSON results: `response.SuccessAny(converted)` - serializes via `notation.MarshalIndent`
- Text results: `response.Success("comment added")`
- Validation errors: `response.Fail("message: %v", err)` - wraps `mcp.NewToolResultError`
- Infrastructure errors: `s.captureFail(e, "message")` - captures to Sentry with event ID
- Error variables progress `e`, `f`, `g`, `h`, `i` - never reuse the same letter. See `naming.md`.
- Always convert results through the `convert/` package - never serialize raw domain objects
- Error handling is two-tier - input validation vs infrastructure failures. See `error-handling/mcp.md`.

## REST Strict Server

Services with REST APIs use oapi-codegen's strict server mode.
The full pattern - configs, handlers, error schemas, recording
middleware, mounting - lives in `generated-api.md`; the error
tiers and REST `captureFail` in `error-handling/rest.md`. MCP and
REST share the `convert/` layer, the reporter, and the recorder.

## Wiring into mount.go

MCP mounts through the daemon's top-level `Mount(...)` beside the
REST tree, on the same `guard.Mux` (full shape in
`generated-api.md`) - REST routes (`/api/...`) and MCP routes
(`/mcp`, `/sse`, `/message`) don't conflict:

```go
model_context.New(s, r, t, version).Mount(g)
```

The `model_context` package is imported without alias. The
recorder comes from the instrument in `Main()` (`i.Recorder()`,
see `entrypoint.md`) - same recorder for both MCP baseline and
REST baseline.

## What Not To Do

- Don't create a separate lifecycle server for MCP - one port, one mux
- Don't name it `mcp/` - use `model_context/` (no acronyms in package names)
- Don't name it `tool/` or `toolset/` - use `model_context/`
