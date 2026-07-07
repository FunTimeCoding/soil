# REST strict server error handling

Recovery middleware remains as a safety net for unexpected panics
even on strict server services - it catches anything that slips
past explicit error handling.

Services using oapi-codegen's strict server mode return typed error
responses. The OpenAPI spec defines two error schemas:

- `Error` - just an `error` string
- `ErrorResponse` - `error` + `event_identifier`

Three tiers:

**Tier 1 - Local validation** (400, `Error`, no Sentry):
oapi-codegen handles parameter binding validation automatically -
the strict handler wrapper returns 400 before the handler runs if
parameters fail binding. For application-level validation (e.g.
instance resolution in multi-instance services), return the 400
response type with `clientError`. No Sentry capture - these are
caller mistakes, not infrastructure failures.

```go
instance, e := s.resolveInstance(r.Params.Instance)

if e != nil {
    return server.ListNodes400JSONResponse(*clientError(e)), nil
}
```

`clientError` returns `*server.Error` (pointer - lint
requires pointer returns for structs).

**Tier 2 - Upstream parseable** (500, `ErrorResponse` with
extracted message, Sentry): the upstream API returned a structured
error we can parse. The extracted message goes in `ErrorResponse`
for diagnostic value, but the status code is 500 - we haven't
classified which upstream errors are truly client mistakes. See
[external-api.md](external-api.md) for the `captureDetail` pattern.

**Tier 3 - Upstream opaque** (500, `ErrorResponse` with
`constant.UnexpectedError`, Sentry): we can't extract anything
meaningful. The caller gets "unexpected error" plus a Sentry
event ID for investigation.

For local store services without external APIs, all errors are
tier 3 via `captureFail`:

```go
records, e := s.store.ByName(r.Params.Name)

if e != nil {
    return server.GetAlerts500JSONResponse(
        *s.captureFail(e, constant.UnexpectedError),
    ), nil
}
```

`captureFail` on REST servers returns `*server.ErrorResponse`.

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
