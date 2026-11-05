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
├── run.go                          # Run(o, i): builds the daemon, wires lifecycle, signal block
├── option/
│   ├── <name>.go                   # Option struct (named after tool/domain, not "Option")
│   └── new.go                      # Factory: New() *<Name>
├── daemon/                         # The one owner - see The daemon
│   ├── daemon.go                   # Daemon struct
│   ├── new.go                      # New(o, l, r) - builds everything from configuration
│   ├── mount.go                    # Mount(...) - hands each surface the services it calls
│   ├── workers.go                  # Workers() - guards first, then clocks
│   └── close.go                    # Close() - closes everything it built
├── service/                        # The use cases - see Services
│   ├── service.go                  # one set: Service struct, its code here
│   ├── new.go
│   └── <operation>.go              # One file per operation the surfaces call
├── store/                          # Persistence (if needed) - see Stores
│   ├── store.go                    # one store: Store struct
│   ├── new.go                      # Constructor: takes the mapper it is given
│   └── <operation>.go              # One file per operation (save.go, resolve.go, by_name.go)
├── face/                           # One full interface per client - see The daemon
├── migrate/                        # AutoMigrate(d) - schema history, see database.md
├── constant/
│   └── constant.go                 # Tool constants (bucket names, identity)
├── unit/
│   ├── option_test.go              # assert.NotNil(t, option.New()) - see test-placement.md
│   ├── guard_test.go               # Guard battery over Mount - see testing.md
│   └── <noun>_tester/new.go        # Builds one service from its parts - see testing.md
└── server/                         # REST implementation (or manual routes)
    ├── <operation>.go              # Route function returning http.HandlerFunc
    └── response.go                 # Response struct (if needed)
