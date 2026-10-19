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

## Everything else: answer and answerWith

Every other error a handler meets - from a service, a store, a client,
even local file I/O - goes to one of two private methods on the Server
struct, each in its own file:

```go
// answer_with.go
func (s *Server) answerWith(
    e error,
    message string,
) (*mcp.CallToolResult, error) {
    return response.AnswerWith(s.reporter, e, message)
}

// answer.go
func (s *Server) answer(e error) (*mcp.CallToolResult, error) {
    return s.answerWith(e, "")
}
```

`response.AnswerWith` reads the class table. A quiet class fails the
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

`answerWith` is for a tool that does several things, where the context
says which one failed - adding, removing or replacing tags in one call.
A single-step tool answers `answer(e)`; restating the tool's own name is
not a context.

A step a daemon must run before every answer - gosentryd records the
reacher's outage episodes - sits in `answer` before it hands over, so a
site that needs the step never calls `answerWith` directly.

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
