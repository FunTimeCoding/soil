# Pillars

Every service wires four infrastructure concerns in a consistent
order. Together they ensure operational visibility, panic capture,
graceful degradation, and usage insight.

## The Pillars

| Pillar | Where constructed | Where threaded |
|--------|-------------------|----------------|
| Reporter | `Main()` via `reporter.New` | `Run()` param -> workers, recovery middleware, model_context |
| Logger | `Run()` via `logger.New(ctx)` | Workers, lifecycle |
| Recovery | `Run()` via middleware + `recovery.New` | HTTP servers, worker loops |
| Telemetry | `Run()` via `telemetry.NewEnvironment()` (inside the server setup callback) | REST strict middleware, MCP recorder hook |

## Vigilance

The pillars are required, not optional. `telemetry.NewEnvironment()`
uses `environment.Required` - the daemon exits at startup when the
telemetry host and port variables are missing. The reporter tolerates
an empty locator in code (noop hub, no nil checks - see below), but
every deployment sets `SENTRY_LOCATOR`; a missing locator is a
deployment mistake, not an operating mode. Fail loud at startup over
running blind.

## Wiring Order

`Main()` creates the reporter unconditionally. Empty locator produces
noop behavior - no branching, no nil. `Run()` receives the reporter
as `face.Reporter` (not on the option struct - option structs hold
configuration, not constructed dependencies). `Run()` creates the
logger and wires everything into lifecycle. The telemetry recorder is
constructed inside the server setup callback, where both surfaces
that consume it live.

```go
func Main(
    version string,
    gitHash string,
    buildDate string,
) {
    r := reporter.New(constant.Identity.Name(), version).Start()
    defer func() { r.RecoverFlush(recover()) }()

    a := argument.NewInstance(constant.Identity)
    // ... register flags
    a.Parse(version, gitHash, buildDate)
    o := option.New()
    // ... populate option fields
    Run(o, r)
}
```

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
                    t := telemetry.NewEnvironment()
                    // REST: strict handler + telemetry middleware
                    // MCP: model_context.New(..., t, ...).Mount(m)
                    // see model-context.md for the full wiring
                },
            ).WithMiddleware(web.RecoveryMiddleware(r)),
        ),
    ).RunUntilSignal()
}
```

## Recovery Layers

Three layers, from outermost to innermost:

1. **Entrypoint** (`Main()`): `defer func() { r.RecoverFlush(recover()) }()`
   catches panics from the entire program. See `entrypoint.md`.

2. **HTTP middleware** (`web.RecoveryMiddleware`): wraps the HTTP mux.
   Panics from handlers are caught, reported via `r.Recover(v)`, and
   converted to 500 responses. Wired via `server.New(...).WithMiddleware(...)`.
   See `lifecycle.md`.

3. **Worker loops** (`recovery.Recovery`): each worker runs
   per-iteration work through the shared recovery component
   (`w.recovery.Run(w.poll)`). Panics are reported via `r.Recover(v)`
   and the worker continues. See `error-handling/infrastructure.md`.

MCP handlers do not need per-handler recover defers - the mcp-go
framework handles recovery internally.

## Telemetry

Every REST operation and MCP tool call is recorded as a baseline
event via `face.Recorder`, feeding gotelemetryd (usage heatmaps by
tool, surface, actor, and outcome).

**REST surface** - an inline `StrictMiddlewareFunc` closure wraps the
generated handlers and calls `web.RecordTelemetry(t, operation, e)`
after each call. The closure is pasted per service and that is the
canon pattern: each generated package declares its own
`StrictHandlerFunc` type, so a shared middleware cannot type-check
against all of them (the old `web.TelemetryMiddleware` was removed
for this reason). If oapi-codegen ever supports emitting the wrapper
inside `generated/`, the closure moves there - until then, paste the
block. Full example in `model-context.md`.

**MCP surface** - `WithRecorder(t)` on the mark server factory
installs an AfterCallTool hook that records every tool call. See
`model-context.md`.

**Fire-and-forget by design** - `send()` silently drops on marshal or
POST failure. No retry, no sentry capture. A telemetry failure must
never hurt the service. Strict at startup (vigilance), tolerant at
runtime.

## Server Configuration

The server builder controls timeout and TLS:

```go
server.New(address, setup).
    WithMiddleware(web.RecoveryMiddleware(r)).  // panic recovery
    WithProtected().                            // 10s read/write timeout
    WithCertificate(cert, key).                 // TLS / HTTP/2
    WithProfiling()                             // pprof endpoints
```

Omit `WithProtected()` for streaming servers (MCP, SSE) - the
streaming endpoint governs the timeout choice. Mixed servers
(REST + MCP on the same mux) omit it for the same reason.

## What Each Component Receives

| Component | Logger | Reporter | Telemetry | Why |
|-----------|--------|----------|-----------|-----|
| Lifecycle workers | Yes (constructor param) | Yes (for recovery.New) | No | Workers need both for recovery logging |
| HTTP route handlers | Via recovery middleware | Via recovery middleware | Via strict middleware | Middleware catches panics and records operations |
| MCP tool handlers | Not directly | Via Server struct + captureFail | Via WithRecorder hook | Tier 2 errors use captureFail -> response.CaptureFail |
| Startup one-shot tasks | Yes (from Run) | Yes (for recovery.New if looping) | No | Same recovery pattern as workers |
