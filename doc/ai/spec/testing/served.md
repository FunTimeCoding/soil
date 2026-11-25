# Served Surfaces

A test that needs the daemon served starts the production `Mount` through
`generative/model_context_server.New(t, setup)` on a dynamic port, and talks
to it through the daemon's generated REST client on `s.Client()`, which
parses typed responses - no hand-rolled JSON getters. A raw lifecycle
server (`system.FindUnusedPort`, `assert.Listen`) is for testing the
lifecycle itself.

A not-found answer through the daemon client cannot tell a missing record
from a missing route - the route's own 404 reads as `not_found` too - so the
double records what it was asked and the test asserts the ask.

## HTTP Clients

The production rule is `http-client.md`; this is the test side. A test talks
HTTP through a client it owns, never `http.DefaultClient`,
`http.DefaultTransport` or `http.Get` and its siblings. The owner closes the
client's idle connections before the server stops: a connection Go's
transport dialled and parked unused counts as new to the server, and
`Shutdown` waits five seconds before treating a new connection as idle. The
tell is a test pinned at exactly 5.0s, or 10s for two, in
`gotestsum tool slowest`: something it opened - a client session, a stream -
was still held when the server stopped. A change to a shared client's
lifecycle runs every consumer's suite under `-race`: the goroutine that races
it lives in the consumer.

- On `generative/model_context_server`: its clients share the server's
  transport, so its `Close` closes them all. Helpers outside the package take
  the server, not a port or a URL.
- On a raw lifecycle server: a `stall_client.New().Build()` pool, its `Close`
  deferred after the lifecycle's `Close`, so it runs first.
- On an `httptest.Server`: its own `Client()`, which `Close` takes care of.

## Guard Battery

Every guarded daemon has `<path>/unit/guard_test.go` (package
`unit`): it starts the production `Mount` on a dynamic port via
`generative/model_context_server.New(t, setup)` - with `WithSignOn()` for a
daemon that signs on - and asserts the full
auth contract with the battery methods, mirroring the mount surface:

- `VerifyBase` — health and version open, and `route:/api/reachable` 401
  bare and 204 with the test token: the token guard and the peer probe
  route every daemon serves (`error-handling/infrastructure.md`, The probe
  route)
- `VerifyGuarded(path)` / `VerifyOpen(path)` / `VerifyOpenPost(path)`
  — one per guarded route worth naming and per open mount, including
  the dashboard root and live path of web-carrying daemons
- `VerifySession(path)` — a session page redirects a bare request to sign-on,
  reported by the server's `Halting()` client without following, and serves
  the test token
- `VerifyModelContext` — 401 bare on `route:/mcp` and `route:/sse`, handshake
  with the test token
- `VerifyStatus(path, status)` — exact status for a bare request,
  where route existence needs pinning (the other verbs accept any
  non-401, so a dead route passes them)

The battery never invokes tool or REST handlers, but every dependency still
enters as something callable: external API clients as their full mocks through
the factories on the option the service is built from, daemon-internal
dependencies as empty constructions (`inventory.New()`, in-memory stores). Nil
is never passed - mocks exist so nothing downstream needs a nil check. The
`nil_dependency` analyzer holds this for `face` interfaces passed to a `New…`
constructor; a genuinely optional dependency is set by a `With…` method. Daemons
with a test base run the guard test through it, over the real service with the
mock clients inside. Every web surface asserts its favicon with `VerifyFavicon`,
a 200 carrying `image/png` - `VerifyOpen` only shows a route needs no sign-on,
and a missing route passes it. Session (SSO) surfaces assert their dashboard
with `VerifySession`; a page whose handler reaches an external client the
battery cannot serve, or a live stream that holds its headers, is not worth the
wait. Tests pass `constant.DefaultVersion` where a mount takes a version.
