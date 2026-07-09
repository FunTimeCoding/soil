# Service Tool Spec

Pattern for long-running service tools that use lifecycle for HTTP servers and background workers.

## Package Structure

```
cmd/go<tool>/
└── main.go                         # Linker vars, delegates to Main()

pkg/tool/go<tool>/
├── main.go                         # Main(): register flags, parse, build option, call Run()
├── run.go                          # Run(o): wiring, lifecycle, signal block
├── option/
│   ├── <name>.go                   # Option struct (named after tool/domain, not "Option")
│   ├── new.go                      # Factory: New() *<Name>
│   └── <name>_test.go              # assert.NotNil(t, New())
├── store/                          # Persistence (if needed)
│   ├── store.go                    # Store struct
│   ├── constant.go                 # Bucket name
│   ├── record.go                   # Record type
│   ├── new.go                      # Constructor: opens db, ensures bucket
│   ├── close.go                    # Close()
│   └── <operation>.go              # One file per operation (save.go, resolve.go, by_name.go)
├── worker/                         # Background worker (if needed)
│   ├── worker.go                   # Worker struct
│   ├── new.go                      # Constructor
│   ├── start.go                    # Start() - goroutine with ticker
│   ├── stop.go                     # Stop()
│   └── poll.go                     # Poll() - single cycle logic
└── server/                         # REST implementation (or manual routes)
    ├── <operation>.go              # Route function returning http.HandlerFunc
    ├── response.go                 # Response struct (if needed)
    └── constant.go                 # Server-specific constants
```

## Entry Point

See `entrypoint.md` for linker variables, `Main()`, and sentry setup.

Every serving daemon registers the common web parameters through the
argument instance - one call, same everywhere:

```go
a := argument.NewInstance(constant.Identity)
a.Web()
a.Parse(version, gitHash, buildDate)
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

Services with a database also call `a.Database()` (registers `--lite`
with chain `LITE_PATH` > `Identity.LitePath()`, and `--postgres` with
chain `POSTGRES_LOCATOR` > empty) or `a.Lite()` alone for
sqlite-only raw stores; key-value (bbolt) stores keep a plain
`--path` - `--lite` and `LITE_PATH` mean sqlite, nothing else. See
`database.md` for the selection convention. The locator flag exists for development; credentials
travel via environment, not argv.

## Run Function

`Run(o, r)` constructs components and wires lifecycle. It receives the
reporter from `Main()` as `face.Reporter` - not on the option struct.
See `pillars.md` for the full wiring rationale.

```go
func Run(o *option.Log, r face.Reporter) {
    l := logger.New(context.Background())
    s := store.New(o.LitePath)
    defer s.Close()
    lifecycle.New(
        l,
        lifecycle.WithWorker(worker.New(client, s, l, r, 1*time.Minute)),
        lifecycle.WithServer(
            server.New(
                o.Address,
                func(m *http.ServeMux) {
                    m.HandleFunc("/api/alerts", server.Alerts(s))
                },
            ).WithMiddleware(web.RecoveryMiddleware(r)),
        ),
    ).RunUntilSignal()
}
```

Key conventions:
- Logger constructed first, threaded to workers and lifecycle
- Reporter threaded to workers (for `recovery.New`) and recovery middleware
- Use `WithServer(server.New(...).WithMiddleware(web.RecoveryMiddleware(r)))` - add `.WithProtected()` for plain REST servers
- Server address is `o.Address`, produced by `a.Address()` in
  `Main()` - see Entry Point. Never format a listen address by hand;
  surfaces without flags (examples, callback servers) use
  `web.AddressPort`, which resolves the same `BIND_ADDRESS` >
  loopback chain
- Routes registered in the `func(*http.ServeMux)` callback
- `RunUntilSignal()` handles run, signal block, and reverse-order stop
- Store closed via `defer` before lifecycle starts
- Avoid declaring intermediate variables for lifecycle or generative server when only used once

## Server Functions

When routes are defined by an OpenAPI spec, use the Server struct
pattern - see `generated-api.md`.

For manually registered routes without a spec, route functions live in
the `server/` package. Each is a function returning `http.HandlerFunc`:

```go
func Alerts(s *store.Store) http.HandlerFunc {
    return func(
        w http.ResponseWriter,
        r *http.Request,
    ) {
        // ...
        web.EncodeNotation(w, result)
    }
}
```

- `web.EncodeNotation(w, result)` - sets Content-Type JSON header and encodes result as JSON (the common case)
- `web.Encode(w, result)` - encodes result as JSON without setting headers (use when headers are set separately, e.g. with a non-200 status code)
- `web.ObjectHeader(w)` - sets Content-Type JSON header only (use with `web.Encode` when you need a custom status code between header and body)
- Use `argument.*` constants for query parameter names
- Response structs in `response.go`, constants in `constant.go`

## HTML Web Package

When a service tool serves an HTML UI (using gomponents), routes are grouped
under a `web/` package rather than `server/`. Use this pattern instead of
`server/` when handlers need shared state and render HTML rather than JSON.

```
pkg/tool/go<tool>d/web/
├── server.go           # type Server struct (dependencies + view)
├── new.go              # New(deps) *Server - constructs view
├── mount.go            # (s *Server) Mount(m *http.ServeMux)
├── constant.go         # inlineCSS (component CSS specific to this service)
├── <route>.go          # handler method: dashboard(), alerts()
└── <component>.go      # HTML builder: alerts_table.go, add_form.go
```

### View and Layout

The `Server` holds a `*view.View` constructed in `New`. The view wraps
a `*layout.Page` template configured with the service's identity, theme,
style, and navigation items. Handlers call `s.view.RenderPage` for full pages
and `s.view.RenderFragment` for HTMX partials.

```go
type Server struct {
    store    *store.Store
    view     *view.View
    registry *palette.Registry
}

