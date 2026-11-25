---
base: pkg/tool/goatlassiand
---

# Service Tool Spec

Pattern for long-running service tools that use lifecycle for HTTP servers and
background workers.

- `daemon.md` - the one owner, services, stores, workers, the CLI split
- `client.md` - external API clients, Must pairs, the Find shape
- `web.md` - the HTML web package, guarding, recovery, form errors, the view

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
├── daemon/                         # The one owner - see daemon.md
│   ├── daemon.go                   # Daemon struct
│   ├── new.go                      # New(o, l, r) - builds everything from configuration
│   ├── mount.go                    # Mount(...) - hands each surface the services it calls
│   ├── workers.go                  # Workers() - guards first, then clocks
│   └── close.go                    # Close() - closes everything it built
├── service/                        # The use cases - see daemon.md
│   ├── service.go                  # one set: Service struct, its code here
│   ├── new.go
│   └── <operation>.go              # One file per operation the surfaces call
├── store/                          # Persistence (if needed) - see daemon.md
│   ├── store.go                    # one store: Store struct
│   ├── new.go                      # Constructor: takes the mapper it is given
│   └── <operation>.go              # One file per operation (save.go, resolve.go, by_name.go)
├── face/                           # One full interface per client - see daemon.md
├── migrate/                        # AutoMigrate(d) - schema history, see database.md
├── constant/
│   └── constant.go                 # Tool constants (bucket names, identity)
├── unit/
│   ├── option_test.go              # assert.NotNil(t, option.New()) - see test-placement.md
│   ├── guard_test.go               # Guard battery over Mount - see testing/served.md
│   └── <noun>_tester/new.go        # Builds one service from its parts - see testing/double.md
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
only along consistency boundaries (`store/<aggregate>/`, see `daemon.md`). Split
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
o.Read(a)
Run(o)
```

Every serving daemon's option struct embeds soil's `option.Web`
(`pkg/web/option`), so `o.Address`, `o.ServiceTokens` and `o.Authorization`
read as its own fields. `Read` fills the listen address from `a.Address()` and
the tokens from `SERVICE_TOKEN`. A daemon with sign-on defaults
`Authorization` to `client.NewEnvironment` in `option.New()`; the client reads
`AUTHORIZATION_ISSUER`, its identifier, secret and encryption secret, and
`PUBLIC_LOCATOR` itself.

Every lifecycle server takes the one shape
`server.New(identity, &o.Web, r, setup)`, and the option's values decide the
guarding: its tokens guard `Token` and `Session` routes, and with no tokens - a
metrics exporter's `o.Metric` - the guard takes `Open` routes only and refuses a
protected one at mount, so nothing private can start unguarded. A second
listener is a named `option.Web` on the daemon's option, filled in `main` or
defaulted in `option.New()` (`o.Metric`, gobotd's `o.Public`, goalpined's
`o.Files`), and what it serves mounts itself -
`server.New(constant.Identity, &o.Metric, r, m.Mount)`, `d.MountFiles` - so
`run.go` only wires. A set `Authorization` adds sign-on, built with the reporter
`r` and closed with the server. The test battery builds its guard the same way:
`model_context_server.New` with the test token, and its `WithSignOn()` option
for a daemon that signs on.

`a.Web()` registers the listen flags, `a.Metric()` a Prometheus exporter's,
and `a.Database()` the database's (`a.Lite()` alone for sqlite-only raw stores;
selection in `database.md`). Each flag resolves flag > environment > default,
computed through `environment.Fallback`, so `--help` shows the live resolved
default. The option carries the joined listen address, never a bare port.
Credentials travel via environment, not argv.

- Local daemons are loopback-only by default; each gets a unique `--port`
  (and `--metric-port`) in the Procfile
- Kubernetes deployments set `BIND_ADDRESS: '0.0.0.0'` in their environment
  ConfigMap and keep the port default
- A local daemon serving clients on other hosts passes
  `--bind-address 0.0.0.0` in its launch entry
- The exporter binds all interfaces unless its bind address names one; a host
  reachable from outside its network narrows it to a private address

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
    c := lifecycle.New(l).WithWorkers(d.Workers()...)
    c.WithServer(
        server.New(
            constant.Identity,
            &o.Web,
            r,
            func(g *guard.Mux) {
                d.Mount(r, i.Recorder(), o.Version, g)
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
- `server.New` builds the one `guard.Mux` from the option and hands it to the
  setup, which hands it to `d.Mount` - routes register there, never on the mux.
  A guard with tokens mounts `route:/api/reachable` itself, the peer probe
  route (`error-handling/infrastructure.md`, The probe route)
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
    web.New(d.book, d.loan, d.member).Mount(g)                // web - web.md
}
```

`Mount` is the production surface: the guard battery
(`<path>/unit/guard_test.go`, see `testing/served.md`) starts it on
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

`web.EncodeNotation` writes the JSON header and body; `web.Encode` and
`web.ObjectHeader` split them for a custom status code between. Query
parameter names are `argument.*` constants; response structs sit in
`response.go`.

## MCP Integration

When a daemon also exposes MCP tools, add a `model_context/`
subpackage and mount it through the daemon's `Mount()` on the same
`guard.Mux` as the REST API. See `model-context.md`.

`<path>/model_context/` is the standard package name - not `<path>/tool/` or
`<path>/toolset/`.
