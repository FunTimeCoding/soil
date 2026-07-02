# Error Handling

## Principle

Minimize `error` returns and `if e != nil` comparisons. Prefer `errors.PanicOnError`
and rely on recovery chains to capture unexpected failures. Explicit error
handling is reserved for cases where it is structurally required.

## Recovery Chain

Three recovery layers wrap all program execution:

**Entrypoint** (`Main()`): `defer func() { r.RecoverFlush(recover()) }()` captures
panics from the entire program. See `entrypoint.md`.

**HTTP middleware** (`web.RecoveryMiddleware`): wraps the HTTP mux. Panics from web
handlers are caught, reported via `r.Recover(v)`, and converted to 500 responses.

**Worker goroutines**: per-iteration work runs through the shared
`recovery.Recovery` component (`pkg/errors/sentry/recovery`). Panics are
reported via the reporter and the worker continues running (the loop is
not killed).

```go
w.recovery.Run(w.poll)
```

The reporter is never nil - `Main()` always creates one. Empty locator
produces noop behavior. No nil-guards needed.

## When to Use Each Strategy

### Default: `PanicOnError`

Use everywhere covered by a recovery layer - web handlers, workers, REST route
handlers. Do not return the error to the caller. Do not write `if e != nil`.

```go
errors.PanicOnError(s.store.Save(record))
```

### Flow control exception: return `error`

Use when the error outcome changes what happens next - not just "something went wrong"
but "this specific failure path has distinct handling." The canonical example is a
job worker marking status transitions: a failed status update means the job must be
skipped, which requires the caller to act differently.

```go
if e := w.store.UpdateStatus(job.ID, "processing"); e != nil {
    w.logger.Plain("failed to mark job %d processing: %v", job.ID, e)

    return
}
```

### Multi-context exception: return error for caller to decide

When a function is called from both HTTP handlers (which PanicOnError) and MCP
handlers (which translate to tool results), it returns `error`. Each caller
handles it according to its context.

```go
// Shared domain function - returns error
func ProcessRecord(path string) (*Result, error) {
    // ... filesystem or network operations that can fail
}

// HTTP caller - panics, caught by recovery middleware
result, e := store.ProcessRecord(path)
errors.PanicOnError(e)

// MCP caller - captures to reporter, returns structured error
result, e := store.ProcessRecord(path)

if e != nil {
    return s.captureFail(e, "process record failed")
}
```

This is distinct from the flow control exception - the error doesn't change
what happens next (both callers abort on error). The difference is how the
error is surfaced.

## Leaves

- [mcp.md](mcp.md) - MCP tiers, captureFail, captureDetail, detail_error, parseDetail
- [rest.md](rest.md) - strict server tiers, clientError, captureFail on REST
- [external-api.md](external-api.md) - captureDetail pattern, default posture, sentinel classification, typed errors, message format
- [infrastructure.md](infrastructure.md) - store method rule, worker recovery, HTTP recovery, Sentry enrichment, external process, self-healing

## Deviations

**vs idiomatic Go**: idiomatic Go returns `(T, error)` from nearly everything. Here,
error returns are the exception, not the rule. See `conventions.md`.

**vs Claude defaults**: Claude will tend to add `if e != nil` handling everywhere and
thread errors up the call stack. Resist this. When in doubt, `PanicOnError`.
