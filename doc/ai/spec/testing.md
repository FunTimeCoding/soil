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
- **A new test is proven by breaking what it guards** - break the code
  path, not its inputs, keep it compiling, watch the test fail, restore.
  `go test -overlay` breaks without touching the tree. One break at a
  time: a break that panics aborts the package and hides every other
  verdict. A break the test survives names a branch no test reaches, or
  code nothing needs. A break run reads its silences: `no tests to run` is a
  pattern that matched nothing, and a break whose pattern no longer matches
  broke nothing - both read green.
- **A filter test keeps a row and drops a row**, and counts the rows
  directly - a filter dropping everything passes a test asserting only
  the drop, and a read that deletes what it filters hides a missing
  cleanup.

## Test names

A test name is chopped: the words that carry the meaning, no articles
(`test_article` flags them), and no tail the verdict already tells apart -
`TestOpenAIServerIsReachable` beside
`TestOpenAIServerThatRefusesIsNotReachable`. Two tests of one behaviour at
different layers carry the layer. Prepositions stay when they carry meaning;
`preposition_name` doesn't read test names.

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

The store borrows the mapper, so the test closes it, as the daemon does
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

A mock reaches a service the way the real client does: the daemon builds
the client from the factory on the option and hands it to the services that
use it (`service-tool.md`, The daemon). Tests come in two kinds.

A service test stands in for the daemon: `<path>/unit/<noun>_tester` builds
the store over a mapper, the mock, and the service from those parts with the
services it was built with, closes what it built on cleanup, and returns
the service and the mock:

```go
func New(t *testing.T) (*loan.Service, *mock_client.Client) {
    t.Helper()
    m := lite.NewMemory()
    t.Cleanup(func() { relational.CloseMapper(m) })
    migrate.AutoMigrate(m)
    c := mock_client.New()
    t.Cleanup(c.Close)
    s := catalog.New(m)

    return loan.New(s, member.New(member_store.New(m)), c, time.Now), c
}
```

A surface test builds the daemon from the option - the factory returning the
mock, `LitePath` a file in `t.TempDir()` - closes it on cleanup, and goes served
through `Mount`. The daemon closes the mock it was handed through the factory,
so what the mock and the store carry before the daemon starts goes in as seeded
values, and a served test reading the store beside the daemon opens its own
connection on the option's `LitePath`. A pooled client's factory answers per
instance, so a pooled tester takes the instance names and returns one mock per
name. Nothing builds a surface over hand-wired services. A guarded service's
test shares one reporter between the guard and the caller, so the events it
counts are the ones the daemon would send.

A mock answers as the client does: never `nil, nil` where the client returns a
value, an unseeded query with an empty result, absence as the daemon answers
it. A double it replaces keeps its old default by seeding it. A mock that only
ever stands in for the client hides the client's own bugs, so each success path
also goes once through the real client against the served daemon. A double with
a `Close` is the test's to close: one helper builds it and closes it on cleanup.
A client too costly to build per test is shared behind a lease whose `Close`
does nothing, closed in `TestMain` after `m.Run()`. A fleet's fixtures become
the face's seeded data through one translator, proven by loading every fixture
back.

A fake upstream - an HTTP server answering the client's real requests with
their content type - stays for testing a client's own wire handling, in the
client's package; a browser-scraping client is tested through a page serving a
captured fixture's HTML. That package's test also holds the client and its mock
to the face, both in one `[]face.Client`. A daemon client's coverage test
reflects on `face.Client` and checks it against the embedded spec (`GetSpec`),
failing on any operation the face lacks; a method name differs from its
operation only by an explicit mapping in the test (gonetboxd's
`GetAnsibleInventory` as `AnsibleInventoryMarkup`). The client and mock follow
by compile through the face test.

## Poll Cycle Testing

Test poll logic directly by calling `Poll()` - no goroutine or ticker needed:

```go
w := worker.New(v.Poll, l, r, time.Minute)
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

A not-found answer through the daemon client cannot tell a missing record
from a missing route - the route's own 404 reads as `not_found` too - so the
double records what it was asked and the test asserts the ask.

## Peer Guard Tests

A guarded daemon has an outage table (`go:pkg/web/outage_tester`): every
REST operation and MCP tool against a peer that is down, each expecting
`unavailable` and no capture. Each operation meets a fresh daemon built
by the caller's `serve` on one shared reporter, so no operation inherits a
guard another one opened, and an operation swallowing its peer's error
shows by name.

A quiet class reads the same guarded or not, so an outage test cannot prove
the guard; what only the guard does can - the next call failing at once,
without asking the peer. The double fails as the client does: a dial failure
is a `connection.Failure` (a hand-built `net.OpError` classifies only around
a real errno, `os.NewSyscallError("connect", syscall.ECONNREFUSED)`), where a
bare `unreachable` error passes through as a refusal and the guard never
opens. The guard answers `unavailable` on the call that opens the outage and
reports it only after `Lasting`, so a fast test counts the peer's calls, not
the reporter's events. A guard comes back only through its probe loop, on a
real timer from `ProbeMinimum`, so a test that needs the peer back fails a
call with a fault the guard passes through untouched.

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
`generative/model_context_server.New(t, setup)` - `NewSignedOn` for a daemon
that signs on - and asserts the full
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

