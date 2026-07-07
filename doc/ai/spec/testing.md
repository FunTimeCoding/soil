# Testing Spec

Integration testing patterns for service tools. Unit testing philosophy lives in `conventions.md`.

## Store Testing

Use a real bbolt database on a temporary path:

```go
path := filepath.Join(t.TempDir(), constant.TestDatabase)
defer system.Remove(path)
s := store.New(path)
defer s.Close()
```

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
    lifecycle.WithWorker(p),
    lifecycle.WithServer(
        address,
        func(m *http.ServeMux) {
            m.HandleFunc("/api/alerts", route.Alerts(s))
        },
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
