# Pillars

Every service wires four infrastructure concerns in a consistent
order. Together they ensure operational visibility, panic capture,
graceful degradation, and usage insight.

## The Pillars

| Pillar | Where constructed | Where threaded |
|--------|-------------------|----------------|
| Reporter | `Main()` via `instrument.New` (the reporter half) | `Run()` pulls `i.Reporter()` -> workers, recovery middleware, model_context |
| Logger | `Run()` via `logger.New(ctx)` | Workers, lifecycle |
| Recovery | `Run()` via middleware + `recovery.New` | HTTP servers, worker loops |
| Telemetry | `Main()` via `instrument.New` (the recorder half) | `Run()` pulls `i.Recorder()` -> `Mount()` recording middleware, MCP recorder hook |

## Vigilance

The pillars are required, not optional - every service constructs
all four unconditionally, with no branching on configuration. The
recorder falls back to localhost and the listen default when
`GOTELEMETRY_HOST` and `GOTELEMETRY_PORT` are unset (https unless
`GOTELEMETRY_INSECURE`); the reporter tolerates an empty locator
(noop hub, no nil checks - see below). Every deployment still sets
`SENTRY_LOCATOR` and the telemetry family - absence is a
deployment mistake, not an operating mode.

## Wiring Order

`Main()` creates the instrument unconditionally - `instrument.New`
bundles the reporter and recorder behind one constructor and one
exit defer (see `entrypoint.md`). Empty Sentry locator produces
noop behavior - no branching, no nil. `Run()` receives it as
`face.Instrument` (not on the option struct - option structs hold
configuration, not constructed dependencies), pulls the halves
where it wires them, and creates the logger.

```go
func Main(
    version string,
    gitHash string,
    buildDate string,
) {
    s := instrument.New(constant.Identity, version)
    defer func() { s.Flush(recover()) }()
    a := argument.NewInstance(constant.Identity)
    // ... register flags
    a.Parse(version, gitHash, buildDate)
    o := option.New()
    // ... populate option fields
    Run(o, s)
}
```

```go
func Run(o *option.Config, i face.Instrument) {
    r := i.Reporter()
    l := logger.New(context.Background())
    lifecycle.New(
        l,
        lifecycle.WithWorker(worker.New(l, r)),
        lifecycle.WithServer(
            server.New(
                constant.Identity,
                o.Address,
                func(m *http.ServeMux) {
                    // domain deps first - see service-tool.md
                    Mount(r, i.Recorder(), o.Version, guard.New(m, o.ServiceTokens))
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

**REST surface** - `web.RecordingMiddleware`, instantiated per
service with its generated `StrictHandlerFunc` type, wraps the
generated handlers and calls `web.RecordTelemetry(t, operation, e)`
after each call - the type parameter is what lets the one helper
serve every generated package. Full example in `generated-api.md`.

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
| HTTP route handlers | Via recovery middleware | Via recovery middleware | Via recording middleware | Middleware catches panics and records operations |
| MCP tool handlers | Not directly | Via Server struct + captureFail | Via WithRecorder hook | Tier 2 errors use captureFail -> response.CaptureFail |
| Startup one-shot tasks | Yes (from Run) | Yes (for recovery.New if looping) | No | Same recovery pattern as workers |
