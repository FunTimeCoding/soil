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
`answerMessage` (`mcp.md`, `rest.md`) and never parses an upstream body
itself. Where an upstream answers with an opaque body or an HTML error
page, its status line is all the message there is.

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
- `reported` - a peer daemon answered with a failure it has
  already captured: its body carries a non-empty
  `event_identifier`. `refusal.New` produces it before the status
  table, so a captured 500 never reaches Sentry a second time.
  The message keeps the peer's reason and its event identifier,
  the handle to the one event. An empty identifier means nothing
  reached Sentry, and the caller still captures.
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

A sentinel moving to a shared class moves every branch testing it in
the same change, or each one falls through to a captured 500. An empty
list is not `not_found`: it answers empty.

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

A client returns classified errors. A recover stays only around a
dependency that still panics, turning the panic into a captured error: a
client panicking on every failure, recovered at the service, answers
every failure as one captured 500. A write whose confirmation timed out
against an external party may have happened, so it answers a captured
fault saying so and asking to check before retrying - a blind retry can
repeat it - and a create that cannot find its new identifier answers the
same, never success with an empty one.

A client classifies at its lowest layer, from the response status
rather than the error a library decoded - a proxy's HTML 503 decodes to
nothing. One table classifies every upstream
(`go:pkg/web/refusal/Table`): 404 is `not_found`, 409 `conflict`, 429,
502, 503 and 504 `unavailable` once the retries are spent, and 400, 401,
403, 422 and every other status `unexpected`. A 401 or 403 is the
daemon's own credential or its privileges refused - a captured fault,
never `forbidden` - unless the client knows its upstream's 403 to be
about a resource the caller chose, and then it sets that row. The
requester carries the table (`WithStatus(http.StatusForbidden,
errorsConstant.ForbiddenClass)`); a library that returns its response
beside the error applies it at each call site, `refusal.Status` for the
default or `refusal.NewTable().With(...).Classify` with the client's rows
(GitLab, Jira, Mattermost, Ollama, OpenAI, Alertmanager, NetBox,
Prometheus). A default row its upstream's source contradicts goes with
`Without`: Prometheus answers 503 for a query timeout as well as while
starting, and 404 only from its router, for a path the client wrote, so
neither is `unavailable` or `not_found` there. A body that
changes what a status means - a refused token inside a 403 - is read by
the client's refusal hook (`WithRefusal`), which runs before the table.
Our own daemons answer through `refusal.NewDaemonTable`, where 400 and
422 are `validation` and 403 `forbidden`: a daemon's status is already
its classified answer. A transport failure keeps its cause - the client
returns the `connection.Failure` that `connection.Classify` builds,
never an `unreachable` built from the error's text, so the guard, the
Sentry fingerprint and the surface all read the same kind and host. An
upstream client that no daemon of ours stands in front of has a
`Reachable` of its own, one cheap authenticated call, so its guard's
probe proves the token as well as the host.

What an upstream answers is checked, not assumed - by its source read at
the deployed version, a read through the deployed tool, a direct call
with the daemon's own credentials, or an example (`../example.md`) with
synthetic values; a closed-source peer is probed live with harmless
requests. A class no source or probe shows is not mapped on a guess. The
same answer can mean different things per call:

- **A status that means the caller on some calls.** An upstream 400 is
  usually the daemon's own malformed request, so the requester's default
  reads it as `unexpected`. Where one call's 400 is seen to be the
  caller's - an invalid value passed straight through - that request
  sets its own row (`request.WithStatus(http.StatusBadRequest,
  errorsConstant.ValidationClass)`), turning it into `validation` with
  the upstream's message; every other call keeps the default. A status
  the upstream sends for the caller's mistake and for ours alike stays
  opt-in per call (Habitica's 400 on allocate); one
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
