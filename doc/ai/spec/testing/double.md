# Test Doubles

Mock external dependencies using hand-rolled structs
(`conventions/interfaces.md`). Mocks expose methods to manipulate state
during a test:

```go
c := mock_client.New(t)
c.Add(&alert.Alert{Fingerprint: "fp1", Name: "HighMemory"})
// ... exercise code ...
c.Remove("fp1")
// ... exercise code again, assert resolution ...
```

A mock reaches a service the way the real client does: the daemon builds
the client from the factory on the option and hands it to the services that
use it (`service-tool/daemon.md`). Tests come in two kinds.

## Service tests

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
    c := mock_client.New(t)
    s := catalog.New(m)

    return loan.New(s, member.New(member_store.New(m)), c, time.Now), c
}
```

A guarded service's test shares one reporter between the guard and the caller,
so the events it counts are the ones the daemon would send.

## Surface tests

A surface test builds the daemon from the option - the factory returning the
mock, `LitePath` a file in `t.TempDir()` - closes it on cleanup, and goes served
through `Mount`. The daemon closes the mock it was handed through the factory,
so what the mock and the store carry before the daemon starts goes in as seeded
values, and a served test reading the store beside the daemon opens its own
connection on the option's `LitePath`. A pooled client's factory answers per
instance, so a pooled tester takes the instance names and returns one mock per
name. Nothing builds a surface over hand-wired services.

## How a double answers

- **As the client does** - never `nil, nil` where the client returns a value,
  an unseeded query with an empty result, absence as the daemon answers it. A
  double replacing another keeps the old default by seeding it.
- **Through the real client once** - a mock that only ever stands in hides the
  client's own bugs, so each success path also goes once through the real
  client against the served daemon.
- **One double, one collaborator** - a double handed to a peer and its guard
  at once makes every mis-wiring between them behave like the right one; only
  the served path proves that wiring.
- **Atomic under a goroutine** - when production gains a goroutine that reads
  a double, such as a guard's probe loop over a mock, the double's switches
  become `atomic`, or a test flipping one races.
- **Closed by the test** - one helper builds a double with a `Close` and closes
  it on cleanup. A client too costly to build per test is shared behind a
  lease whose `Close` does nothing, closed in `TestMain` after `m.Run()`.
- **Fixtures through one translator** - a fleet's fixtures become the face's
  seeded data through one translator, proven by loading every fixture back.

## Fake upstreams

A fake upstream - an HTTP server answering the client's real requests with
their content type - stays for testing a client's own wire handling, in the
client's package; a browser-scraping client is tested through a page serving a
captured fixture's HTML. That package's test also holds the client and its mock
to the face, both in one `[]face.Client`. A daemon client's coverage test
reflects on `face.Client` and checks it against the embedded spec (`GetSpec`),
failing on any operation the face lacks; a method name differs from its
operation only by an explicit mapping in the test. The client and mock follow
by compile through the face test.

A peer that only answers is `go:pkg/web/upstream_server` at the call site, and
the test reads what reached it through `Last()`, `LastBody()` and `Count()`. A
builder in the home's `<subject>_tester` takes the address and builds the
client against it. A peer that routes or keeps state is a `<peer>_server` of
its own (`test-placement.md`). The fake carries `WithObject()` only when its
client chooses a decoder by `Content-Type`: generated OpenAPI clients and the
openai SDK refuse a plain-text answer. A refusal fake whose client reads errors
through `refusal.Reason` needs none, since that parses any body.
