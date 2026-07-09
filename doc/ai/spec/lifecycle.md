# Lifecycle Spec

Reusable application lifecycle manager for service-style applications that serve HTTP routes and run background workers. Handles startup ordering and reverse-order shutdown.

## Package

`pkg/lifecycle/`

## Core Concept

Components (servers and workers) are registered via functional options. They start in registration order and stop in reverse registration order. `RunUntilSignal()` handles the run/block/stop sequence.

## File Layout

```
pkg/lifecycle/
  constant.go                         - Option type
  lifecycle.go                        - Lifecycle struct
  new.go                              - constructor
  run.go                              - Run() starts all components in order
  run_until_signal.go                 - RunUntilSignal() runs, blocks, stops
  stop.go                             - Stop() stops all in reverse order
  with_server.go                      - WithServer option (takes *server.Server)
  with_verbose.go                     - WithVerbose option
  with_worker.go                      - WithWorker option
  component/
    component.go            - internal component adapter
    start.go
    stop.go
  server/
    constant.go             - readWriteTimeout for protected servers
    server.go               - Server struct
    new.go                  - New(address, setup) constructor
    start.go                - sets up mux, registers pprof if enabled,
                              applies middleware, creates http.Server,
                              serves TLS or plain
    stop.go                 - graceful HTTP shutdown
    with_certificate.go     - WithCertificate(cert, key) - enables TLS/HTTP/2
    with_listener.go        - WithListener(net.Listener) - for tests
    with_middleware.go       - WithMiddleware(func(http.Handler) http.Handler)
    with_profiling.go        - WithProfiling() - registers pprof endpoints
    with_protected.go        - WithProtected() - 10s read/write timeout
```

## Worker Interface

The `face.Worker` interface from `pkg/face`:

```go
type Worker interface {
    Start()
    Stop()
}
```

Implemented by: `metric.Server`, `ticker.Ticker`, `reporter.Reporter`, and any struct with `Start()`/`Stop()`.

## WithServer

One option for all server configurations. The caller builds a
`*server.Server` using the builder methods and passes it in.

```go
lifecycle.WithServer(
    server.New(address, setup).
        WithMiddleware(web.RecoveryMiddleware(r)),
)
```

Builder methods on `*server.Server`:
- `WithMiddleware(fn)` - wraps the mux handler. Typically
  `web.RecoveryMiddleware(r)` for panic recovery + sentry.
- `WithProtected()` - adds 10s read/write timeout (slowloris
  protection). Use for plain REST servers that don't serve
  streaming or long-lived connections.
- `WithCertificate(cert, key)` - enables TLS. Go enables HTTP/2
  automatically over TLS.
- `WithProfiling()` - registers pprof endpoints at `/debug/pprof/`.
- `WithListener(l)` - serves on a pre-bound listener instead of
  an address. Used in tests with `system.ClaimPort()`.

Mixed servers (REST routes + MCP/SSE on the same mux) omit
`WithProtected()` - the streaming endpoint governs the timeout.

## Registration Order Matters

Components start in the order they are registered. This is intentional:

```go
// Worker initializes before servers accept traffic
lifecycle.New(
    lifecycle.WithWorker(w),
    lifecycle.WithServer(srv),
)
```

Stop happens in reverse: the server shuts down before the worker stops.

## What Stays Outside Lifecycle

- **Sentry** - lives in `Main()`, not `Run()`. See `entrypoint.md`. The `recover()` defer in `Main()` catches panics from `Run()`.
- **App-specific setup** - database connections, client construction, configuration parsing. All happen before `lifecycle.New()`.

## Usage Pattern

```go
func Run(o *option.Config, r face.Reporter) {
    l := logger.New(context.Background())
    lifecycle.New(
        l,
        lifecycle.WithWorker(worker.New(l, r)),
        lifecycle.WithServer(
            server.New(
                o.Address,
                func(m *http.ServeMux) {
                    m.HandleFunc("/health", health)
                },
            ).WithMiddleware(web.RecoveryMiddleware(r)),
        ),
    ).RunUntilSignal()
}
```

See `pillars.md` for the full wiring pattern including Main()
reporter creation.

## Adapting Existing Types

To make a type usable as a `lifecycle.Worker`:

- **Move runtime parameters into constructor.** If `Start(hourly bool)` takes args, move them to `New(..., hourly)` so `Start()` takes none.
- **Move callbacks into options.** If `Start(fn)` takes a callback, add a `WithSubscriber(fn)` option so `Start()` reads from the struct.
- **HTTP servers become route registrars.** If a server type owns its own `*http.Server`, extract a `Setup(m *http.ServeMux)` method and let lifecycle own the HTTP serving via `WithServer`.

## Soil Components with Start()/Stop()

These types implement the Worker interface:

| Package                      | Type       | What it does                     |
|------------------------------|------------|----------------------------------|
| `pkg/metric`                 | `Server`   | Prometheus metrics HTTP endpoint |
| `pkg/ticker`                 | `Ticker`   | Periodic function execution      |
| `pkg/errors/sentry/reporter` | `Reporter` | Error capture + panic recovery   |
| `pkg/system/reaper`          | `Reaper`   | Zombie process reaper            |

## Reaper

Services that spawn subprocesses (`git`, `ansible-playbook`,
`terraform`) run as PID 1 inside their container. When a
subprocess spawns its own children (e.g. `git` spawns `ssh`)
and the child outlives its parent, it gets reparented to PID 1.
Go's runtime does not reap adopted orphans, so they accumulate
as zombies.

`reaper.New(reporter)` starts a goroutine that calls
`wait4(-1, WNOHANG)` every 30 seconds and reports each reaped
process to Sentry. Register it as a lifecycle worker before
other workers so it starts reaping before subprocesses begin.

```go
lifecycle.New(
    l,
    lifecycle.WithWorker(reaper.New(r)),
    lifecycle.WithWorker(provisionRunner),
    lifecycle.WithServer(srv),
)
```

Only needed for services that exec subprocesses. Pure HTTP
clients and MCP bridges do not need it.
