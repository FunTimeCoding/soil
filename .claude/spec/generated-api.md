# Generated API

Pattern for typed HTTP APIs using `oapi-codegen`. Use when a service
exposes an HTTP API that needs a generated client (or when type safety
on the server side is worth the codegen overhead).

## Package Structure

```
pkg/tool/go<tool>d/
├── generated/
│   ├── client/
│   │   ├── config.yaml    # oapi-codegen client config
│   │   └── generated.go   # Generated: Client, typed request/response methods
│   └── server/
│       ├── openapi.yaml   # Hand-written spec
│       ├── config.yaml    # oapi-codegen server config
│       └── generated.go   # Generated: ServerInterface, HandlerFromMux, types
├── client/                # Domain wrapper for CLI (imports generated/client/)
│   ├── client.go          # Client struct (wraps generated client)
│   ├── new.go             # New(host string) *Client
│   └── <operation>.go     # One file per operation
├── server/                # REST implementation (imports generated/server/)
│   ├── server.go          # Server struct (holds dependencies)
│   ├── new.go             # New(dep) *Server
│   └── <operation>.go     # One file per endpoint (post_deploy.go, get_status.go)
├── model_context/         # MCP implementation
├── convert/               # Type filtering shared by model_context/ and server/
├── constant/
├── option/
├── main.go
└── run.go
```

Everything under `generated/` is machine output - don't edit it. `client/`
and `server/` are hand-written code that consumes the generated types.

## Configs

`generated/server/config.yaml`:
```yaml
package: server
generate:
  std-http-server: true
  strict-server: true
  models: true
  embedded-spec: true
output: generated.go
```

`generated/client/config.yaml`:
```yaml
package: client
generate:
  client: true
  models: true
output: generated.go
```

## Generation

Run via taskfile:
```yaml
generate:
  cmds:
  - cd pkg/<service>/generated/client && oapi-codegen --config config.yaml ../server/openapi.yaml
  - cd pkg/<service>/generated/server && oapi-codegen --config config.yaml openapi.yaml
```

No lint exclusions needed - the lint tools skip generated files
automatically by detecting the `Code generated ... DO NOT EDIT.`
header.

## Implementing the Server

`server/server.go` - plain struct, no embedding:
```go
type Server struct {
    store    *store.Store
    reporter face.Reporter
}
```

`server/<operation>.go` - implements the generated
`StrictServerInterface` method. Typed request objects in, typed
response objects out, no `http.ResponseWriter`:

```go
func (s *Server) PostDeploy(
    _ context.Context,
    r server.PostDeployRequestObject,
) (server.PostDeployResponseObject, error) {
    result, e := s.store.TriggerTargets(r.Body.Targets)

    if e != nil {
        return server.PostDeploy500JSONResponse(
            *s.captureFail(e, constant.UnexpectedError),
        ), nil
    }

    return server.PostDeploy200JSONResponse{
        Tag: result,
    }, nil
}
```

The framework handles Content-Type headers, status codes, and JSON
serialization. No `web.EncodeNotation` or manual `w.WriteHeader`
needed.

When a server operation uses converters shared with `model_context/`,
import the `convert/` package:
```go
func (s *Server) GetIssue(
    _ context.Context,
    r server.GetIssueRequestObject,
) (server.GetIssueResponseObject, error) {
    return server.GetIssue200JSONResponse(
        convert.JiraIssue(s.store.Issue(r.Key)),
    ), nil
}
```

## CLI Access

Multi-command CLIs call the generated client directly. `Main()`
constructs it once and hands it to every subcommand:

```go
v, e := client.NewClient(
    locator.New(
        environment.Fallback(constant.HostEnvironment, web.Localhost),
    ).Port(web.ListenPort).Insecure().String(),
)
errors.PanicOnError(e)
o.AddCommand(listItems(v))
```

