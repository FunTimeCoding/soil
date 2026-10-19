---
base: pkg/tool
---

# MCP error handling

MCP tool handlers run without panic recovery of their own: mcp-go's
recovery is an opt-in server option that is not enabled, so a panic
escapes the handler and is caught by the HTTP recovery middleware,
which answers an HTML 500 the model cannot read as a tool result.
Every error a handler meets must therefore be translated to a tool
result, and no client call a handler makes may panic underneath it -
the Must/non-Must pair in `service-tool.md` exists for this.

Do not add per-handler recover defers as a workaround; the fix is
in the client or in the server builder.

Whether Sentry hears of an error is its class, in the table in
`README.md` (Answering an error); this leaf is how a handler builds
each answer.

## Uncaptured answers

Bad params from the model, and every class the table answers without
Sentry: `response.Fail`.

```go
id, f := r.RequireString(parameter.Identifier)

if f != nil {
    return response.Fail("identifier is required: %v", f)
}
```

`response.Fail` wraps `mcp.NewToolResultError` with `fmt.Sprintf` and returns
the standard `(*mcp.CallToolResult, error)` tuple.

## Captured answers

Every other error a handler meets - a store, a database, an external
call, even local file I/O - is captured to Sentry and answered with a
structured error carrying the event ID.

Two layers:

### The capture primitive

`captureFail` is the primitive - a private method on the Server struct
that wraps `response.CaptureFail`. Takes the error and a model-facing
message. The error goes to sentry, the message goes to the model.

```go
func (s *Server) captureFail(
    e error,
    message string,
) (*mcp.CallToolResult, error) {
    return response.CaptureFail(s.reporter, e, message)
}
```

### Per-service capture

`captureDetail` is the per-service wrapper that relays the message of
recognized API-specific typed errors through `captureFail` and hands
everything else to `response.Answer`, which applies the class table:
an uncaptured class fails with its own message, a captured class is
captured with it, and anything unclassified is captured as
`constant.UnexpectedError`. Lives in its own file (`capture_detail.go`).
A handler routes every error that did not come from parsing the request
through it - `return s.captureDetail(e)` - never through `response.Fail`.

```go
func (s *Server) captureDetail(e error) (*mcp.CallToolResult, error) {
    if d, okay := errors.AsType[*detail_error.Detail](e); okay {
        return s.captureFail(e, d.Detail)
    }

    return response.Answer(s.reporter, e)
}
```

Each service's `captureDetail` whitelists the error types its own
client produces - the inventory is the `capture_detail.go` file in
every `<path>/model_context/` package, not this list. The captured
shapes that recur:

- a typed detail from the HTTP client, `*detail_error.Detail` from
  `parseDetail` or a `wrapError` - the message field is relayed
- an SDK error type with a message field, such as Mattermost's
  `*model.AppError`
- a client's own sentinels matched with `errors.Is` - where one is
  a shared class in disguise (a driver's not-found, a context
  deadline), its class answers it instead
- a shared extractor when REST and MCP surfaces need the same
  parsing, as gonetboxd's `common.ExtractMessage`

### Calling them from a handler

Most handlers call `captureDetail(e)` directly:

```go
result, e := s.client.SearchIssues(...)

if e != nil {
    return s.captureDetail(e)
}
```

Use `captureFail(e, message)` directly only when the handler knows
a specific message that `captureDetail` can't derive - e.g. after
a multi-step operation where the context matters.

## Shared pieces

### The detail error type

`detail_error.Detail` (`pkg/web/detail_error/`) is a shared typed
error carrying a `Detail` string and `Status` string. HTTP clients
return this from `parseDetail` when the API response contains a
known error message field. The `Detail` field is what gets relayed
to the model.

### The unexpected-error fallback

`constant.UnexpectedError` (`pkg/constant/`) is the shared fallback
message - `"unexpected error"`. Honest about not knowing what went
wrong. The model gets this alongside the sentry event ID and can
look up the full error if needed. Never lie about the cause - don't
use "API unreachable" or "database unreachable" as catch-all messages.

### Parsing upstream error bodies

`parseDetail` in HTTP clients checks the response status code
and parses the error body for a known message field. Each API has
its own shape:

- Sentry: `{"detail": "..."}` (JSON)
- Confluence: `{"message": "..."}` (JSON)
- Habitica: `{"message": "..."}` (JSON)
- Jellyfin: `{"title": "..."}` (ASP.NET ProblemDetails JSON),
  or plain text body, or empty
- Salt: `<p>...</p>` in CherryPy HTML error pages

If the known field is found, returns `detail_error.New(message, status)`.
Otherwise returns `fmt.Errorf("%s", status)`.

### Capturing through the reporter

`response.CaptureFail` captures the exception via the reporter and
returns structured JSON with `error` and `event_identifier` fields via
`response.FailAny`. The event ID lets the model look up the
stacktrace via Sentry MCP tools and diagnose the problem in the same
conversation.

The reporter is threaded into the MCP server (same pattern as web and
workers). The reporter is never nil - it always exists, even in noop
mode.
