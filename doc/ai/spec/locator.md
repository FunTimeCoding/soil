# Locator

`pkg/web/locator` is a fluent URL builder used throughout the codebase. Use it whenever constructing a URL - never concatenate strings manually.

## Core Behavior

- `locator.New(host)` defaults to **HTTPS** - no need to specify the scheme for normal use
- All builder methods return `*Locator` for chaining
- `String()` produces the final URL string
- `Copy()` deep-copies a locator - use this when storing a base locator in a client and building per-request URLs from it

## Common Patterns

### Base URL for a client

Store the locator on the client, copy it per request:

```go
type Client struct {
    base *locator.Locator
}

func New(host string) *Client {
    return &Client{base: locator.New(host)}
}

func (c *Client) Get(path string, params map[string]string) {
    l := c.base.Copy().Path(path)
    for k, v := range params {
        l.Set(k, v)
    }
    // use l.String()
}
```

### Ad-hoc URL construction

```go
locator.New(host).Path("/api/alerts").Set("name", name).String()
```

### Non-standard port

```go
locator.New(host).Port(8080).String()
// → https://host:8080
```

### HTTP (insecure)

```go
locator.New(host).Insecure().String()
// → http://host
```

### Custom scheme (WebSocket, etc.)

```go
locator.New(host).Port(port).Scheme("wss").String()
```

### Path with fmt-style formatting

```go
locator.New(host).Path("/%s/-/merge_requests", project).String()
```

### Reusable base path

```go
locator.New(host).Port(port).Base("/api").Path("/alerts").String()
// → https://host:port/api/alerts
```

### Basic auth

```go
locator.New(host).UserPassword(user, password).String()
// → https://user:password@host
```

### Passing to oapi-codegen generated client

Generated clients expect a full URL with scheme. Use `locator.New(host).String()`:

```go
c, err := client.NewClient(locator.New(host).String())
```

## Method Reference

| Method | Purpose |
|--------|---------|
| `New(host)` | Create locator defaulting to HTTPS |
| `Copy()` | Deep copy - use when building per-request URLs from a stored base |
| `String()` | Produce final URL string |
| `Pointer()` | Produce `*string` of the URL |
| `Path(f, ...a)` | Set request path (fmt.Sprintf style) |
| `Base(f, ...a)` | Set reusable base path prefix |
| `Port(p)` | Set port (omitted from URL when 0) |
| `Scheme(s)` | Override scheme |
| `Insecure()` | Switch to HTTP |
| `Set(k, v)` | Set query parameter (replaces) |
| `Add(k, v)` | Add query parameter (allows multiple values) |
| `SetInteger(k, v)` | Set int query parameter |
| `SetInteger64(k, v)` | Set int64 query parameter |
| `UserPassword(u, p)` | Set basic auth credentials |
| `Trail()` | Add trailing slash |

## What Not To Do

- Don't concatenate `"https://" + host` - use `locator.New(host).String()`
- Don't mutate a stored base locator - always `Copy()` it first
- Don't pass a bare hostname to a generated HTTP client - wrap with `locator.New(host).String()`