Each subcommand builds its typed request and prints the response - a
hand-written pass-through method per operation would add code without
behavior. Don't write wrapper methods just to hide generated types
from the CLI.

## Client Wrapper

The `client/` package wraps the generated REST client with a clean
interface independent of generated types. This is the internal REST
client - it talks to our own daemon, not to a third-party service.
It earns its existence when there is a consumer the raw generated
client doesn't serve well:

- **Daemon-to-daemon calls** - a service layer in one daemon calling
  another daemon wants domain methods, not request objects
- **Test mocks** - a `mock_client/` next to the wrapper lets
  consumers test without a live daemon
- **Domain formatting** - when responses should flow back into
  domain entities with `Format()` support (the daemon bridge
  factory pattern - see `entity-wrapper.md`)

`client/client.go`:
```go
type Client struct {
    context context.Context
    client  *generated.ClientWithResponses
}
```

`client/new.go`:
```go
func New(host string) *Client {
    c, e := generated.NewClientWithResponses(locator.New(host).String())
    errors.PanicOnError(e)
    return &Client{context: context.Background(), client: c}
}
```

`client/<operation>.go`:
```go
func (c *Client) Alerts() string {
    result, e := c.client.GetAlerts(c.context, &generated.GetAlertsParams{})
    errors.PanicOnError(e)
    return web.ReadString(result)
}
```

For services with a third-party upstream, the external API client lives
at `pkg/<domain>/` (see `service-tool.md`). The `client/` wrapper and
the external API client serve different consumers - `client/` talks to
our own daemon, `pkg/<domain>/` is for the daemon (and MCP-only
services) talking to the upstream.

## Wiring into run.go

The generated server package is aliased `generated`. Wrap the
server in `NewStrictHandler` and pass to `HandlerFromMux`:

```go
import (
    generated "github.com/funtimecoding/go-library/pkg/tool/go<tool>d/generated/server"
)

lifecycle.WithServer(
    server.New(
        web.AddressPort(o.Port),
        func(m *http.ServeMux) {
            generated.HandlerFromMux(
                generated.NewStrictHandler(server.New(dep, r), nil),
                m,
            )
        },
    ).WithMiddleware(web.RecoveryMiddleware(r)),
)
```

When MCP is also mounted, add the `model_context` import and call
`Mount(m)`:

```go
import (
    "github.com/funtimecoding/go-library/pkg/tool/go<tool>d/model_context"
    generated "github.com/funtimecoding/go-library/pkg/tool/go<tool>d/generated/server"
)

lifecycle.WithServer(
    server.New(
        web.AddressPort(o.Port),
        func(m *http.ServeMux) {
            generated.HandlerFromMux(
                generated.NewStrictHandler(server.New(dep, r), nil),
                m,
            )
            model_context.New(dep, r).Mount(m)
        },
    ).WithMiddleware(web.RecoveryMiddleware(r)),
)
```

REST routes (`/api/...`) and MCP routes (`/mcp`, `/sse`, `/message`)
don't conflict on the same mux.

## OpenAPI Spec Patterns

Optional arrays of objects generate `*[]*Type` when the items are
nullable. Use `nullable: true` with `allOf` wrapping:

```yaml
checklist:
  # Optional: only present on certain task types.
  type: array
  items:
    nullable: true
    allOf:
    - $ref: "#/components/schemas/ChecklistItem"
```

Without `nullable: true` on items, oapi-codegen generates
`*[]Type` (pointer to slice of values). The nullable pattern
keeps pointer convention consistent between the generated types
and the `convert/` layer.

## What Not To Do

- Don't put hand-written code in `generated/` - that's machine output
- Don't edit `generated.go` - regenerate from the spec instead
- Don't write a `client/` wrapper just to hide generated types from
  the CLI - wrappers exist for daemon-to-daemon consumers, mocks,
  and domain formatting
- Don't manually register routes that the spec already defines -
  let `HandlerFromMux` do it
