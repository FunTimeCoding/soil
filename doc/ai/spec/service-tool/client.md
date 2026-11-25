---
base: pkg/tool/goatlassiand
---

# External API Client

When a daemon proxies a third-party service (Habitica, Mattermost,
Jellyfin, Sentry, etc.), the client library for that service
lives at `pkg/<domain>/`. This is the external API client - it talks
directly to the upstream service, not to our daemon.

```
pkg/<domain>/
├── client.go                       # Client struct (raw HTTP or SDK wrapper)
├── new.go                          # New(host, token string) *Client
├── new_environment.go              # NewEnvironment() - reads env vars
├── <operation>.go                  # Non-Must: returns (value, error)
├── must_<operation>.go             # Must: panics on error
└── constant/
    └── constant.go                 # HostEnvironment, TokenEnvironment
```

The daemon builds this client through its factory (`daemon.md`); no
surface imports it. **`pkg/<domain>/` is client-only.** All server-side concerns
(store, domain types, `model_context/`) belong in `pkg/tool/go<tool>d/`, not in
the shared library.

## Must and Non-Must Pattern

External API client methods come in pairs:

- `<operation>.go` - returns `(value, error)`. Used by daemon handlers
  that need to propagate errors gracefully (MCP and REST answer them by
  class through `answer`).
- `must_<operation>.go` - panics on error. Used by CLI tools and
  background workers running under lifecycle recovery.

```go
// links.go
func (c *Client) Links() ([]*link.Link, error) {
    // ... returns error on failure
}

// must_links.go
func (c *Client) MustLinks() []*link.Link {
    result, e := c.Links()
    errors.PanicOnError(e)
    return result
}
```

A Must twin exists only where a caller panics. Its callers are counted
by receiver, not by name - a `MustLinks` on another type is not a
caller - and a twin with none is deleted before the client's full face
is written, so the face carries no operation nobody calls. Every
consumer of an upstream client moves onto its full face and mock in the
change that writes them. A client that owns nothing - a connection per
call - has no `Close`, and neither does its face.

## Find Pattern

Lookups where absence is a valid outcome return the Find shape -
`Find<Noun>(...) (*T, bool, error)`. The flag carries absence,
the error carries failure; never a `nil, nil` soft-miss that
callers must remember to check. When absence is exceptional for
every caller, return a typed `not_found` instead of the flag.
`<Noun>Exists` variants wrap a lookup and swallow only
`not_found`. List endpoints are neither: an empty slice is a
complete answer, not a miss. Examples: `FindBranchRequest`,
`FindSession`, `FindLatestPulse`, `FindDocument`, `FileExists`.

Helpers that compose other calls (`ProjectsWithFile`, `Jobs`,
`RunnersByTag`) follow the same pair: the plain form composes the
plain variants and returns the first error, and a Must twin panics
on it. A method body never calls a Must variant - the panic belongs
to the caller's choice of variant, not to the helper, and a plain
signature that panics underneath breaks the MCP and REST callers
the pair exists for.

`pkg/<domain>/` talks to the third-party service; `<toold>/client/` talks
to our own daemon (`generated-api.md`, The Daemon Client).
