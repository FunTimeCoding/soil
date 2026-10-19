# REST strict server error handling

Recovery middleware remains as a safety net for unexpected panics
even on strict server services - it catches anything that slips
past explicit error handling.

Handler bodies split by service class: external API services call
non-Must client methods and route failures through captureDetail;
local store services may keep `Must*`/`PanicOnError` and lean on
the recovery net. The Must/non-Must pair convention lives in
`service-tool.md`.

Services using oapi-codegen's strict server mode return typed error
responses. Which status, which schema and whether Sentry hears of it
are the error's class, in the table in `README.md` (Answering an
error); this leaf is how a handler builds each answer.

## Schemas

The OpenAPI spec defines two error schemas (`generated-api.md`, Error
schemas):

- `Error` - just an `error` string, for an answer Sentry does not hear of
- `ErrorResponse` - `error` + `event_identifier`, for a captured one

## Uncaptured answers

The status the class names, with `Error`. oapi-codegen handles parameter
binding itself - the strict handler wrapper returns 400 before the
handler runs if parameters fail binding. For a class the handler
answers, return the response type with `clientError`, or the message
inline:

```go
instance, e := s.resolveInstance(r.Params.Instance)

if e != nil {
    return server.ListNodes400JSONResponse(*clientError(e)), nil
}
```

`clientError` returns `*server.Error` (pointer - lint
requires pointer returns for structs).

## Captured answers

`captureFail` captures the error and returns `*server.ErrorResponse`
carrying the message it is given and the event identifier. A service
with only internal stores passes `constant.UnexpectedError`:

```go
records, e := s.store.ByName(r.Params.Name)

if e != nil {
    return server.GetAlerts500JSONResponse(
        *s.captureFail(e, constant.UnexpectedError),
    ), nil
}
```

A service over an external API extracts the upstream's own message
where it can, through `captureDetail` (see
[the captureDetail pattern](external-api.md#the-capturedetail-pattern)).
A transport failure with no guard in front answers 503 the same way,
relaying its own message:

```go
if unreachable.Is(e) {
    return server.GetTabs503JSONResponse(
        *s.captureFail(e, e.Error()),
    ), nil
}
```

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
