---
base: pkg/tool/goatlassiand
---

# Service Tool Spec

Pattern for long-running service tools that use lifecycle for HTTP servers and
background workers.

## Package Structure

```
cmd/go<tool>/
└── main.go                         # delegates to Main()

pkg/tool/go<tool>/
├── main.go                         # Main(): register flags, parse, build option, call Run()
├── run.go                          # Run(o, i): wiring, lifecycle, signal block
├── mount.go                        # Mount(deps..., g *guard.Mux) - the served surface
├── option/
│   ├── <name>.go                   # Option struct (named after tool/domain, not "Option")
│   └── new.go                      # Factory: New() *<Name>
├── unit/
│   ├── option_test.go              # assert.NotNil(t, option.New()) - see test-placement.md
│   ├── guard_test.go               # Guard battery over Mount - see testing.md
│   └── service_tester/new.go       # Builds the service from an option with mocks - see testing.md
├── constant/
│   └── constant.go                 # Tool constants (bucket names, identity)
├── service/                        # The daemon's single owner - see Service
│   ├── service.go                  # Service struct
│   ├── new.go                      # New(o, ...) - builds everything from configuration
│   ├── close.go                    # Close() - closes everything it built
│   └── <operation>.go              # One file per operation the surfaces call
├── face/                           # One full interface per client - see Service
├── migrate/                        # AutoMigrate(d) - schema history, see database.md
├── store/                          # Persistence (if needed)
│   ├── store.go                    # Store struct
│   ├── record.go                   # Record type
│   ├── new.go                      # Constructor: takes the mapper it is given
│   └── <operation>.go              # One file per operation (save.go, resolve.go, by_name.go)
├── worker/                         # Background worker (if needed)
│   ├── worker.go                   # Worker struct
│   ├── new.go                      # Constructor
│   ├── start.go                    # Start() - goroutine with ticker
│   ├── close.go                    # Close()
│   └── poll.go                     # Poll() - single cycle logic
└── server/                         # REST implementation (or manual routes)
    ├── <operation>.go              # Route function returning http.HandlerFunc
    └── response.go                 # Response struct (if needed)
```

## Entry Point

See `entrypoint.md` for `main`, `Main()`, and sentry setup.

Every serving daemon registers the common web parameters through the
argument instance - one call, same everywhere:

```go
a := argument.NewInstance(constant.Identity)
a.Web()
a.Parse()
o := option.New()
o.Address = a.Address()
Run(o)
```

`a.Web()` registers `--port` and `--bind-address`. Each resolves
through the chain flag > environment > default: `--port` falls back
to `PORT` then 8080, `--bind-address` falls back to `BIND_ADDRESS`
then `127.0.0.1`. The chain is implemented by computing the flag
default through `environment.Fallback` - `--help` therefore shows
the live resolved default.

- Local daemons are loopback-only by default; each gets a unique
  `--port` in the Procfile
- Kubernetes deployments set `BIND_ADDRESS: '0.0.0.0'` in their
  environment ConfigMap and keep the port default
- A local daemon that serves remote clients on other hosts passes
  `--bind-address 0.0.0.0` in its launch entry
- The option struct carries `Address string` - the joined listen
  string from `a.Address()` - not a bare port

Daemons that expose a Prometheus exporter also call `a.Metric()` -
registers `--metric-port` with chain `METRIC_PORT` > 9090. The
option struct carries `MetricAddress string` from
`a.MetricAddress()`; `Run` mounts the exporter on its own lifecycle
server at that address, binding all interfaces. Locally co-hosted
exporter daemons each get a unique `--metric-port` in the Procfile,
like `--port`.

Services with a database also call `a.Database()` (registers `--lite` with chain
`LITE_PATH` > `Identity.LitePath()`, and `--postgres` with chain
`POSTGRES_LOCATOR` > empty) or `a.Lite()` alone for sqlite-only raw stores -
`--lite` and `LITE_PATH` mean sqlite, nothing else. See `database.md` for the
selection convention. The locator flag exists for development; credentials
travel via environment, not argv.

