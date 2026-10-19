# External API error handling

Services that wrap external APIs (NetBox, Jira, Confluence, Alertmanager,
Proxmox) meet errors no class answers yet. A NetBox "name must be
unique" is the caller's mistake - but we haven't seen enough of these
errors to know which are safe to exclude from Sentry.

## Default posture: capture everything

Capture everything to Sentry. The status is for the caller and the
capture is for us - independent questions (`README.md`, Answering an
error). Only stop capturing specific errors after they've been seen
enough to be identified as noise - by classifying them, never by
skipping the capture.

## Upstream errors are 500 until classified

We haven't classified which
upstream errors are truly client mistakes vs infra failures vs
auth issues. Until a specific error has been observed enough to
route differently, everything from an external API is 500 with
`ErrorResponse`. The split is message quality, not status code:
parseable errors get the extracted message, opaque errors get
`constant.UnexpectedError`.

Not all upstream APIs return structured errors. Some return
opaque errors or HTML error pages. For those, the fallback
message is all we can offer.

## The captureDetail pattern

A per-service method that checks for the upstream's error types and
extracts the best message.
Returns `*server.ErrorResponse` - always captured; the handler answers
with the status the class names, 500 for anything unclassified.

```go
// gonetboxd - NetBox returns GenericOpenAPIError with parseable bodies
func (s *Server) captureDetail(e error) *server.ErrorResponse {
    if m, okay := common.ExtractMessage(e); okay {
        return s.captureFail(e, m)
    }

    return s.captureFail(e, constant.UnexpectedError)
}
```

The handler returns 500 with the result:

```go
if e != nil {
    return server.Foo500JSONResponse(*s.captureDetail(e)), nil
}
```

**MCP layer equivalent:** MCP handlers use the same `captureDetail`
pattern but return `(*mcp.CallToolResult, error)` via `captureFail`
+ `response.CaptureFail`. The extraction logic is identical - only
the return type differs.

**When to add captureDetail vs just captureFail:** if the service
wraps an external API with parseable errors, use `captureDetail`
to extract the best message. If the service only has internal
stores (GORM, pgx), use `captureFail` with
`constant.UnexpectedError` - there's no meaningful message to
extract for the caller.

## Sentinel classification

The default posture (500, Sentry, best-effort message) applies to
errors we haven't classified yet. When a specific error has been
observed enough in Sentry to understand what it means, classify it:
give it the shared class it turned out to be, and the table in
`README.md` (Answering an error) answers its status, its capture and
its message together.

Classification is per-error. The handler checks for the classified
error before falling through to `captureDetail`. The default remains
capture-everything until proven otherwise.

### Simple sentinels and typed errors

Two shapes for classified errors.

*Simple sentinel* -
`var ErrorBrowserUnreachable = errors.New("browser unreachable")`. Fixed
message, matched with `errors.Is`. Use when the message doesn't vary by
instance - the condition is the full story.

*Typed error* - a struct carrying a formatted message, matched
with `errors.As`. Use when the message needs instance context
(which resource, which identifier). The type provides
classification, the instance carries the specifics.

```go
type NotFoundError struct {
    Message string
}

func (e *NotFoundError) Error() string { return e.Message }

func Is(e error) bool {
    var target *NotFoundError

    return errors.As(e, &target)
}
```

Handlers that only classify - route to 404, pick a response
shape - call the package's `Is` function (the `os.IsNotExist`
shape) instead of spelling out `errors.As` with a target they
never read. No package-level sentinel variable exists for typed
errors; the type is the classification.

The dividing line: does the message need context from the call
site? If yes, typed error. If no, simple sentinel.

### Shared error classes

The shared typed classes live in `pkg/errors`, each with
`Format` (or `New`) and `Is`. Reach for these before inventing
a per-package sentinel for the same class:

- `not_found` - a lookup missed. `New(kind, identifier)` for
  the standard message, `Format` for free-form.
- `ambiguous` - an exactly-one lookup matched several. Zero
  matches is `not_found`, never `ambiguous`.
- `conflict` - the state is already occupied: already live,
  already in use, already queued, still running.
  `Exists(kind, identifier)` renders the standard
  "kind already exists: identifier" message.
- `validation` - a caller-shaped request error (`New`);
  `captureDetail` surfaces its message without a Sentry
  capture. Upstream data failing our own checks is not the
  caller's mistake and stays unclassified.
- `not_selected` - a required session selection is missing;
  the message carries the actionable hint.
- `not_configured` - the deployment lacks the config or the
  installation to serve the operation (a cluster without the CRD
  a tool reads); only the operator can fix it, so it is captured.
- `unreachable` - a backing dependency is down or not connected.
- `unavailable` - a peer guard's verdict that the peer is down;
  the guard has reported it, so a caller never does
  (`infrastructure.md`, Peer Outages and Refusals).
- `timeout` - waiting ran out. `timeout.Is` also matches
  `context.DeadlineExceeded`.
- `dropped` - the answer broke off mid-way: an unexpected EOF, a
  reset, a closed HTTP/2 stream. Retried like a timeout on idempotent
  requests.
- `unexpected` - an upstream replied with a surprise; message
  shape `"<subject> status: %d"` or `"unexpected <subject>:
  <value>"`. The package also carries the panic helpers
  (`Count`, `Integer`, ...).
- `unreadable_body` - a response arrived but reading its body
  failed. `New(wrapped, format, ...)` keeps the cause chain;
  `web.ReadBytes` panics with it, so the panic path classifies
  too. An upstream's own well-formed error message is neither
  this nor `unexpected` - pass it through.

Multi-instance daemons resolve their session-scoped instance
through `go:pkg/inventory.Resolve`, which produces `not_found`,
`not_selected` and `not_configured`. Every session-scoped choice - an
instance, a cluster, a module - is held by
`go:pkg/generative/model_context/selection`, which looks the session up
once; a context without a session is a broken invariant, never a class.

### Translate at the boundary

A driver's or library's own sentinel - `gorm.ErrRecordNotFound`,
`proxmox.ErrNotFound`, a Kubernetes `IsNotFound` - is translated into
its shared class at the call that knows what it was looking for, the
store or client method, as `not_found.New(kind, identifier)`. No surface
tests a driver sentinel, and no store returns one it did not receive:
the surface sees only classes, and the message names the subject and the
identifier instead of the driver's "record not found".

## Message format

An error message reads `"subject condition: identifier"` - the what
before the which. Examples: `"machine not found: 123"`,
`"collection not found: favorites"`, `"browser unreachable"`.
The subject tells you what failed. The identifier tells you
which one. Construct the message where the knowledge is - the
function that did the lookup, not the handler that routes the
response.

When wrapping an underlying error with the operation that was
attempted, the wrap label is `"verb noun:
%w"` - `"list tabs: %w"`, `"create entry: %w"`, `"download
vocabulary: %w"`. Never `"failed to X"` or `"X fail"` - the
error position already says it failed.
