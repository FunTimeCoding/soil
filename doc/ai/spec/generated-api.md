---
base: pkg/tool/goatlassiand
---

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
├── client/                # The daemon's one client - its CLI and other daemons
│   ├── client.go          # Client struct (wraps the generated client)
│   ├── new.go             # New(locator, token) *Client
│   ├── new_environment.go # NewEnvironment() - the tool's env family
│   ├── reachable.go       # Reachable(x) - the peer probe
│   ├── close.go           # Close() - the client's own pool
│   ├── <operation>.go     # One file per operation
│   └── mock_client/       # face.Client for tests
├── server/                # REST implementation (imports generated/server/)
│   ├── server.go          # Server struct (holds dependencies)
│   ├── new.go             # New(dep) *Server
│   └── <operation>.go     # One file per endpoint (post_deploy.go, get_status.go)
├── model_context/         # MCP implementation
├── convert/               # Type filtering shared by model_context/ and server/
├── constant/
├── option/
├── unit/
│   └── guard_test.go      # Guard battery over Mount - see testing.md
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

`server/server.go` - plain struct, no embedding, holding the daemon's
service as `*service.Service`; the service has no face
(`service-tool.md`, Service):
```go
type Server struct {
    service  *service.Service
    reporter soil.Reporter
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
    result, e := s.service.TriggerTargets(r.Body.Targets)

    if e != nil {
        return server.PostDeploy500JSONResponse(*s.answer(e)), nil
    }

    return server.PostDeploy200JSONResponse{
        Tag: result,
    }, nil
}
```

The framework handles Content-Type headers, status codes, and JSON
serialization. No `web.EncodeNotation` or manual `w.WriteHeader`
needed.

### Shared converters

When a server operation uses converters shared with `model_context/`,
import the `convert/` package:
```go
func (s *Server) GetIssue(
    _ context.Context,
    r server.GetIssueRequestObject,
) (server.GetIssueResponseObject, error) {
    return server.GetIssue200JSONResponse(
        convert.JiraIssue(s.service.Issue(r.Key)),
    ), nil
}
```

## The Daemon Client

A daemon has one client, `client/`, and everything that talks to the
daemon uses it: its own CLI, other daemons, and tests through
`client/mock_client`. A second shape - a CLI on the generated client, a
wrapper that relays raw bodies - is a second client in waiting. The
external API client a daemon uses toward its upstream is a different
thing, at `pkg/<domain>/` (`service-tool.md`).

`client.Client` is the struct, the package's one receiver: it wraps the
generated client on a pool it owns, which `Close` closes. `New(locator,
token)` builds it; `NewEnvironment()` resolves the tool's env family -
`GO<TOOL>_HOST` falling back to localhost, `GO<TOOL>_PORT` optional,
`GO<TOOL>_INSECURE` switching to http, `GO<TOOL>_TOKEN` required - with
`locator.Environment`.

The interface is `face.Client` - the operations, `face.Reachable` and
`Close` - which `client/mock_client` implements too.

Each operation takes a context, calls the `…WithResponse` method, and
returns typed values and a classified error, never a raw body. The plain
generated method fails only on transport, so a 500 or 503 would arrive as
if it were the answer. The operation takes the status it expects and turns
every other into an error with `go:pkg/web/refusal/New`, which reads the
body's reason and makes a 400 a `validation`, a 403 a `forbidden`, a 404 a
`not_found` and a 503 an `unavailable`; a status the caller acts on
differently - a 404 meaning "nothing here yet" - is matched before it. An
operation answering 204 returns only the error:

```go
func (c *Client) Items(x context.Context) ([]generated.Item, error) {
    r, e := c.client.GetItemsWithResponse(x, &generated.GetItemsParams{})

    if e != nil {
        return nil, e
    }

    if r.JSON200 == nil {
        return nil, refusal.New("items", r.HTTPResponse, r.Body)
    }

    return *r.JSON200, nil
}
```

`Reachable` calls `route:/api/reachable` through
`go:pkg/web/reachable_probe/Reach` - the probe a peer's start-up wait and
guard use (`error-handling/infrastructure.md`, The probe route).