## Run Function

`Run(o, i)` constructs components and wires lifecycle. It receives
the instrument from `Main()` as `face.Instrument` - not on the
option struct - and pulls the reporter and recorder where it wires
them (see `entrypoint.md`). See `pillars.md` for the full wiring
rationale.

```go
func Run(o *option.Log, i face.Instrument) {
    r := i.Reporter()
    l := logger.New(context.Background())
    v := service.New(o, l, r)
    defer v.Close()
    lifecycle.New(l).
        WithWorker(worker.New(v, l, r, 1*time.Minute)).
        WithServer(
            server.New(
                constant.Identity,
                o.Address,
                func(m *http.ServeMux) {
                    Mount(v, r, i.Recorder(), o.Version, guard.New(m, o.ServiceTokens))
                },
            ).WithMiddleware(web.RecoveryMiddleware(r)),
        ).
        RunUntilSignal()
}
```

### Key conventions

- The reporter goes to the workers (`recovery.New`) and the recovery
  middleware; add `.WithProtected()` to a plain REST server
- The listen address is `o.Address` from `a.Address()`; a surface without
  flags (examples, callback servers) uses `web.AddressPort`, the same
  `BIND_ADDRESS` > loopback chain - never format one by hand
- The server callback builds one `guard.Mux` and hands it to `Mount` -
  routes register there, never on the mux. `guard.New` mounts
  `route:/api/reachable` itself, the peer probe route
  (`error-handling/infrastructure.md`, The probe route)
- The service's `Close` is deferred before the lifecycle starts, so it
  closes after every worker and server has
- A service with peers is waited on next -
  `errors.PanicOnError(v.Wait(context.Background()))` - and its
  `Workers()` registered before the daemon's own, so the guards run
  before anything calls through them

## Mount Function

`mount.go` sits beside `run.go` and owns everything the daemon
serves: the token-guarded REST tree, the MCP mount, and the web
mount when present. The service comes first - then reporter, recorder,
version, and the `*guard.Mux` last:

```go
func Mount(
    v log.Log,
    r face.Reporter,
    t face.Recorder,
    version string,
    g *guard.Mux,
) {
    g.TokenMount(...)                             // REST - generated-api.md
    model_context.New(v, r, t, version).Mount(g)  // MCP - model-context.md
    web.New(v).Mount(g)                           // web - HTML Web Package below
}
```

`Mount` is the production surface: the guard battery
(`<path>/unit/guard_test.go`, see `testing.md`) starts it on
a dynamic port and asserts the full auth contract, and goaudit's
guard rules enforce that the battery and the mount shape exist. A
route registered in the run.go callback instead of `Mount` escapes
both the guard and the battery.

## Server Functions

When routes are defined by an OpenAPI spec, use the Server struct
pattern - see `generated-api.md`.

For manually registered routes without a spec, route functions live in
the `server/` package. Each is a function returning `http.HandlerFunc`:

```go
func Alerts(v *service.Service) http.HandlerFunc {
    return func(
        w http.ResponseWriter,
        r *http.Request,
    ) {
        // ...
        web.EncodeNotation(w, result)
    }
}
```

- `web.EncodeNotation(w, result)` - sets Content-Type JSON header and encodes
  result as JSON (the common case)
- `web.Encode(w, result)` - encodes result as JSON without setting headers (use
  when headers are set separately, e.g. with a non-200 status code)
- `web.ObjectHeader(w)` - sets Content-Type JSON header only (use with
  `web.Encode` when you need a custom status code between header and body)
- Use `argument.*` constants for query parameter names
- Response structs in `response.go`, constants in the tool's `constant/`

## HTML Web Package

When a service tool serves an HTML UI (using gomponents), routes are grouped
under a `web/` package rather than `server/`. Use this pattern instead of
`server/` when handlers need shared state and render HTML rather than JSON.

