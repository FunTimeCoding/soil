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
`ErrorResponse`, and the message is the error's own text.

## The upstream's own message

The message is extracted where the error is made, not at the surface. A
client on the requester (`pkg/web/requester`) reads the upstream's
message out of its error body through `refusal.Reason`; a library client
converts the library's error type into `detail_error.Detail` at the
boundary. The surface answers what arrives through `answer` and
`answerWith` (`mcp.md`, `rest.md`) and never parses an upstream body
itself. Where an upstream answers with an opaque body or an HTML error
page, its status line is all the message there is.

Daemons not yet moved may still carry a `captureDetail` reading a
library's error type at the surface, as gonetboxd's
`common.ExtractMessage` does.

## Sentinel classification

The default posture (500, Sentry, best-effort message) applies to
errors we haven't classified yet. When a specific error has been
observed enough in Sentry to understand what it means, classify it:
give it the shared class it turned out to be, and the table in
`README.md` (Answering an error) answers its status and its capture
together.

Classification is per-error. The handler checks for the classified
error before falling through to `answer`. The default remains
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
- `forbidden` - the daemon refuses the operation as it is
  running: a write on a daemon started without `--arm`. Only
  the operator can change it, by restarting the daemon, so the
  caller is told and nothing is captured.
- `validation` - a caller-shaped request error (`New`),
  answered as a tool failure or a 400 without a Sentry
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
  (`infrastructure.md`, Peer Outages and Refusals). A client maps
  its upstream's 503 to it; the guard takes it as an outage and
  answers with its own verdict.
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
store or client method, as `not_found.New(kind, identifier)`; so is a
delete or update the driver reports as zero rows affected without an
error. No surface tests a driver sentinel.

A client classifies at its lowest layer, from the response status
rather than the error a library decoded - a proxy's HTML 503 decodes to
nothing. `go:pkg/web/refusal/Status` is the one table, applied by the
requester's refusal (`request.WithRefusal`, or the client's default) or
at each call site of a library that returns its response beside the
error; a 401 is the daemon's own credential refused, a captured fault,
never `forbidden`. A transport failure keeps its cause - the client
returns the `connection.Failure` that `connection.Classify` builds,
never an `unreachable` built from the error's text, so the guard, the
Sentry fingerprint and the surface all read the same kind and host. An
upstream client that no daemon of ours stands in front of has a
`Reachable` of its own, one cheap authenticated call, so its guard's
probe proves the token as well as the host.

What an upstream answers is checked, not assumed - by a read through
the deployed tool, a direct call with the daemon's own credentials, or
an example (`../example.md`) with synthetic values - and the same answer
can mean different things per call:

- **A status that means the caller on some calls.** An upstream 400 is
  usually the daemon's own malformed request, so the requester's default
  reads it as `unexpected`. Where one call's 400 is seen to be the
  caller's - an invalid value passed straight through - that request
  carries its own refusal (`request.WithRefusal`) turning it into
  `validation` with the upstream's message; every other call keeps the
  default. A status the upstream sends for the caller's mistake and for
  ours alike stays opt-in per call (Habitica's 400 on allocate); one
  only the caller can cause is classified for every call (LinkAce's 400
  and 422).
- **An answer that means different things per call.** OPNsense answers
  a save that failed without validation details: on set that is an
  unknown identifier, `not_found`; on add it is a request missing its
  wrapper, the client's own fault, left `unexpected`.
- **An answer shared with a fault.** LinkAce answers a note for an
  unknown link with the 403 a token without rights would get, so its
  client looks the link up first: the lookup's 404 is the caller's
  mistake, a 403 after it a real fault.
- **An upstream in the tree.** One we own and that answers imprecisely
  is fixed at its source rather than worked around in its client (the
  iTerm bridge, `graft/py/iterm_bridge`).

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