The CLI builds the client once in `Main()`, defers its `Close`, and hands
it to every subcommand. A subcommand prints what an operation returned
with `terminal.EmitValue(value, error)` - the value as compact JSON, or
the classified error as its exit - and an operation answering 204 with
`terminal.EmitDone(error)`.

## Mounting

Every guarded daemon has a top-level `mount.go` beside `run.go`.
`Mount(...)` takes the daemon's `*service.Service`, the reporter,
the recorder, the version, and a `*guard.Mux` last. It wraps the server in
`NewStrictHandler` with the generic recording middleware, builds
the tree on a fresh sub-mux via `HandlerFromMux`, and token-mounts
it at the `route:/api/` prefix (`webConstant.InterfacePath`):

```go
import (
    generated "github.com/funtimecoding/soil/pkg/tool/go<tool>d/generated/server"
    "github.com/funtimecoding/soil/pkg/tool/go<tool>d/model_context"
    "github.com/funtimecoding/soil/pkg/web"
    webConstant "github.com/funtimecoding/soil/pkg/web/constant"
    "github.com/funtimecoding/soil/pkg/web/guard"
)

func Mount(
    v *service.Service,
    r soil.Reporter,
    t soil.Recorder,
    version string,
    g *guard.Mux,
) {
    g.TokenMount(
        webConstant.InterfacePath,
        generated.HandlerFromMux(
            generated.NewStrictHandler(
                server.New(v, r),
                []generated.StrictMiddlewareFunc{
                    web.RecordingMiddleware[generated.StrictHandlerFunc](t),
                },
            ),
            http.NewServeMux(),
        ),
    )
    model_context.New(v, r, t, version).Mount(g)
}
```

### Inside Mount

`web.RecordingMiddleware` records baseline telemetry per
operationID via `web.RecordTelemetry` - the type parameter lets
the one helper serve every generated package's own
`StrictHandlerFunc` type. Telemetry posture and wiring order live
in `pillars.md`.

Daemons without MCP skip the `model_context` line (see
`model-context.md` for the transport routes); web-carrying daemons
add `u.Mount(g)` for their HTML surface. REST routes (`route:/api/...`),
MCP routes (`route:/mcp`, `route:/sse`, `route:/message`), and web routes don't
conflict on the same mux.

### The mux in run.go

run.go builds one `guard.Mux` per lifecycle server from the service
tokens `Main` read with `web.ServiceTokens()`, and hands it to `Mount`
(`service-tool.md`, Run Function). `Mount` is the production surface
the guard battery starts (`testing.md`); a route registered in the
run.go callback escapes both the guard and the battery.

## OpenAPI Spec Patterns

### Error schemas

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

`Error` carries an answer Sentry does not hear of, `ErrorResponse` a
captured one with its event ID. Which status takes which schema is the
error's class, in the table in `error-handling/README.md` (Answering
an error); building each answer is `error-handling/rest.md`.

### Nullable arrays

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

### Responses and paths

Every non-2xx response carries a content body with a descriptive
message - no bodyless error responses. Specs don't document
framework-level 400s: oapi-codegen handles binding validation,
and documenting it is the framework's job.

API paths are unversioned: `route:/api/<resource>`, never
`route:/api/<version>/...`. APIs here break and roll forward rather than
maintain versions, so a version segment would suggest a guarantee nobody keeps.
Enforced by goaudit (`versioned_path`).

Root-level `additionalProperties: true` breaks strict-server
marshaling (the type alias loses its custom MarshalJSON) - use a
`fields` sub-object instead.

## What Not To Do

- Don't put hand-written code in `generated/` - that's machine output
- Don't edit `generated.go` - regenerate from the spec instead
- Don't call the generated client outside `client/` - the CLI and other
  daemons go through the daemon client
- Don't manually register routes that the spec already defines -
  let `HandlerFromMux` do it
- Don't register routes in the run.go callback - everything the
  daemon serves mounts through `Mount()`
