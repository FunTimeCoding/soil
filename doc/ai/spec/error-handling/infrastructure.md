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

Web handlers rely on panic + recovery middleware. Services whose
web package exposes a `Recovery` delegate (wrapping `view.Recovery`)
render the panic into the page layout: a notification item carrying
the panic message and the Sentry event ID - a full shell page for
page loads, a marked fragment for HTMX requests (the layout's
global `htmx:responseError` listener inserts it into the
notification region). Services not yet migrated fall back to
`web.RecoveryMiddleware`, which returns a plain-text 500. Handler
style is unchanged either way: `PanicOnError`, never explicit error
rendering. Validation errors that should re-render a form are not
panics - see the form round-trip convention in `service-tool.md`.

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

### Sentry enrichment providers

`pkg/errors/sentry/start.go` installs a `BeforeSend` hook that
checks every error - on both `CaptureException` and `Recover`
paths (via `OriginalException` and `RecoveredException` in the
`EventHint`) - for two provider interfaces. No caller changes
needed; any error satisfying a provider is automatically
enriched:

- `face.BodyProvider` (`Body() []byte`) - e.g.
  `*netbox.GenericOpenAPIError` - attaches the response body as
  a `response` context.
- `face.ContextProvider` (`ErrorContext() (string, map[string]any)`)
  - e.g. `*command.CommandError` ("process"), `*job.JobError`
  ("job"), `*detail_error.Detail` ("upstream": status, detail,
  body) - attaches the map under the key the error names.

This is the enrichment mechanism: errors carry their own story,
and the single capture at the recovery boundary attaches it.
Reporting from a throw site before panicking creates a second
Sentry issue for the same failure - a context-free decoy - so
never capture and panic with the same error.

## External Process Failure

When a command run through `run.Start` fails, the error is a
`*command.CommandError` (`pkg/errors/command`) carrying the
command line, stdout, and stderr, wrapping the underlying
`exec.ExitError`. Its `Error()` stays short and stable for
Sentry grouping ("git reset --hard origin/main: exit status
128"); the output travels as structured context via
`face.ContextProvider` and is attached by the `BeforeSend`
hook at whichever recovery layer catches the panic:

```go
c := run.New()
c.Directory = directory
c.Start("terraform", "init")
// on failure: panics with *command.CommandError - the recovery
// chain reports once, with command/output/stderr attached
```

For cases where the caller needs to inspect the error before
deciding what to do (e.g. parsing terraform's JSON output to
decide whether to retry with `-upgrade`), use `NoPanic()` and
handle manually - `c.Error` holds the same rich error:

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
