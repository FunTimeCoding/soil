# Error Handling

## Principle

Minimize `error` returns and `if e != nil` comparisons. Prefer
`errors.PanicOnError` and rely on recovery chains to capture unexpected
failures. Explicit error handling is reserved for cases where it is structurally
required.

The same posture for closing: prefer `errors.PanicClose` over checking
close errors; use `LogClose` only in loops or uncertain contexts (e.g.
validation), and `LogReclose` for a value handed to something that may
already have closed it - a file or pipe given to an HTTP request, a
listener a server shuts down. Eliminate error variables entirely when possible:
`errors.PanicOnError(someCall())`. Error variable naming (`e`/`f`/`g`
escalation) lives in `../naming.md`.

## Recovery Chain

Three recovery layers wrap all program execution:

**Entrypoint** (`Main()`): `defer func() { r.RecoverFlush(recover()) }()`
captures panics from the entire program. See `entrypoint.md`.

**HTTP middleware** (`web.RecoveryMiddleware`): wraps the HTTP mux. Panics from
web handlers are caught, reported via `r.Recover(v)`, and converted to 500
responses.

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

### Default: panic on error

Use everywhere covered by a recovery layer - web handlers, workers, REST route
handlers. Do not return the error to the caller. Do not write `if e != nil`.

```go
errors.PanicOnError(s.store.Save(record))
```

### Flow control exception: return the error

Use when the error outcome changes what happens next - not just "something went
wrong" but "this specific failure path has distinct handling." The canonical
example is a job worker marking status transitions: a failed status update means
the job must be skipped, which requires the caller to act differently.

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

// MCP caller - answers by the error's class
result, e := store.ProcessRecord(path)

if e != nil {
    return s.answer(e)
}
```

This is distinct from the flow control exception - the error doesn't change
what happens next (both callers abort on error). The difference is how the
error is surfaced.

## Answering an error

A surface that answers a caller - a REST route, an MCP tool - asks three
questions of every error, and they are independent:

- **Whose move is it?** The status says what the caller does next: fix the
  request, accept that nothing is there, see what already is, try again
  later, or wait while we look.
- **Do we need to look?** Sentry is the anomaly channel. A caller's mistake
  is not ours to look at, and neither is a failure something else has
  already reported; everything else is captured, and nothing captured is
  also logged.
- **What can we honestly say?** The error's own text, after the step that
  failed when the caller cannot infer it - never a guessed cause.

The error's class answers whose move it is and whether we look; the
message does not depend on it:

| Class | REST | MCP | Sentry |
|-------|------|-----|--------|
| `validation`, `ambiguous`, `not_selected` | 400 `Error` | tool failure | no |
| `not_found` | 404 `Error` | tool failure | no |
| `conflict` | 409 `Error` | tool failure | no |
| `unavailable` | 503 `Error` | tool failure | no - the guard reports the outage |
| `unreachable`, `timeout`, `dropped` | 503 `ErrorResponse` | captured | yes |
| `not_configured` | 500 `ErrorResponse` | captured | yes |
| anything else | 500 `ErrorResponse` | captured | yes |

The message is `context: error`, or the error alone without context. A
context names a step the caller cannot infer - which of several writes in
one tool failed - and restating the tool's own name is not one.

The schema follows the Sentry column: a captured answer carries
`ErrorResponse` with its event identifier, an uncaptured one `Error`. On
REST, reaching the 500 is always a fault: a quiet class arriving there is
an operation missing its branch, captured under the title
`missing <class> branch: <error>`, never logged instead.

The table is code: `Answers` in `go:pkg/errors/constant`, read through
`go:pkg/errors/classify` - `Class` names an error's class, `Answer` its
row, `Reportable` its Sentry column, `Message` the text and `Capture` the
REST answer. The `error_answer` analyzer and both surfaces' answers
(`mcp.md`, `rest.md`) read the same table.

`unavailable` is a guard's verdict that a peer is down, and the guard
reports a lasting outage once (`infrastructure.md`, Peer Outages and
Refusals). A transport failure that reaches a surface without a guard in
front has nobody else to report it, so the surface captures it - and one
that keeps arriving names a call path that wants a guard.

The last row is where every error starts. Classifying one moves it into a
row above it - by a shared class from `pkg/errors` (`external-api.md`,
Shared error classes), or for a one-off type by answering
`Expected() true` (`classify/face.Expected`), which `classify.Reportable`
honours - and only once it has been seen enough to know what it means.

## Leaves

- [mcp.md](mcp.md) - answer and answerWith on MCP, response.Fail for
  request parsing, detail_error
- [rest.md](rest.md) - strict server schemas, the 4xx body, answer and
  answerWith on REST
- [external-api.md](external-api.md) - default posture, sentinel
  classification, translating at the client boundary, message format
- [infrastructure.md](infrastructure.md) - store method rule, web handler
  panics, worker recovery, HTTP recovery, Sentry enrichment, external process,
  self-healing

## Deviations

**vs idiomatic Go**: idiomatic Go returns `(T, error)` from nearly everything.
Here, error returns are the exception, not the rule.

**vs Claude defaults**: Claude will tend to add `if e != nil` handling
everywhere and thread errors up the call stack. Resist this. When in doubt,
`PanicOnError`.
