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
├── integration/
│   └── guard/             # Guard battery over Mount - see testing.md
├── main.go
├── mount.go               # Mount(deps..., g *guard.Mux) - the served surface
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
constructs it once and hands it to every subcommand. The locator
resolves through the tool's environment family via
`locator.Environment` - `GO<TOOL>_HOST` falling back to localhost,
`GO<TOOL>_PORT` optional, `GO<TOOL>_INSECURE` switching to http -
and every client authenticates with the daemon's token
(`GO<TOOL>_TOKEN`, the `TokenEnvironment` constant beside the
others) via `web.BearerEditor`:

```go
v, e := client.NewClient(
    locator.Environment(
        constant.HostEnvironment,
        constant.PortEnvironment,
        constant.InsecureEnvironment,
    ).String(),
    client.WithRequestEditorFn(
        web.BearerEditor(environment.Required(constant.TokenEnvironment)),
    ),
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

`client/new.go` - takes the daemon locator and token; the
`NewEnvironment()` wrapper builds both from the tool's env family:
```go
func New(
    l *locator.Locator,
    token string,
) *Client {
    c, e := generated.NewClientWithResponses(
        l.String(),
        generated.WithRequestEditorFn(web.BearerEditor(token)),
    )
    errors.PanicOnError(e)
    return &Client{context: context.Background(), client: c}
}
```

`client/new_environment.go`:
```go
func NewEnvironment() *Client {
    return New(
        locator.Environment(
            constant.HostEnvironment,
            constant.PortEnvironment,
            constant.InsecureEnvironment,
        ),
        environment.Required(constant.TokenEnvironment),
    )
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

## Mounting

Every guarded daemon has a top-level `mount.go` beside `run.go`.
`Mount(...)` takes the domain dependencies (through the daemon's
`face/` interfaces where they exist), the reporter, the recorder,
the version, and a `*guard.Mux` last. It wraps the server in
`NewStrictHandler` with the generic recording middleware, builds
the tree on a fresh sub-mux via `HandlerFromMux`, and token-mounts
it at the `/api/` prefix (`webConstant.InterfacePath`):

```go
import (
    generated "github.com/funtimecoding/soil/pkg/tool/go<tool>d/generated/server"
    "github.com/funtimecoding/soil/pkg/tool/go<tool>d/model_context"
    "github.com/funtimecoding/soil/pkg/web"
    webConstant "github.com/funtimecoding/soil/pkg/web/constant"
    "github.com/funtimecoding/soil/pkg/web/guard"
)

func Mount(
    s *store.Store,
    r face.Reporter,
    t face.Recorder,
    version string,
    g *guard.Mux,
) {
    g.TokenMount(
        webConstant.InterfacePath,
        generated.HandlerFromMux(
            generated.NewStrictHandler(
                server.New(s, r),
                []generated.StrictMiddlewareFunc{
                    web.RecordingMiddleware[generated.StrictHandlerFunc](t),
                },
            ),
            http.NewServeMux(),
        ),
    )
    model_context.New(s, r, t, version).Mount(g)
}
```

`web.RecordingMiddleware` records baseline telemetry per
operationID via `web.RecordTelemetry` - the type parameter lets
the one helper serve every generated package's own
`StrictHandlerFunc` type. Telemetry posture and wiring order live
in `pillars.md`.

Daemons without MCP skip the `model_context` line (see
`model-context.md` for the transport routes); web-carrying daemons
add `u.Mount(g)` for their HTML surface. REST routes (`/api/...`),
MCP routes (`/mcp`, `/sse`, `/message`), and web routes don't
conflict on the same mux.

run.go builds exactly one `guard.Mux` per lifecycle server and
hands it to `Mount` - service tokens thread main - option - run
via `web.ServiceTokens()`:

```go
lifecycle.WithServer(
    server.New(
        constant.Identity,
        o.Address,
        func(m *http.ServeMux) {
            Mount(s, r, i.Recorder(), o.Version, guard.New(m, o.ServiceTokens))
        },
    ).WithMiddleware(web.RecoveryMiddleware(r)),
)
```

`Mount` is the production surface the guard battery starts (see
`testing.md`). Everything the daemon serves belongs inside it - a
route registered directly in the run.go callback escapes both the
guard and the battery.

API paths are unversioned: `/api/<resource>`, never `/api/v1/...`.
APIs here break and roll forward rather than maintain versions, so
a version segment would suggest a guarantee nobody keeps. Enforced
by goaudit (`versioned_path`).

## OpenAPI Spec Patterns

Every spec includes two error schemas, defined inline on each
endpoint (not via `components/responses/` - those generate
struct embeddings instead of direct type aliases, bug #1864):

```yaml
Error:
  type: object
  required: [error]
  properties:
    error:
      type: string
ErrorResponse:
  type: object
  required: [error, event_identifier]
  properties:
    error:
      type: string
    event_identifier:
      type: string
```

`Error` carries tier 1 rejections (400/404, no Sentry event ID),
`ErrorResponse` carries tier 2/3 failures (500, with the event
ID) - the tiers and `captureFail` live in
`error-handling/rest.md`.

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

Every non-2xx response carries a content body with a descriptive
message - no bodyless error responses. Specs don't document
framework-level 400s: oapi-codegen handles binding validation,
and documenting it is the framework's job.

Root-level `additionalProperties: true` breaks strict-server
marshaling (the type alias loses its custom MarshalJSON) - use a
`fields` sub-object instead.

## What Not To Do

- Don't put hand-written code in `generated/` - that's machine output
- Don't edit `generated.go` - regenerate from the spec instead
- Don't write a `client/` wrapper just to hide generated types from
  the CLI - wrappers exist for daemon-to-daemon consumers, mocks,
  and domain formatting
- Don't manually register routes that the spec already defines -
  let `HandlerFromMux` do it
- Don't register routes in the run.go callback - everything the
  daemon serves mounts through `Mount()`