```
pkg/tool/go<tool>d/web/
├── server.go           # type Server struct (dependencies + view)
├── new.go              # New(deps) *Server - constructs view
├── mount.go            # (s *Server) Mount(g *guard.Mux)
├── recovery.go         # (s *Server) Recovery(r) - delegates to view.Recovery
├── constant.go         # inlineCSS (component CSS specific to this service)
├── <route>.go          # handler method: dashboard(), alerts()
└── <component>.go      # HTML builder: alerts_table.go, add_form.go
```

### Recovery

A web service wires the layout-aware recovery middleware instead of
the generic one: construct the web server before `lifecycle.New`,
mount with `w.Mount`, and pass `w.Recovery(r)` to `WithMiddleware`.
Panics then render into the page layout (notification item with the
Sentry event ID; fragment form for HTMX requests) instead of a
plain-text 500. The `recovery.go` delegate is two lines:

```go
func (s *Server) Recovery(r face.Reporter) func(http.Handler) http.Handler {
    return s.view.Recovery(r)
}
```

### Form Errors

Validation errors are not panics. Full-page forms use the round
trip: the submit handler calls `form.Redirect(w, r, path, message,
values)` to bounce back with the message and the submitted values
in query parameters; the GET handler reads `form.ErrorText(r)` and
renders `layout.Alert(message)` above the repopulated form.
HTMX-swapped forms re-render the fragment with the alert inside the
swap target and values preserved. Sentinel-to-friendly-message
mapping stays in the service, next to the sentinel.

### View and Layout

The `Server` holds a `*view.View` constructed in `New`. The view wraps
a `*layout.Page` template configured with the service's identity, theme,
style, and navigation items. Handlers call `s.view.RenderPage` for full pages
and `s.view.RenderFragment` for HTMX partials.

```go
type Server struct {
    service  *service.Service
    view     *view.View
    registry *registry.Registry
}

func New(v *service.Service) *Server {
    r := registry.New()
    r.Register(
        palette.NewCommand(constant.DashboardTitle, constant.DashboardPath, webConstant.PaletteNavigate),
        palette.NewCommand(constant.EntriesTitle, constant.EntriesPath, webConstant.PaletteNavigate),
    )

    return &Server{
        service:  v,
        registry: r,
        view: view.New(
            layout.New(constant.Identity).
                WithTheme(theme.Archive).
                WithStyle(inlineCSS).
                WithCommandPalette("/palette").
                WithItems(
                    navigation_item.New(constant.DashboardPath, constant.DashboardTitle),
                    navigation_item.New(constant.EntriesPath, constant.EntriesTitle),
                ),
        ),
    }
}
```

#### Summary strip and live updates

**Summary strip** - `layout.WithSummary(items ...string)` renders
a dot-separated muted line inside main. The view exposes
`RenderLivePageWithSummary` for pages that need dynamic summary
values (alongside `RenderLivePage` for live pages without one).
The strip is absent when the summary slice is nil or empty.

**Live updates** - `layout.WithLiveEndpoint(path)` adds the htmx
SSE extension script and `hx-ext="sse"` + `sse-connect` attributes
to `main`. The summary strip gets `sse-swap="summary_strip"`
automatically. Services push updates via `layout.SummaryStripContent`
with event name `layout.SummaryStrip`.

#### Handlers use the view

Handlers call the view directly, with no wrapper functions of their own:

```go
func (s *Server) dashboard(w http.ResponseWriter, r *http.Request) {
    if s.view.IsExtendedRequest(r) {
        s.view.RenderFragment(w, s.recentTable())
        return
    }
    s.view.RenderPage(w, constant.DashboardTitle, constant.DashboardPath,
        html.H1(gomponents.Text(constant.DashboardTitle)),
        s.recentTable(),
    )
}
```

### Page titles and paths

Each service declares page titles and paths as constants in its `constant/`
package. Both navigation items and handlers reference the same constants:

```go
const (
    DashboardTitle = "Dashboard"
    DashboardPath  = "/"
    EntriesTitle   = "Entries"
    EntriesPath    = "/entries"
)
```

