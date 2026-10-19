# REST strict server error handling

Recovery middleware remains as a safety net for unexpected panics
even on strict server services - it catches anything that slips
past explicit error handling.

Handler bodies split by service class: external API services call
non-Must client methods and answer failures explicitly; local store
services may keep `Must*`/`PanicOnError` and lean on the recovery net.
The Must/non-Must pair convention lives in `service-tool.md`.

Services using oapi-codegen's strict server mode return typed error
responses. Which status, which schema and whether Sentry hears of it
are the error's class, in the table in `README.md` (Answering an
error); this leaf is how a handler builds each answer.

## Schemas

The OpenAPI spec defines two error schemas (`generated-api.md`, Error
schemas):

- `Error` - just an `error` string, for an answer Sentry does not hear of
- `ErrorResponse` - `error` + `event_identifier`, for a captured one

## The 4xx answers

oapi-codegen handles parameter binding itself - the strict handler
wrapper returns 400 before the handler runs if parameters fail binding.
For every quiet class an operation can meet, the spec declares its
status on the `Error` schema and the handler branches on the class
before anything else, writing the body out:

```go
if not_found.Is(e) {
    return server.DeleteLink404JSONResponse{Error: e.Error()}, nil
}
```

An operation that names an existing resource - in its path or its
body - declares a 404; one whose upstream rejects a caller's value
declares a 400. Listing, searching and creating have nothing of the
caller's to miss, so an upstream 404 there stays a captured 500.

## Captured answers

Everything else answers 500 - or the status its class names, such as
503 for a transport failure no guard stands in front of - through
`answer` and `answerWith`, private methods on the Server struct that
fill the generated `ErrorResponse` from `classify.Capture`:

```go
// answer_with.go
func (s *Server) answerWith(
    e error,
    message string,
) *server.ErrorResponse {
    text, event := classify.Capture(s.reporter, e, message)

    return &server.ErrorResponse{Error: text, EventIdentifier: event}
}

// answer.go
func (s *Server) answer(e error) *server.ErrorResponse {
    return s.answerWith(e, "")
}
```

```go
records, e := s.store.ByName(r.Params.Name)

if e != nil {
    return server.GetAlerts500JSONResponse(*s.answer(e)), nil
}
```

`classify.Capture` always captures: reaching the helper on REST means a
fault. A quiet class arriving there is an operation missing its branch,
captured under the title `missing <class> branch: <error>` so the issue
list says what to add. The body text follows the message rule in
`README.md`.

Daemons not yet moved still carry `captureFail`, which builds the same
`ErrorResponse` with the message it is given, and a `clientError` helper
for the 4xx body.

## Process failures with context

For external process failures with rich context (stdout, stderr),
construct the response inline with `CaptureWithContext`:

```go
return server.PostGenerate500JSONResponse{
    Error: "parser failed",
    EventIdentifier: s.reporter.CaptureWithContext(
        parser.Error,
        "parser",
        map[string]any{
            "output": parser.OutputString,
            "stderr": parser.ErrorString,
        },
    ),
}, nil
```