```

A daemon with one set of use cases keeps them in `service/` as
`service.Service`. One with several - a library daemon's books, loans and
members - splits by noun: `<path>/service/book/`, `<path>/service/loan/`,
`<path>/service/member/`, each a `book.Service`, `loan.Service`,
`member.Service`,
and `service/` holds only those directories. A store splits the same way
only along consistency boundaries (`store/<aggregate>/`, see Stores). Split
by what each set touches, not by kind of operation: a package per noun
stays small; a package per concern moves the bulk down a level.

goaudit holds a daemon with `<path>/daemon/` to this layout on every run
(`missing_service`, `mixed_service`, `service_imports_mapper`,
`run_builds_clock`, `run_builds_client`, `daemon_not_closed`,
`missing_service_mount`); `surface_holds_daemon` runs everywhere.

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

`Run(o, i)` builds the daemon and wires lifecycle. It receives the
instrument from `Main()` as `face.Instrument` - not on the option struct -
and pulls the reporter and recorder where it wires them (see
`entrypoint.md`). See `pillars.md` for the full wiring rationale.

```go
func Run(o *option.Library, i face.Instrument) {
    r := i.Reporter()
    l := logger.New(context.Background())
    d := daemon.New(o, l, r)
    defer d.Close()
    errors.PanicOnError(d.Wait(context.Background()))
    c := lifecycle.New(l)

    for _, w := range d.Workers() {
        c.WithWorker(w)
    }

    c.WithServer(
        server.New(
            constant.Identity,
            o.Address,
            func(m *http.ServeMux) {
                d.Mount(r, i.Recorder(), o.Version, guard.New(m, o.ServiceTokens))
            },
        ).WithMiddleware(web.RecoveryMiddleware(r)),
    ).RunUntilSignal()
}
```

### Key conventions

- `Run` holds the daemon and nothing else: it builds no client, no service
  and no clock
- The reporter goes to the recovery middleware; add `.WithProtected()` to a
  plain REST server
- The listen address is `o.Address` from `a.Address()`; a surface without
  flags (examples, callback servers) uses `web.AddressPort`, the same
  `BIND_ADDRESS` > loopback chain - never format one by hand
- The server callback builds one `guard.Mux` and hands it to `d.Mount` -
  routes register there, never on the mux. `guard.New` mounts
  `route:/api/reachable` itself, the peer probe route
  (`error-handling/infrastructure.md`, The probe route)
- The daemon's `Close` is deferred before the lifecycle starts, so it
  closes after every worker and server has
- A daemon with required peers is waited on next, before the lifecycle
  starts; its `Workers()` are registered in the order it returns them,
  guards first

## Mount Function

`Mount` is the daemon's method, in `daemon/mount.go`, and owns everything
the daemon serves: the token-guarded REST tree, the MCP mount, and the web
mount when present. It hands each surface the services that surface calls,
and takes only what its surfaces need beside them - here reporter,
recorder and version - with the `*guard.Mux` last. A web the daemon built
in `New` for its recovery leaves `Mount` only the mux. A daemon serving a
second listener has a mount method per listener (`MountLocal`,
`MountBridge`) beside `Mount`:

```go
func (d *Daemon) Mount(
    r face.Reporter,
    t face.Recorder,
    version string,
    g *guard.Mux,
) {
    g.TokenMount(...)                                         // REST - generated-api.md
    model_context.New(d.book, d.loan, r, t, version).Mount(g) // MCP - model-context.md
    web.New(d.book, d.loan, d.member).Mount(g)                // web - HTML Web Package below
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
the generic one. The daemon builds its web surface in `New`, `Mount`
mounts it beside the other surfaces, and the daemon hands its recovery
out - `Run` passes `d.Recovery(r)` to `WithMiddleware`, which is
evaluated before the mount callback runs, and still holds only the
daemon. Panics then render into the page layout (notification item with
the Sentry event ID; fragment form for HTMX requests) instead of a
plain-text 500. Both delegates are two lines:

```go
// web/recovery.go
func (s *Server) Recovery(r face.Reporter) func(http.Handler) http.Handler {
    return s.view.Recovery(r)
}

// daemon/recovery.go
func (d *Daemon) Recovery(r face.Reporter) func(http.Handler) http.Handler {
    return d.web.Recovery(r)
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
a `*layout.Page` template configured with the daemon's identity, theme,
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
- Methods that access server state (its services) stay as methods; pure
  renderers are standalone functions

## The daemon

Every daemon has one `daemon.Daemon`, small ones included. It is the only
owner `Run` holds: it builds everything else and closes what it built. It
does no work of its own.

The daemon takes configuration - the option, a logger, a reporter, and a metric
registry when its workers export gauges - and builds the rest: the database
(`relational.Open`, `migrate.AutoMigrate`, then each store over it,
`database.md`), its clients, pools and peer guards, then each service from the
parts that service uses. It holds each thing it owns in one field - the same
guard in two fields reads as an owner never closed. Its `Close` closes
everything it built and nothing it was given. A daemon that owns nothing - a
reader over the local system, paths and a logger - has no `Close`, and `Run`
defers nothing.

A daemon is never handed a finished client, in a parameter or an option
function: whoever closes a handed client closes something it did not
build. A client enters as a factory on the option - a function returning
the client's face, set in `option.New()` to the real constructor, or set
in `Main` when the environment decides whether there is one. The daemon
calls it once and closes what it returns when that has a `Close`. A client
the environment may leave out - a factory returning nil, a configuration
that is off - the daemon builds only when present; the services that use
it then answer `not_configured`, and a surface asks the service rather
than testing a field for nil. A pooled client's factory takes the
instance - `func(*instance.Instance) (face.<Name>Source, error)` - the
build function `pool.New` takes. An option holds what `Main` decides once:
addresses, tokens, a parsed inventory, and factories, which own nothing
until the daemon calls them. Nothing live sits on an option.

A daemon with required peers has a `Wait` that waits on each required
peer's guard in turn. `Run` calls it right after building the daemon and
panics on its error, before the lifecycle starts, so the daemon serves
only once every required peer answers; a peer only a clock reads is
guarded and not waited on (`error-handling/infrastructure.md`, Required
and guarded peers).

A face describes one client in full and its mock implements all of it: the
upstream client's face is what the factory returns, the daemon client's is
`face.Client` (`generated-api.md`). Code shared between daemons and tools
takes a role interface - only the methods it calls, beside the full face
(`FloatSource`), filled by the client, its mock or a service's guarded
method through a function type in a sub-package of the face
(`float_source_function.FloatSourceFunction`). No other
face is partial or mirrors another; neither the daemon nor a service has
one.

## Services

The use cases live in services. A service holds only what the daemon built
it with - the stores, clients, guards and clock it uses, and the other
services it calls - and never the daemon: the daemon holds the service, so
the service holding the daemon is an import cycle. A service calling
another is built with it as a constructor parameter, so the daemon builds
them in dependency order and the calls between services run one way:

```go
// daemon/new.go
m := member.New(s.member)
b := book.New(s.catalog, o.Random)
l := loan.New(s.catalog, m, b, d.clock)
```

State only one service uses - a waiter map, a cache it reads, who is
watching - is that service's own field. A view gathering several services -
a member's page with their loans and holds - is a service of its own, built
with the services it reads, named for what it shows.

The service validates. A surface parses its request into plain values, calls one
operation, and answers the error's class (`error-handling/README.md`); a check
that lived in a route moves into the service, where every surface shares it. An
update gets the record, applies the change, checks and saves - a change applied
inside the store hides the result from the check - and sets only the fields it
was sent. A write that may still be refused is staged and moved into place once
every check passed. A value joined into a path, a URL, a selector or a script is
checked as the narrowest shape it can be - a bare name, a bare number - since an
encoded `%2e%2e` passes `http.ServeMux`'s cleaning. A surface holds the services
it calls and nothing beside them - no client, store or guard of its own to reach
around them, and never the daemon. The daemon keeps no method that only forwards
to a service: each surface is handed the service itself. The daemon's work
belongs in a service, never in a worker or the daemon: a tick or an event calls
one operation that returns its outcomes.

A service test builds that service from its parts and the services it was
built with; a surface test goes served through `Mount` (`testing.md`).
Nothing builds a surface over hand-wired services.

## Stores

A store borrows the mapper the daemon opened and has no `Close`
(`database.md`, Store Shape). It saves and fires nothing: the service
notifies after a write that succeeded. One store serves a daemon whose tables
commit independently. Where some tables must commit together and others
never do, the store splits along those consistency boundaries - one store
per aggregate, `store/<aggregate>/` - and a transaction never spans two.
Tables one service owns outright get its own store; tables several services
write in one transaction stay one store, shared by every service writing
them.

A sequence that must commit together runs through the store's
`Atomically`. The store opens, commits and rolls back; the service names
the steps and passes each the transaction's store:

```go
// store/catalog/atomically.go
func (s *Store) Atomically(apply func(*Store) error) error {
    return s.mapper.Transaction(
        func(t *gorm.DB) error {
            return apply(New(t))
        },
    )
}

// service/loan/return.go
func (s *Service) Return(identifier int) error {
    return s.catalog.Atomically(
        func(t *catalog.Store) error {
            if e := s.close(t, identifier, s.clock()); e != nil {
                return e
            }

            return s.restock(t, identifier)
        },
    )
}
```

A step inside reads the store it was passed, never the service's own
field. No service imports the mapper's package or opens a transaction
itself, and nothing calls a peer between `Atomically`'s open and commit:
what a sequence reads from a peer it reads before, and what it tells a
peer it tells after the commit.

## Workers

Workers implement `face.Worker` (`Start()` + `Close()`) and close only
what they started or built - a service or client they were given
belongs to whoever built it. See `lifecycle.md` for details.

A worker that holds state or a connection - a peer guard, a runner a
service launches work into, a socket - is built by the daemon and reaches
back only through function values it was given at construction, never the
daemon: the daemon holds the worker, so the worker holding the daemon is an
import cycle.

A worker that holds nothing but time is a clock. The daemon builds it over
one service operation per tick - goclauded's
`worker.New(v.Reconcile, l, r, constant.ReconcileInterval)` on
`pkg/worker` - and it keeps only what is its own: the ticker, its gauges,
its stop channel and the recovery around the tick. Its `Close` cancels the
context the tick was handed and waits for a tick in flight, so nothing the
daemon closes next is still in use; what a cancelled tick returns is not
reported. No service holds a clock.

The daemon hands its workers out through `Workers()`, guards first, so
they run before anything calls through them, then its clocks:

```go
// daemon/workers.go
func (d *Daemon) Workers() []face.Worker {
    return []face.Worker{
        d.catalogGuard,
        worker.New(d.loan.Remind, d.logger, d.reporter, constant.RemindInterval),
    }
}
```

The lifecycle closes them, and the daemon closes what it built again, so
their `Close` is safe twice (`lifecycle.md`).

## Daemon and CLI Split

A daemon `go<tool>d` pairs with a CLI `go<tool>`, which talks to it
through the daemon client in `<toold>/client/` - the one client other
daemons use too (`generated-api.md`, The Daemon Client). A rule deciding
what the daemon's data means - which threads count as unresolved - is a
daemon operation, never the CLI's, so every caller gets the same answer.

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

The daemon builds this client through its factory (The daemon); no
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
