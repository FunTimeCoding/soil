# Testing Spec

Testing philosophy and integration patterns for service tools. Where
test files live is `test-placement.md`.

## Philosophy

- **Unix approach: silence on success, noise on failure.** No
  progress logs (`t.Log("doing X...")`), no success logs
  (`t.Logf("✓ X succeeded")`). Only `t.Fatalf()` for must-succeed
  setup, `t.Errorf()` for validation failures.
- **Assert package** - `pkg/assert` for every assertion, over raw `if`
  checks, argument order `(t, expected, actual)`. The `forbidden_import`
  analyzer refuses testify.
- **No defensive programming in tests** - "should work or blow up".
  Extract logic into helpers that panic instead of returning errors;
  helpers return only success values.

## Service First

Behaviour is tested where it lives: the service, called directly. A
surface - MCP tool, REST route, web handler - parses its input, calls the
service and maps the error, so a test through it repeats what the
service test already holds. A surface earns a test only for what it alone
decides: an error it classifies differently (a not-found answered rather
than captured), a parse no other surface shares, a form round-trip a
reader would not notice breaking. Logic two surfaces share moves into the
service, where one test covers both. The guard battery is the standing
exception - the auth contract belongs to the mount, not the service.

## Store Testing

Gorm stores test against in-memory sqlite by default:

```go
m := lite.NewMemory()
t.Cleanup(func() { relational.CloseMapper(m) })
migrate.AutoMigrate(m)
s := store.New(m)
```

The store borrows the mapper, so the test closes it, as the service does
in production (`database.md`). Raw `go:database/sql` stores use
`store.New(connection.NewMemory())` the same way. Use the file openers
with `t.TempDir()` only when the test genuinely needs a file on disk. Never call
`gorm.Open` directly - the `restricted_call` analyzer confines it
to `pkg/relational/` (see `database.md`).

## Mock Clients

Mock external dependencies using hand-rolled structs
(`conventions/interfaces.md`). Mocks expose methods to manipulate state
during a test:

```go
c := mock_client.New()
c.Add(&alert.Alert{Fingerprint: "fp1", Name: "HighMemory"})
// ... exercise code ...
c.Remove("fp1")
// ... exercise code again, assert resolution ...
```

A mock reaches the service the way the real client does: through the factory on
the option (`service-tool.md`, Service). `<path>/unit/service_tester` creates
the mock, builds the option with the factory returning it, builds the service
and closes it on cleanup, and returns both; a pooled client's factory answers
per instance, so a pooled tester takes the instance names and returns one mock
per name -
`New(t, names...) (*service.Service, map[string]*mock_client.Client)`. The
service closes the mock it was handed through the factory, so no caller holds
one it never closes - data the mock should carry before the service starts goes
in as values, not as a mock the caller built. A guarded service's test shares
one reporter between the guard and the caller, so the events it counts are the
ones the daemon would send. Tests of the service, its workers and the mount all
go through it:

```go
func New(t *testing.T) (*service.Service, *mock_client.Client) {
    t.Helper()
    c := mock_client.New()
    o := option.New()
    o.LitePath = filepath.Join(t.TempDir(), "daemon.db")
    o.Source = func() face.AlertSource { return c }
    result := service.New(o, logger.New(t.Context()))
    t.Cleanup(result.Close)

    return result, c
}
```

A service built from the option opens its database from `LitePath`, so
service-level tests run on a file in `t.TempDir()`; store tests keep the
in-memory mapper. A fake upstream - an HTTP server answering the
client's real requests with their content type - stays for testing a
client's own wire handling, in the client's package. That package's
test also holds the client and its mock to the face, both in one
`[]face.Client`, and a daemon client's to its embedded spec
(`GetSpec`), failing on any operation the client lacks.

## Poll Cycle Testing

Test poll logic directly by calling `Poll()` - no goroutine or ticker needed:

```go
w := worker.New(c, s, 1*time.Minute)
w.Poll()
records := s.ByName("HighMemory")
assert.Count(t, 1, records)
```

## Served Surface Testing

A test that needs the daemon served starts the production `Mount` through
`generative/model_context_server.New(t, setup)` on a dynamic port, and talks
to it through the daemon's generated REST client on `s.Client()`, which
parses typed responses - no hand-rolled JSON getters. A raw lifecycle
server (`system.FindUnusedPort`, `assert.Listen`) is for testing the
lifecycle itself.

## HTTP Clients

The production rule is `http-client.md`; this is the test side. A test talks
HTTP through a client it owns, never `http.DefaultClient`,
`http.DefaultTransport` or `http.Get` and its siblings. The owner closes the
client's idle connections before the server stops: a connection Go's
transport dialled and parked unused counts as new to the server, and
`Shutdown` waits five seconds before treating a new connection as idle.
A change to a shared client's lifecycle runs every consumer's suite under
`-race`: the goroutine that races it lives in the consumer.

- On `generative/model_context_server`: `Connect(t)` for the MCP client,
  `Client()` for raw requests and generated REST clients
  (`client.WithHTTPClient(s.Client())`), `Halting()` where a redirect is the
  answer, `Address(target)` for the URL - a target may carry a query. `Close`
  closes them all, because they share the server's transport. Helpers outside
  the package take the server, not a port or a URL.
- On a raw lifecycle server: a `stall_client.New().Build()` pool, its `Close`
  deferred after the lifecycle's `Close`, so it runs first.
- On an `httptest.Server`: its own `Client()`, which `Close` takes care of.

## Guard Battery

Every guarded daemon has `<path>/unit/guard_test.go` (package
`unit`): it starts the production `Mount` on a dynamic port via
`generative/model_context_server.New(t, setup)` and asserts the full
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

The battery never invokes tool or REST handlers, but every
dependency still enters as something callable: external API
clients as their full mocks through the factories on the option the
service is built from, daemon-internal dependencies as empty
constructions (`inventory.New()`, in-memory stores). Nil is never
passed - mocks exist so nothing downstream needs a nil check. The
`nil_dependency` analyzer holds this for `face` interfaces passed to a
`New…` constructor; a genuinely optional dependency is set by a `With…`
method. Daemons with a test base run the guard test through it, over the
real service with the mock clients inside. Session (SSO) web surfaces
assert their favicon open and their
dashboard with `VerifySession`; a page whose handler reaches an
external client the battery cannot serve, or a live stream that
holds its headers, is not worth the wait. Tests pass
`constant.DefaultVersion` where a mount takes a version.

