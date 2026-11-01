# Infrastructure error handling

## Store Methods

Store methods return `error`. Each caller handles it by its context - the
web handler with `PanicOnError` under the recovery middleware, MCP and REST
by the error's class (`README.md`, Multi-context exception and Answering an
error). A store method only a worker calls may `PanicOnError` inside, under
the worker's recovery.

## Web Handler Answers

A web handler answers an error from the service by its class, as REST and
MCP do (`README.md`, Answering an error), through `view.Answer`: it renders
a notification item into the page layout - a full shell page for page
loads, a marked fragment for HTMX requests, which the layout's global
`htmx:responseError` listener inserts into the notification region - at the
class's status. A quiet class carries its message alone and Sentry hears
nothing; anything else is captured and the item carries the event ID.
Validation errors that re-render a form are the form round-trip instead
(`service-tool.md`, Form Errors).

Panics stay the net for faults the handler does not expect. A web package
exposes a `Recovery` delegate (wrapping `view.Recovery`) that renders a
panic the same way at 500, and reads its class: a panicking quiet-class
error is shown without a capture, so a forgotten `PanicOnError` on a peer's
`unavailable` does not flood Sentry. A daemon without a web package uses
`web.RecoveryMiddleware`, a plain-text 500.

## Worker and HTTP Recovery

A worker's loop runs each iteration through `recovery.Recovery`
(`pkg/errors/sentry/recovery`), built in the worker's constructor from the
logger and reporter: `w.recovery.Run(w.poll)` reports a panic and keeps the
loop running. The HTTP servers wrap their mux in the recovery middleware
(`lifecycle.md`, WithServer).

## Peer Outages and Refusals

Everything for talking to another service lives under `pkg/peer/`: the
wait and the guard at its root, `pkg/peer/reacher/` for outage and
refusal edges, `pkg/peer/retry/` for the bounded loops, and
`pkg/peer/constant/` for the timings.

### Required and guarded peers

Every peer is guarded once the daemon runs (The guard). Whether it must
also answer before the daemon serves follows from what calls it.

A peer a surface has no answer without is required: the daemon serves only
once it answers, and fails at start-up when it does not, rather than
answering for work it cannot do. A peer the surfaces can answer without -
one only a worker reads, or one whose absence a surface answers from the
daemon's own store or as `unavailable` (goqueryd's embedder) - is guarded
and not waited on, so a peer that is away stops what needs it while every
other answer keeps going.

The service decides which, per peer, where it builds the guard; it is a
property of how the peer is used, so no flag or variable changes it. The
environment decides only whether a peer exists: `Main` sets its factory
when the environment names it.

`Run` waits for each required peer before the lifecycle starts and panics
when one never comes. `go:pkg/peer/Await` polls the peer's probe for a minute
(`Wait`, every `Poll`), logs `<name>_connected`, and returns why the peer
never came; a guard's `Wait` is the same wait. The wait takes
`context.Background()`, since a signal before the lifecycle ends the
process.

### The probe route

Every token-guarded daemon serves `route:/api/reachable`: `guard.New`
mounts it behind the token, it answers 204 and touches nothing - no
upstream, no store, no browser. It is the token guard's route, not the
daemon's, so no OpenAPI spec declares it. A peer's `face.Reachable` is
one call to it with the caller's token, through
`go:pkg/web/reachable_probe/Reach`, so the wait and the guard prove the
peer accepts this daemon, not only that it runs. A 401 there is the
peer refusing the token: the wait fails with it at start-up, and later
the reacher reports it once as a refusal.

`route:/health` is the lifecycle server's liveness answer for the
orchestrator, the supervisor and load balancers, open by design, and
never a peer probe - it passes with a wrong or expired token.
`route:/api/health` and `route:/api/status` stay whatever a daemon makes
of them.

### Outages

Every failure talking to a peer goes through one `reacher.Reacher` per
daemon, by `Handle`. A transport failure - refused, dropped, timed out,
by `connection.Classify` - is an outage, and so is an `unavailable`
returned by the call, reported as an episode: logged
when it begins, captured to Sentry once if it lasts `Lasting`, as the
transport error wrapped in `go:pkg/peer/outage/OutageError`, logged
again when it ends. A blip a restart explains stays two log lines.
Sentry's connection fingerprint (kind and host) keeps every episode of
one peer in one issue. The reacher takes the daemon's reporter and
notices time only when a call fails: a path that stops calling reports
nothing. Anything else is a refusal, reported once per reason - to
Sentry where it is an anomaly, to the log where it is expected.

