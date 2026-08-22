# External API error handling

Services that wrap external APIs (NetBox, Jira, Confluence, Alertmanager,
Proxmox) encounter errors that don't fit cleanly into tier 1 or tier 2.
A NetBox "name must be unique" is the caller's mistake - but we haven't
seen enough of these errors to know which are safe to exclude from Sentry.

**Default posture: capture everything to Sentry.** The 400/500 split
is for the caller. The Sentry capture is for us. These are independent
decisions. An error can be a 400 to the caller (they sent a duplicate
name) AND captured to Sentry (we want visibility into what errors our
services encounter). Only stop capturing specific errors after they've
been seen enough to be identified as noise - by adding sentinel-based
exclusions, not by skipping the tier.

**Upstream errors are always 500.** We haven't classified which
upstream errors are truly client mistakes vs infra failures vs
auth issues. Until a specific error has been observed enough to
route differently, everything from an external API is 500 with
`ErrorResponse`. The split is message quality, not status code:
parseable errors get the extracted message, opaque errors get
`constant.UnexpectedError`.

Not all upstream APIs return structured errors. Some return
opaque errors or HTML error pages. For those, the fallback
message is all we can offer.

**The captureDetail pattern:** a per-service method that checks
for the upstream's error types and extracts the best message.
Returns `*server.ErrorResponse` - always 500, always Sentry.

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
observed enough in Sentry to understand what it means, classify it
and route it specifically. What changes depends on what we learned:

- Status code: route to 404 (resource doesn't exist), 503
  (dependency down), or 400 (confirmed caller mistake)
- Sentry capture: skip for errors confirmed as caller noise
- Message: use the classified error's message instead of
  extraction

Classification is per-error, not per-tier. The handler checks
for the classified error before falling through to
`captureDetail`. The default remains capture-everything until
proven otherwise.

**Two shapes for classified errors:**

*Simple sentinel* - `var ErrorBrowserUnreachable = errors.New("browser unreachable")`.
Fixed message, matched with `errors.Is`. Use when the message
doesn't vary by instance - the condition is the full story.

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
  capture.
- `not_selected` - a required session selection is missing;
  the message carries the actionable hint.
- `not_configured` - the deployment lacks the config to serve
  the operation; only the operator can fix it.
- `unreachable` - a backing dependency is down, not connected,
  or not installed. The 503-shaped class.
- `timeout` - waiting ran out. `timeout.Is` also matches
  `context.DeadlineExceeded`.
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
through `pkg/inventory.Resolve`, which produces `not_found`
and `not_selected`.

**Message format:** `"subject condition: identifier"` - the what
before the which. Examples: `"machine not found: 123"`,
`"collection not found: favorites"`, `"browser unreachable"`.
The subject tells you what failed. The identifier tells you
which one. Construct the message where the knowledge is - the
function that did the lookup, not the handler that routes the
response.

**Wrap-label format:** when wrapping an underlying error with
the operation that was attempted, the label is `"verb noun:
%w"` - `"list tabs: %w"`, `"create entry: %w"`, `"download
vocabulary: %w"`. Never `"failed to X"` or `"X fail"` - the
error position already says it failed.
