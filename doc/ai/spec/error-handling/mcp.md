# MCP error handling

MCP handlers have no recovery middleware. Panics would crash the handler without
producing a response. All errors must be translated to tool results.

Do not add per-handler recover defers to MCP tools - the mcp-go framework
handles recovery internally.

Two tiers:

**Tier 1 - Input validation** (bad params from the model): use `response.Fail`. No
Sentry - these are model mistakes, not infrastructure failures.

```go
id, f := r.RequireString(parameter.Identifier)

if f != nil {
    return response.Fail("identifier is required: %v", f)
}
```

`response.Fail` wraps `mcp.NewToolResultError` with `fmt.Sprintf` and returns the
standard `(*mcp.CallToolResult, error)` tuple. Use it for all input validation.

**Tier 2 - Infrastructure failure** (store, DB, external call): capture
to Sentry and return a structured error with the event ID. Two layers:

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

Each service's `captureDetail` checks for different error types
depending on its API:

- Sentry, Habitica, Jellyfin, Confluence: `*detail_error.Detail`
  from `parseDetail` in the HTTP client
- GitLab: `*gitlab.ErrorResponse` → `e.Message`
- Mattermost: `*model.AppError` → `e.Message`
- NetBox: `*netbox.GenericOpenAPIError` → `common.ExtractMessage`
  parses `detail` string, then field-level validation errors
  (`__all__` and per-field arrays). Shared between REST and MCP
  surfaces via `gonetboxd/common/`.
- Jira/Confluence (goatlassiand): `*jira.Error` → `jiraMessage`
  (joins all ErrorMessages + Errors map with "; "),
  `*detail_error.Detail`
- Proxmox: sentinel errors via `errors.Is` (`ErrNotFound`,
  `ErrNotAuthorized`, `ErrTimeout`)
- GORM services: `gorm.ErrRecordNotFound` via `errors.Is`

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

**`detail_error.Detail`** (`pkg/web/detail_error/`): a shared typed
error carrying a `Detail` string and `Status` string. HTTP clients
return this from `parseDetail` when the API response contains a
known error message field. The `Detail` field is what gets relayed
to the model.

**`constant.UnexpectedError`** (`pkg/constant/`): the shared fallback
message - `"unexpected error"`. Honest about not knowing what went
wrong. The model gets this alongside the sentry event ID and can
look up the full error if needed. Never lie about the cause - don't
use "API unreachable" or "database unreachable" as catch-all messages.

**`parseDetail`** in HTTP clients: checks the response status code
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

`response.CaptureFail` captures the exception via the reporter and
returns structured JSON with `error` and `event_identifier` fields via
`response.FailAny`. The event ID lets the model look up the
stacktrace via Sentry MCP tools and diagnose the problem in the same
conversation.

The reporter is threaded into the MCP server (same pattern as web and
workers). The reporter is never nil - it always exists, even in noop
mode.

Design principle: any `error` value from a function call in an MCP
handler is worth capturing. Even local file I/O errors.
`response.Fail` is only for validation where the handler constructs
the error message itself (e.g. "service is required").