### The guard

After start-up every peer sits behind one `go:pkg/peer/Guard` over its probe,
shared by every path that calls it. The service builds the guards and hands them
to `Run` through `Workers()`. `Attempt` runs a call through the reacher: while
the peer is down it fails at once with `errors/unavailable`.

A peer goes down two ways, and each comes back the way it failed. A
transport failure is a peer that does not answer, so the guard probes it
back through `Reachable` on a backoff (`ProbeMinimum` to `ProbeMaximum`).
An `unavailable` from the call itself - the upstream's own 503, or a
guarded daemon passing on its guard's verdict - is a peer that answers
and refuses, where `Reachable` would pass while every operation fails.
That outage is tested by the calls: once the backoff delay has passed,
one call goes through as the trial while every other still fails at
once. A trial that succeeds brings the peer back; one answering
`unavailable` again doubles the delay and checks the episode for
`Lasting`. A peer nobody calls stays down and quiet, since nothing waits
on it.

A caller answers `unavailable` as what it is - 503 over
REST, a tool failure naming the peer - and never reports it again; the
guard has.

`Attempt` hands every other error back untouched, so the caller answers
it by its class: every surface, and every worker that reports its own
failures - returning them from its tick rather than panicking under
recovery, and reporting what `classify.Reportable` admits - calls
through `Attempt`. `Do` is `Attempt` that also reports each distinct
refusal itself, once per reason, without seeing its class, so a quiet
`not_found` or `validation` reaches Sentry; it is for a caller that
reports nothing itself.

### Held connections

A long-lived connection - a WebSocket, an event stream - is held by
`go:pkg/peer/hold/Hold` rather than guarded: there are no calls to fail
fast, and the guard's probe cannot tell whether it reconnected. The
connection is a `face.Connection`: `Open` dials, authenticates and
subscribes, `Serve` reads until the connection ends and says why, `Close`
releases it. The hold runs one loop with its own reacher - a failed `Open`
through `Handle`, an ended `Serve` through `Fail` - and reopens on the
guard's backoff (`ProbeMinimum` to `ProbeMaximum`, back to the minimum once
a session ran), each session under recovery. A peer answering the dial
with an HTTP status maps it through `refusal.New`, so its 503 is an outage
and anything else a refusal.

## Sentry Enrichment Providers

`pkg/errors/sentry/start.go` installs a `BeforeSend` hook that checks
every captured or recovered error for two provider interfaces, so an
error satisfying one is enriched with no caller change:

- `face.BodyProvider` (`Body() []byte`) - e.g.
  `*netbox.GenericOpenAPIError` - attaches the response body as a
  `response` context
- `face.ContextProvider` (`ErrorContext() (string, map[string]any)`) -
  e.g. `*command.CommandError` ("process"), `*job.JobError` ("job"),
  `*detail_error.Detail` ("upstream": status, detail, body) - attaches
  the map under the key the error names

Errors carry their own story and the single capture at the recovery
boundary attaches it. Capturing at a throw site before panicking makes a
second, context-free Sentry issue for the same failure, so never capture
and panic with the same error.

## External Process Failure

A command run through `run.Start` fails with a `*command.CommandError`
(`pkg/errors/command`) carrying the command line, stdout and stderr and
wrapping the `exec.ExitError`. Its `Error()` stays short and stable for
Sentry grouping ("git reset --hard origin/main: exit status 128"); the
output rides as `face.ContextProvider` context. A caller that inspects
the failure before deciding - parsing terraform's JSON output to retry
with `-upgrade` - uses `NoPanic()` and reads `c.Error`, the same error:

```go
c := run.New().NoPanic()
c.Start("terraform", "init", "-json")

if c.Error != nil && needsUpgrade(c.OutputString) {
    // retry with different args
}
```

## Self-Healing File Operations

In a batch - cleanup loops, file processing - a failed file remove is
captured, not panicked: the file stays and the next run removes it. It
is the exception to the `PanicOnError` default, transient and not worth
killing the batch over, but still visible in Sentry.

```go
if e := os.Remove(path); e != nil {
    c.reporter.CaptureException(e)
}
```