### Themes

Theme constants (`Theme*` in `pkg/web/constant/`, one `theme_<name>.go` per
palette) are pico.css custom property overrides. Each service picks one in its
`New`; the files are the current set. Additional palettes can be defined in
downstream repos.

Key conventions:
- `Mount()` takes a `guard.Mux` and registers all routes through its verbs using
  method values: `g.Open(route.Get("/alerts"), s.alerts)` for board routes on
  non-SSO services; SSO services call `s.authorization.Mount(g)` first, which
  installs the client's session check and opens the auth trio, and register
  their protected routes with `g.Session` (favicon stays `Open`). A service
  whose session check differs, the way gogated's admin check does, calls
  `g.WithSession` itself and opens the client's three methods directly. A
  session route also answers to the service bearer token, so the tokens that
  open REST and MCP open the pages too, with the same full access - in the
  `Authorization` header only, never a cookie or a query string. Bare
  `m.HandleFunc` in a web package is flagged by the `route_guard` analyzer.
- Handler methods named after the route, no `handle` prefix: `alerts()`,
  `dashboard()`, `addSubmit()`
- Standalone HTML builders named after the component they produce:
  `alertsTable()`, `addForm()`, `detailRow()`
- Handler and builder names never collide because the handler is named after the
  route (`add`), not the component (`addForm`)
- Methods that access server state (the service) stay as methods; pure
  renderers are standalone functions

## Service

Every daemon has one `service.Service`, even around a single client or
none. It is the only owner `Run` holds: everything else in the daemon
either is built by it or borrows from it.

The service takes configuration - the option, a logger, a reporter, and
a metric registry when its workers export gauges - and builds the rest:
the database (`relational.Open`, `migrate.AutoMigrate`, then the store
over it, `database.md`), its clients, pools, peer guards and workers.
Its `Close` closes everything it built and nothing it was given. A
service that owns nothing - a reader over the local system, paths and
a logger - has no `Close`, and `Run` defers nothing.

A service is never handed a finished client, in a parameter or an option
function: whoever closes a handed client closes something it did not
build. A client enters as a factory on the option - a function
returning the client's face, set in `option.New()` to the real
constructor, or set in `Main` when the environment decides whether
there is one. The service calls it once and closes what it returns when
that has a `Close`. A part the environment may leave out - a factory
returning nil, a configuration that is off - the service builds only
when present; its operations then answer `not_configured`, and a
surface asks the service rather than testing a field for nil. A pooled
client's factory takes the instance -
`func(*instance.Instance) (face.<Name>Source, error)` - the build function
`pool.New` takes. An option holds what `Main` decides once: addresses,
tokens, a parsed inventory, and factories, which own nothing until the
service calls them. Nothing live sits on an option.

Workers the service builds reach the lifecycle through `Workers()`,
which `Run` registers in order. The lifecycle closes them, and the
service closes them again, so their `Close` is safe twice
(`lifecycle.md`).

A service with required peers has a `Wait` that waits on each required
peer's guard in turn. `Run` calls it right after building the service and
panics on its error, before the lifecycle starts, so the daemon serves only
once every required peer answers; a peer only a worker reads is guarded and
not waited on (`error-handling/infrastructure.md`, Required and guarded
peers).

The service validates. A surface parses its request into plain values,
calls one operation, and answers the error's class
(`error-handling/README.md`); a check that lived in a route moves into
the service, where every surface shares it. A surface holds the service
and nothing beside it - no client, store or guard of its own to reach
around it. Collection belongs in the service too: a worker gets one
operation per tick that returns its outcomes, and keeps only what is its
own - the ticker, its gauges, its stop channel and the recovery around
the tick.

