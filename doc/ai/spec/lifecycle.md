# Lifecycle Spec

Reusable application lifecycle manager for service-style applications that serve
HTTP routes and run background workers. Handles startup ordering and
reverse-order shutdown. It lives in `pkg/lifecycle/`.

## Core Concept

Components (servers and workers) are registered through builder methods on
the lifecycle. They start in registration order and close in reverse
registration order. `RunUntilSignal()` handles the run/block/close sequence.

The lifecycle owns what is registered with it: `Close()` releases every
component, so a server or worker handed to `WithServer` or `WithWorker` is
not closed again by the caller.

## File Layout

```
pkg/lifecycle/
  constant/
    constant.go                       - grace period, interrupt margin,
                                        ReadWriteTimeout for protected servers
  lifecycle.go                        - Lifecycle struct
  new.go                              - New(logger) constructor
  run.go                              - Run() starts all components in order
  run_until_signal.go                 - RunUntilSignal() runs, blocks, closes
  close.go                            - Close() drains and closes in reverse order
  with_server.go                      - WithServer(*server.Server)
  with_worker.go                      - WithWorker(face.Worker)
  with_grace.go                       - WithGrace(time.Duration)
  worker/
    worker.go               - adapter from face.Worker to face.Component,
                              absorbing a second Close
  server/
    server.go               - Server struct
    new.go                  - New(identity, address, setup) constructor
    start.go                - sets up mux, registers pprof if enabled,
                              applies middleware, creates http.Server,
                              serves TLS or plain
    close.go                - graceful HTTP shutdown, then the listener
    with_certificate.go     - WithCertificate(cert, key) - enables TLS/HTTP/2
    with_listener.go        - WithListener(net.Listener) - for tests
    with_middleware.go      - WithMiddleware(func(http.Handler) http.Handler)
    with_profiling.go       - WithProfiling() - registers pprof endpoints
    with_protected.go       - WithProtected() - 10s read/write timeout
    with_verbose.go         - WithVerbose(bool) - logs the shutdown
```

## Worker Interface

The `face.Worker` interface from `pkg/face`:

```go
type Worker interface {
    Start()
    Close()
}
```

`Close` is the release, as for every other resource: it ends what the
worker started - goroutines, tickers, stop channels - and whatever the
worker built. A service or client the worker was given belongs to
whoever built it, and the worker never closes it. `Close` must be safe
on a worker that was never started - the lifecycle closes everything
registered, and tests close workers they only call into - so anything
`Close` uses is created in the constructor, not in `Start`. It must also
be safe twice: a worker the daemon's service builds is closed by the
lifecycle and again by the service (`service-tool.md`, Service).

## Draining Workers

A worker whose work must not be cut off mid-flight also implements
`face.Drainer`:

```go
type Drainer interface {
    Finished() <-chan struct{}
    Interrupt()
}
```

For a drainer, `Close` asks it to finish and may return before the work
has: `Finished` closes when the worker has really stopped, and `Interrupt`
is the must-stop. Lifecycle detects the interface at `WithWorker`; plain
workers close as before. Safe on a never-started drainer includes
`Finished` closing on `Close`, or lifecycle waits out the whole grace
period for work that never began.

`Close()` runs in three phases against a deadline of the grace period
minus `InterruptMargin` (5 s), measured from the call. The grace
period is `TerminationGrace` (30 s, the Kubernetes default) unless
`WithGrace(d)` sets another, which is how tests reach the interrupt
path in milliseconds:

1. Every drainer's `Close()`, so it finishes while servers shut down.
2. All components in reverse order. A drainer's second `Close()` is
   absorbed by its adapter.
3. Each drainer's `Finished()` until the deadline, then `Interrupt()`
   and up to `InterruptMargin` more.

The provision runner is the implementation: `Finished` closes when its
loop returns, and `Interrupt` sends SIGINT to every process the
reaper owns, which lets terraform release its state lock before the
kill.

## WithServer

One method for all server configurations. The caller builds a
`*server.Server` using the builder methods and passes it in.

```go
lifecycle.New(l).WithServer(
    server.New(constant.Identity, address, setup).
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
  an address. Used in tests with `system.ClaimPort()`; the server
  takes the listener and closes it.
- `WithVerbose(v)` - logs the graceful shutdown.

Mixed servers (REST routes + MCP/SSE on the same mux) omit
`WithProtected()` - the streaming endpoint decides the timeout.

## Registration Order Matters

Components start in the order they are registered. This is intentional:

```go
// Worker initializes before servers accept traffic
lifecycle.New(l).WithWorker(w).WithServer(srv)
```

Close happens in reverse: the server shuts down before the worker closes.

## What Stays Outside Lifecycle

- **Sentry** - lives in `Main()`, not `Run()`. See `entrypoint.md`. The
  `recover()` defer in `Main()` catches panics from `Run()`.
- **App-specific setup** - database connections and client construction
  happen in the daemon's service, configuration parsing in `Main`. `Run`
  builds the service before `lifecycle.New()`, registers the workers it
  hands over, and defers its `Close` (`service-tool.md`, Service).

## Usage Pattern

```go
func Run(o *option.Config, r face.Reporter) {
    l := logger.New(context.Background())
    lifecycle.New(l).
        WithWorker(worker.New(l, r)).
        WithServer(
            server.New(
                constant.Identity,
                o.Address,
                func(m *http.ServeMux) {
                    m.HandleFunc("/health", health)
                },
            ).WithMiddleware(web.RecoveryMiddleware(r)),
        ).
        RunUntilSignal()
}
```

When registration depends on a condition, keep the lifecycle in a variable
and add to it; `RunUntilSignal` closes it either way.

See `pillars.md` for the full wiring pattern including Main()
reporter creation.

## Adapting Existing Types

To make a type usable as a lifecycle worker:

- **Move runtime parameters into constructor.** If `Start(hourly bool)` takes
  args, move them to `New(..., hourly)` so `Start()` takes none.
- **Move callbacks into options.** If `Start(fn)` takes a callback, add a
  `WithSubscriber(fn)` option so `Start()` reads from the struct.
- **HTTP servers become route registrars.** If a server type owns its own
  `*http.Server`, extract a `Setup(m *http.ServeMux)` method and let lifecycle
  own the HTTP serving via `WithServer`.

## Soil Components with Start and Close

These types implement the Worker interface:

| Package             | Type     | What it does                |
|---------------------|----------|-----------------------------|
| `pkg/ticker`        | `Ticker` | Periodic function execution |
| `pkg/system/reaper` | `Reaper` | Zombie process reaper       |

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
lifecycle.New(l).
    WithWorker(reaper.New(r)).
    WithWorker(provisionRunner).
    WithServer(srv)
```

It is also the process registry `run` records its children in, and
`Interrupt()` sends each of them SIGINT.

Only needed for services that exec subprocesses. Pure HTTP
clients and MCP bridges do not need it.
