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

Two tiers:

## Tier 1: input validation

Bad params from the model: use `response.Fail`.
No Sentry - these are model mistakes, not infrastructure failures.

```go
id, f := r.RequireString(parameter.Identifier)

if f != nil {
    return response.Fail("identifier is required: %v", f)
}
```

`response.Fail` wraps `mcp.NewToolResultError` with `fmt.Sprintf` and returns
the standard `(*mcp.CallToolResult, error)` tuple. Use it for all input
validation.

## Tier 2: infrastructure failure

Store, DB, external call: capture to Sentry and return a structured
error with the event ID.

Any `error` value from a function call in an MCP handler is worth
capturing. Even local file I/O errors. `response.Fail` is only for
validation where the handler constructs the error message itself
(e.g. "service is required").

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

`captureDetail` is the per-service wrapper that whitelists known error
types. It checks for API-specific typed errors, relays their message
to the model if recognized, and falls back to `constant.UnexpectedError`
for anything unknown. Lives in its own file (`capture_detail.go`).

```go
func (s *Server) captureDetail(e error) (*mcp.CallToolResult, error) {
    if d, okay := errors.AsType[*detail_error.Detail](e); okay {
        return s.captureFail(e, d.Detail)
    }

    return s.captureFail(e, constant.UnexpectedError)
}
```

Each service's `captureDetail` whitelists the error types its own
client produces - the inventory is the `capture_detail.go` file in
every `<path>/model_context/` package, not this list. The shapes
that recur:

- a typed detail from the HTTP client, `*detail_error.Detail` from
  `parseDetail` or a `wrapError` - the message field is relayed
- a caller-contract error, `*validation.Detail` - relayed the same
  way when the cause is upstream data rather than model input
- an SDK error type with a message field, such as Mattermost's
  `*model.AppError`
- sentinel errors matched with `errors.Is` - a driver's not-found,
  a context deadline, a client's own sentinels
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
