# Infrastructure error handling

## Store Method Rule

Three caller categories determine how store methods handle errors:

| Caller | Store returns | Caller handles |
|--------|-------------|----------------|
| Web handler (HTML, recovery middleware) | `error` | `PanicOnError` - recovery middleware catches, returns 500 |
| Worker / sweep | void, `PanicOnError` inside | `recovery.Run` catches, worker continues |
| MCP handler | `error` | `captureFail` / `captureDetail` → structured tool result |
| REST handler (strict server) | `error` | `captureFail` → typed 500 response with Sentry event ID |

The default direction is: store methods return `error`. Web callers wrap
with `PanicOnError`; MCP and REST callers handle explicitly. Worker-only
store methods may `PanicOnError` internally when they have no other
callers.

```go
// Returns error - called from MCP, REST, and web
func (s *Store) UpdateStatus(identifier string, status string) error {
    return s.database.Model(&Record{}).
        Where("identifier = ?", identifier).
        Update("status", status).Error
}

// Web caller - panics, caught by recovery middleware
errors.PanicOnError(s.store.UpdateStatus(id, status))

// MCP caller - captures to Sentry, returns structured error
if e := s.store.UpdateStatus(id, status); e != nil {
    return s.captureFail(e, "update status failed")
}

// REST caller (strict server) - returns typed error response
if e := s.service.UpdateStatus(id, status); e != nil {
    return server.PostUpdate500JSONResponse(
        *s.captureFail(e, constant.UnexpectedError),
    ), nil
}
```

Web handlers rely on panic + recovery middleware because there is no
structured error rendering in the web UI yet. This is intentional -
the panic path is correct and visible via Sentry until a proper error
display mechanism (toast/notification) is built into `pkg/web/layout`.

## Worker Recovery Pattern

Workers that loop (pollers, watchers, schedulers) wrap per-iteration work
in the shared `recovery.Recovery` component
(`pkg/errors/sentry/recovery`). It catches panics from a single
iteration, reports via `r.Recover(v)`, logs, and lets the loop continue.

The worker constructor builds it from the logger and reporter it
already receives:

```go
func New(
    v *service.Service,
    interval time.Duration,
    l *logger.Logger,
    r face.Reporter,
) *Worker {
    return &Worker{
        service:  v,
        interval: interval,
        recovery: recovery.New(l, r),
        stop:     make(chan struct{}),
    }
}
```

Usage in a loop:

```go
for {
    select {
    case <-t.C:
        w.recovery.Run(w.poll)
    case <-w.stop:
        return
    }
}
```

See `pillars.md` for where this fits in the overall recovery
layer design.

## HTTP Recovery Middleware

`pkg/web/RecoveryMiddleware` is the shared HTTP recovery layer. It wraps
the mux, catches panics from any handler, reports via `r.Recover(v)`,
and returns 500. Wired into lifecycle via
`server.New(...).WithMiddleware(web.RecoveryMiddleware(r))`.
See `lifecycle.md`.

### Sentry body enrichment

`pkg/errors/sentry/start.go` installs a `BeforeSend` hook that
checks every error for `face.BodyProvider` (an interface with
`Body() []byte`). When present - e.g. `*netbox.GenericOpenAPIError`
- the response body is attached as a `response` context on the
Sentry event. This works for both `CaptureException` and `Recover`
paths (via `OriginalException` and `RecoveredException` in the
`EventHint`). No caller changes needed - any error satisfying
`BodyProvider` is automatically enriched.

## External Process Failure

When shelling out to an external process (terraform, git, ansible),
use `SetReporter` on `run.Run` to enrich Sentry
with stdout and stderr automatically on failure:

```go
c := run.New()
c.Directory = directory
c.SetReporter(r.reporter, "terraform init")
c.Start("terraform", "init")
```

`SetReporter` stores the reporter and a label on the `Run` struct.
When `Start` encounters an error and `Panic` is true (the default),
it calls `CaptureWithContext` with the process output before
panicking. The Sentry event gets stdout and stderr as structured
context - visible in the Sentry UI without being stuffed into the
panic message.

For cases where the caller needs to inspect the error before
deciding what to do (e.g. parsing terraform's JSON output to
decide whether to retry with `-upgrade`), use `NoPanic()` and
handle manually:

```go
c := run.New().NoPanic()
c.Start("terraform", "init", "-json")

if c.Error != nil && needsUpgrade(c.OutputString) {
    // retry with different args
}
```

## Self-Healing File Operations

In batch operations (cleanup loops, file processing), file remove failures
are captured without panicking. The file stays on disk and gets cleaned
up on the next run.

```go
if e := os.Remove(path); e != nil {
    c.reporter.CaptureException(e)
}
```

This is the exception to the PanicOnError default - the failure is
transient, self-healing, and not worth killing the batch over. But it
should still be visible in Sentry.
