# HTTP Client

Convention for outbound HTTP: every client is built by `stall_client`,
owned by exactly one holder, and closed by it. No process-wide transport,
no `http.DefaultClient`.

## Building

`stall_client.New()` returns a builder; `Build()` returns a
`stall_client/pool.Pool`, which owns one `*http.Transport` and hands out
views over it:

```go
p := stall_client.New().Insecure().Long().Timeout(5 * time.Second).Build()
defer p.Close()
c := p.Client()
```

- `Insecure()` - skip TLS verification
- `Long()` - wait thirty minutes for response headers instead of thirty
  seconds; bodies are never bounded
- `Timeout(d)` - an overall client timeout, on every view
- `Verbose()` - log each request and response
- `BasicAuth(user, password)` - add basic authentication
- `BearerAuth(token)` - add a bearer token
- `Ping(send, wait)` - the HTTP/2 health check below, with other durations
- `Observe(o)` - record every request's duration, labelled `host`, `code`
  and `method`, into `metric.Upstream(identity)`, the daemon's slice of the
  shared `upstream_request_duration_seconds`

`code` is the response status, or, when no response arrived, the
`connection.Classify` kind - `timeout`, `unreachable`, `dropped` - then
`canceled` when the caller gave up and `error` for the rest. Prometheus's own
round-tripper instrumentation records none of these. `host` comes from each
request, always with a port - the scheme's default when the locator leaves it
out, so a public and an in-cluster address both read `host:port`. Observe only
clients with a fixed set of hosts: a client that
follows arbitrary links would grow a series per site.

`Verbose` always wraps outside `BasicAuth` and `BearerAuth`, so the log never
carries the credentials.

Every transport checks its HTTP/2 connections: a connection that has received
nothing for thirty seconds gets a ping, and one whose ping goes unanswered for
fifteen is closed. Without it, a connection whose path went silent stays in the
pool, `ResponseHeaderTimeout` cancels only the request, and every later request
is multiplexed onto the dead connection and times out in turn
(`TestStallClientDropsSilentConnection`).

The transport is the resource: it holds the connections. A client is a
policy view over it - `Client()` follows redirects, `Halting()` stops at
the first one and returns it. An SDK that builds its own wrapper and takes a
base round-tripper gets `p.Client().Transport`, so the switches still apply;
go-jira's auth transport and gowebdav take it that way, and oauth2 takes
`p.Client()` through its context. `Transport()` is the bare transport
underneath.

## Owning

Whoever builds a pool owns it and closes it; whoever is given a client
never closes it.

- A type that builds a client keeps the pool, or the type wrapping it, and
  has a `Close()` that closes everything it built. A wrapper around a
  generated client builds its pool in `New` and passes `p.Client()` through
  `WithHTTPClient`.
- A function that would hand out a bare generated or SDK client takes an
  `*http.Client` from its caller instead, so the caller owns it.
- A daemon's outermost owner is its service: it builds the clients from
  factories on the option, and `Run` builds it and defers its `Close()`
  (`service-tool.md`, Service). A CLI's `main` builds its clients itself
  and defers their `Close()`. A one-shot function owns a local pool for
  the length of its call.

Closing drops idle connections only; a client closed twice, or used again
after `Close`, simply dials again.

## Enforcement

goanalyze's `restricted_call` refuses `http.Get`, `Post`, `Head`,
`PostForm`, `DefaultClient` and `DefaultTransport`; `restricted_literal`
refuses `http.Client` and `http.Transport` literals outside `stall_client`.
`unclosed_resource` follows ownership: an owner held in a local that is
never closed, a holder whose `Close` never reaches a field it built, and a
value handed to a function that only borrows it - the fixtures in
`pkg/lint/integration/unclosed_resource/` show what it derives.

## Exceptions

- `generative/model_context_client` wraps the test server's transport in a
  bearer round-tripper.
- gokubernetesd's port-forward uses client-go's SPDY upgrade transport, and
  client-go builds its transport from `rest.Config`, cluster TLS and auth
  plugins included.
- A test that needs a transport the builder cannot express - a failing
  dialer - marks its literal `goanalyze:ignore restricted_literal`.

Test servers own their client too; `testing.md` covers that side.