A face describes one client in full and its mock implements all of it: the
upstream client's face is what the factory returns, the daemon client's is
`face.Client` (`generated-api.md`). Code shared between daemons and tools
takes a role interface - only the methods it calls, beside the full face
(`FloatSource`), filled by the client, its mock or a service's guarded
method through a function type beside it (`FloatSourceFunction`). No
other face is partial or mirrors another; the service has none, and the
surfaces, workers and tests take `*service.Service` (`testing.md`).

## Workers

Workers implement `face.Worker` (`Start()` + `Close()`) and close only
what they started or built - a service or client they were given
belongs to whoever built it. See `lifecycle.md` for details.

## Daemon and CLI Split

A daemon `go<tool>d` pairs with a CLI `go<tool>`, which talks to it
through the daemon client in `<toold>/client/` - the one client other
daemons use too (`generated-api.md`, The Daemon Client).

## External API Client

When a daemon proxies a third-party service (Habitica, Mattermost,
Jellyfin, Sentry, etc.), the client library for that service
lives at `pkg/<domain>/`. This is the external API client - it talks
directly to the upstream service, not to our daemon.

```
pkg/<domain>/
├── client.go                       # Client struct (raw HTTP or SDK wrapper)
├── new.go                          # New(host, token string) *Client
├── new_environment.go              # NewEnvironment() - reads env vars
├── <operation>.go                  # Non-Must: returns (value, error)
├── must_<operation>.go             # Must: panics on error
└── constant/
    └── constant.go                 # HostEnvironment, TokenEnvironment
```

The daemon's service builds this client through its factory (Service); no
surface imports it. **`pkg/<domain>/` is client-only.** All server-side concerns
(store, domain types, `model_context/`) belong in `pkg/tool/go<tool>d/`, not in
the shared library.

### Must and Non-Must Pattern

External API client methods come in pairs:

- `<operation>.go` - returns `(value, error)`. Used by daemon handlers
  that need to propagate errors gracefully (MCP and REST answer them by
  class through `answer`).
- `must_<operation>.go` - panics on error. Used by CLI tools and
  background workers running under lifecycle recovery.

```go
// links.go
func (c *Client) Links() ([]*link.Link, error) {
    // ... returns error on failure
}

// must_links.go
func (c *Client) MustLinks() []*link.Link {
    result, e := c.Links()
    errors.PanicOnError(e)
    return result
}
```

A Must twin exists only where a caller panics. Its callers are counted
by receiver, not by name - a `MustLinks` on another type is not a
caller - and a twin with none is deleted before the client's full face
is written, so the face carries no operation nobody calls. Every
consumer of an upstream client moves onto its full face and mock in the
change that writes them. A client that owns nothing - a connection per
call - has no `Close`, and neither does its face.

### Find Pattern

Lookups where absence is a valid outcome return the Find shape -
`Find<Noun>(...) (*T, bool, error)`. The flag carries absence,
the error carries failure; never a `nil, nil` soft-miss that
callers must remember to check. When absence is exceptional for
every caller, return a typed `not_found` instead of the flag.
`<Noun>Exists` variants wrap a lookup and swallow only
`not_found`. List endpoints are neither: an empty slice is a
complete answer, not a miss. Examples: `FindBranchRequest`,
`FindPostBefore`, `FindSession`, `FindLatestPulse`,
`FindDocument`, `FileExists`.

Helpers that compose other calls (`ProjectsWithFile`, `Jobs`,
`RunnersByTag`) follow the same pair: the plain form composes the
plain variants and returns the first error, and a Must twin panics
on it. A method body never calls a Must variant - the panic belongs
to the caller's choice of variant, not to the helper, and a plain
signature that panics underneath breaks the MCP and REST callers
the pair exists for.

`pkg/<domain>/` talks to the third-party service; `<toold>/client/` talks
to our own daemon (`generated-api.md`, The Daemon Client).

## MCP Integration

When a daemon also exposes MCP tools, add a `model_context/`
subpackage and mount it through the daemon's `Mount()` on the same
`guard.Mux` as the REST API. See `model-context.md`.

`<path>/model_context/` is the standard package name - not `<path>/tool/` or
`<path>/toolset/`.
