# Testing Spec

Integration testing patterns for service tools. Unit testing philosophy lives in `conventions.md`; where test files live is `test-placement.md`.

## Store Testing

Gorm stores test against in-memory sqlite by default - fast, no
disk, production-matching foreign key enforcement:

```go
s := store.New(lite.NewMemory())
```

Raw `database/sql` stores use `store.New(connection.NewMemory())`
the same way. Use the file openers with `t.TempDir()` only when the
test genuinely needs a file on disk. Never call
`gorm.Open` directly - the `restricted_call` analyzer confines it
to `pkg/relational/` (see `database.md`).

## Mock Clients

Mock external dependencies using hand-rolled structs (see Interfaces section in `conventions.md`). Mocks expose methods to manipulate state during a test:

```go
c := mock_client.New()
c.Add(&alert.Alert{Fingerprint: "fp1", Name: "HighMemory"})
// ... exercise code ...
c.Remove("fp1")
// ... exercise code again, assert resolution ...
```

## Poll Cycle Testing

Test poll logic directly by calling `Poll()` - no goroutine or ticker needed:

```go
w := worker.New(c, s, 1*time.Minute)
w.Poll()
records := s.ByName("HighMemory")
assert.Count(t, 1, records)
```

This pattern tests the full save/resolve lifecycle: add alerts to mock, poll, assert store state, remove from mock, poll again, assert resolution.

## Lifecycle HTTP Testing

Spin up a real lifecycle with HTTP server on a dynamic port:

```go
port := system.FindUnusedPort(19400)
address := fmt.Sprintf(":%d", port)
l := lifecycle.New(
    logger.New(context.Background()),
    lifecycle.WithWorker(p),
    lifecycle.WithServer(
        server.New(
            constant.Identity,
            address,
            func(m *http.ServeMux) {
                m.HandleFunc("/api/alerts", route.Alerts(s))
            },
        ),
    ),
)
l.Run()
defer l.Stop()
assert.Listen(t, port)
```

Key helpers:
- `system.FindUnusedPort(startPort)` - finds an available port
- `assert.Listen(t, port)` - waits up to 2s for port to accept connections
- `assert.NotListen(t, port)` - asserts port is closed
- `assert.HTTPStatus(t, url, expectedStatus)` - GET request + status assertion

## Guard Battery

Every guarded daemon has `integration/guard/guard_test.go` (package
`guard`): it starts the production `Mount` on a dynamic port via
`generative/model_context_server.New(t, setup)` and asserts the full
auth contract with the battery methods, mirroring the mount surface:

- `VerifyBase` — health and version open
- `VerifyInterface` — the `/api/` tree rejects bare requests (probes
  `/api/guard-probe`; never spell the probe path at a call site)
- `VerifyGuarded(path)` / `VerifyOpen(path)` / `VerifyOpenPost(path)`
  — one per guarded route worth naming and per open mount, including
  the dashboard root and live path of web-carrying daemons
- `VerifyModelContext` — 401 bare on `/mcp` and `/sse`, handshake
  with the test token
- `VerifyStatus(path, status)` — exact status for a bare request,
  where route existence needs pinning (the other verbs accept any
  non-401, so a dead route passes them)

The battery never invokes tool or REST handlers, so domain
dependencies enter as typed nils or empty constructions
(`inventory.New()`, in-memory stores); clients that validate their
environment at construction always enter as typed nils. Daemons with
an integration base run the guard test through the base, and bases
run the full production `Mount` — mock clients flow through it
because `Mount`, the REST server, and the model_context package all
consume the daemon's `face/` interfaces, never the concrete clients.
Session (SSO) web surfaces assert their favicon instead of the
dashboard — the sign-in redirect points at a fake gate the test
client cannot follow. Tests pass `constant.DefaultVersion` where a
mount takes a version.

## Typed Response Parsing

Use a generic helper to parse JSON responses in tests:

```go
func getJSON[T any](
    t *testing.T,
    locator string,
) T {
    t.Helper()
    r, e := http.Get(locator)

    if e != nil {
        t.Fatal(e)
    }

    defer errors.PanicClose(r.Body)
    body, e := io.ReadAll(r.Body)
    errors.PanicOnError(e)

    var result T
    errors.PanicOnError(json.Unmarshal(body, &result))

    return result
}
```

Usage:

```go
status := getJSON[route.StatusResponse](t, base+"/api/status")
assert.Integer(t, 2, status.TotalRecords)
```

## Mid-Test State Manipulation

Integration tests can manipulate mock state between assertions to verify behavior changes:

```go
// Initial state: alert firing
p.Poll()
alerts := getJSON[[]route.AlertsResponse](t, base+"/api/alerts?name=X")
assert.String(t, route.Firing, alerts[0].Status)

// Remove alert, poll again: now resolved
c.Remove("fp1")
p.Poll()
alerts = getJSON[[]route.AlertsResponse](t, base+"/api/alerts?name=X")
assert.String(t, route.Resolved, alerts[0].Status)
```