func New(s *store.Store) *Server {
    registry := palette.NewRegistry()
    registry.Register(
        palette.Command{Label: constant.DashboardTitle, Path: constant.DashboardPath, Category: "navigate"},
        palette.Command{Label: constant.EntriesTitle, Path: constant.EntriesPath, Category: "navigate"},
    )

    return &Server{
        store:    s,
        registry: registry,
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

**Summary strip** - `layout.WithSummary(items ...string)` renders
a dot-separated muted line inside main. The view exposes
`RenderLivePageWithSummary` for pages that need dynamic summary
values (alongside `RenderLivePage` for live pages without one).
The strip is absent when the summary slice is nil or empty.

**Live updates** - `layout.WithLiveEndpoint(path)` adds the htmx
SSE extension script and `hx-ext="sse"` + `sse-connect` attributes
to `main`. The summary strip gets `sse-swap="summary_strip"`
automatically. Services push updates via `layout.SummaryStripContent`
with event name `layout.SummaryStrip`. Replaces manual
`WithScript(web.ServerSide)` + wrapper div pattern.

Handlers use the view directly - no `pageLayout`, `renderPage`,
`renderFragment`, `isHTMX`, or `navigationLink` wrapper functions:

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

### Packages involved

| Package | Location | Role |
|---------|----------|------|
| `layout` | `pkg/web/layout/` | Fluent HTML page builder. `New(identity)` + `With*` methods + `Render()`. Owns the full page skeleton (head, nav, main, footer). |
| `layout/navigation_item` | `pkg/web/layout/navigation_item/` | `New(path, label)` / `NewExternal(path, label)` for nav links. Renders with active-state highlighting. |
| `palette` | `pkg/web/palette/` | Command palette framework. `NewRegistry()`, `Register()`, `NewServe(registry)` handler. fzf-style fuzzy matching, highlighted results. Layout adds `⌘K` nav hint and `<dialog>` automatically via `WithCommandPalette`. |
| `view` | `pkg/web/view/` | HTTP layer wrapping layout. `RenderPage`, `RenderFragment`, `IsExtendedRequest`. Clones the layout template per request. |
| `theme/constant` | `pkg/web/theme/constant/` | CSS palette constants (pico.css custom property overrides). |

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

Theme constants in `pkg/web/theme/constant/` are pico.css custom property
overrides. Each service picks one in its `New`. One file per palette in
that package (Amethyst, Archive, Cortex, Ember, Hearth, Sentinel, Slate,
Sprout, Tangerine, Tyria - check the directory for the current set).
Additional palettes can be defined in downstream repos.

Key conventions:
- `Server` struct holds dependencies and a `*view.View`
- `Mount()` registers all routes using method values: `m.HandleFunc("GET /alerts", s.alerts)`
- Handler methods named after the route, no `handle` prefix: `alerts()`, `dashboard()`, `addSubmit()`
- Standalone HTML builders named after the component they produce: `alertsTable()`, `addForm()`, `detailRow()`
- Handler and builder names never collide because the handler is named after the route (`add`), not the component (`addForm`)
- Methods that access server state (store, worker) stay as methods; pure renderers are standalone functions

## Workers

Workers implement `face.Worker` (`Start()` + `Stop()`). See `lifecycle.md` for details.

## Daemon / CLI Split

Long-running services come in pairs:

| Binary | Package | Role |
|--------|---------|------|
| `go<tool>d` | `pkg/tool/go<tool>d/` | Daemon - lifecycle, store, HTTP server |
| `go<tool>` | `pkg/tool/go<tool>/` | CLI - calls the daemon's REST API, prints output |

After the standard entrypoint setup (see `entrypoint.md`), the CLI
constructs the daemon's generated REST client once and hands it to
its subcommands - see `generated-api.md` for the access pattern and
when a hand-written `client/` wrapper is warranted instead.

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

The daemon imports this client library. MCP-only services (without
REST) import it directly into their `model_context/` handlers. Daemons
with REST import it into both `model_context/` and `server/`.

**`pkg/<domain>/` is client-only.** All server-side concerns (store,
domain types, `model_context/`) belong in `pkg/tool/go<tool>d/`, not
in the shared library.

### Must/non-Must Pattern

External API client methods come in pairs:

- `<operation>.go` - returns `(value, error)`. Used by daemon handlers
  that need to propagate errors gracefully (MCP `captureFail`, REST
  `http.Error`).
- `must_<operation>.go` - panics on error. Used by CLI tools and
  background workers running under lifecycle recovery.

```go
// links.go
func (c *Client) Links(collection int32) ([]*link.Link, error) {
    // ... returns error on failure
}

// must_links.go
func (c *Client) MustLinks(collection int32) []*link.Link {
    result, e := c.Links(collection)
    errors.PanicOnError(e)
    return result
}
```

Helper methods that iterate over results (e.g. `CollectionByName`,
`TagByName`) call the Must variants internally since their callers
expect panics.

## Internal REST Client

When a daemon exposes a REST API (via OpenAPI codegen), consumers
that want domain methods instead of generated request objects -
other daemons, test mocks, formatting layers - use a hand-written
wrapper inside the daemon's tree at `<toold>/client/`, not at
`pkg/<domain>/`. CLIs call the generated client directly. See
`generated-api.md` for the full pattern.

The distinction:
- `pkg/<domain>/` - external API client, talks to a third-party service
- `<toold>/client/` - internal REST client, talks to our own daemon

## MCP Integration

When a daemon also exposes MCP tools, add a `model_context/` subpackage
and mount it on the same mux as the REST API. See `model-context.md`.

`model_context/` is the standard package name - not `tool/` or
`toolset/`.
