---
base: pkg/tool
---

# MCP error handling

The server builder (`pkg/generative/mark/server`) recovers every tool
and resource handler once, for every daemon: a panic is reported
through the reporter the builder is constructed with and answered as
an error the model can read - a tool with `response.RecoverFail`'s
`error` and `event_identifier`, a resource with an error naming the
event. The reporter is a required argument of `New`, so no server
recovers without reporting.

Recovery makes the answer readable, not the panic fine: a recovered
panic is still a captured anomaly, so the code that panics is a bug
to fix at its cause. Every error a handler meets is translated to a
tool result, and no client call a handler makes panics underneath
it; the Must/non-Must pair in `service-tool.md` exists for this. A
step that can panic is checked before it runs. Do not add per-handler
recover defers; the builder's is the one.

A panic inside a peer guard's `Attempt` unwinds through it before the
builder recovers it: the guard releases a trial it admitted and
records no outcome, since a panic is a bug here rather than evidence
about the peer, so the next caller is admitted
(`TestPanickingTrialLeavesNextTrialAdmitted`).

Whether Sentry hears of an error is its class, in the table in
`README.md` (Answering an error); this leaf is how a handler builds
each answer.

## Request parsing

Bad parameters from the model are answered on the spot with
`response.Fail`, which wraps `mcp.NewToolResultError` with
`fmt.Sprintf` and returns the standard `(*mcp.CallToolResult, error)`
tuple.

```go
id, f := r.RequireString(parameter.Identifier)

if f != nil {
    return response.Fail("identifier is required: %v", f)
}
```

## Everything else: answer and answerMessage

Every other error a handler meets - from a service, a store, a client,
even local file I/O - goes to one of two private methods on the Server
struct, each in its own file:

```go
// answer_message.go
func (s *Server) answerMessage(
    e error,
    message string,
) (*mcp.CallToolResult, error) {
    return response.AnswerMessage(s.reporter, e, message)
}

// answer.go
func (s *Server) answer(e error) (*mcp.CallToolResult, error) {
    return s.answerMessage(e, "")
}
```

`response.AnswerMessage` reads the class table. A quiet class fails the
tool with the message alone; anything captured goes to Sentry and
answers structured JSON carrying `error` and `event_identifier`, so the
model can look the event up through the Sentry tools in the same
conversation. The message follows the rule in `README.md`: the error's
own text, after a context only when the context names a step the caller
cannot infer.

```go
records, e := s.service.Records(name)

if e != nil {
    return s.answer(e)
}
```

`answerMessage` is for a tool that does several things, where the context
says which one failed - adding, removing or replacing tags in one call.
A single-step tool answers `answer(e)`; restating the tool's own name is
not a context.

A step a daemon must run before every answer - gosentryd records the
reacher's outage episodes - sits in `answer` before it hands over, so a
site that needs the step never calls `answerMessage` directly.

An inline class branch before the answer goes when the answer says the
same or better; one that re-words or re-classes the error stays.

Daemons not yet moved still carry `captureFail`, which captures every
error with the message it is given through `response.CaptureFail`, and
`captureDetail`, which hands to `response.Answer` - the same table,
under the old names.

## Shared pieces

### The detail error type

`detail_error.Detail` (`pkg/web/detail_error/`) is a shared typed error
carrying a `Detail` string and `Status` string. Library clients convert
their own error types into it at the boundary (a `wrapError`), and its
`Detail` is the message a surface relays. It implements
`ContextProvider`, so Sentry gets the structure through the `BeforeSend`
hook (`infrastructure.md`, Sentry Enrichment Providers).

### Upstream messages

Clients on the requester (`pkg/web/requester`) read an upstream's own
message out of its error body through `refusal.Reason`, which checks the
fields APIs commonly use (`message`, `detail`, `title`, `error` and
others). A client that knows a status means the caller's mistake on one
call passes a refusal on that request (`external-api.md`, Translate at
the boundary).

### The reporter

The reporter is threaded into the MCP server (same pattern as web and
workers) and is never nil - it always exists, even in noop mode.
