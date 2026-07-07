# Naming Conventions

Enforced by the `goanalyze` naming analyzer (`pkg/lint/analyzer/naming/`). Auto-fixed by `gofix`.

The full banned-segment list lives in `pkg/lint/segment/` - `suggestions.go`
(~50 abbreviations with letter/word replacements: `id`, `url`, `mcp`, `dir`,
`src`, `ctx`, `msg`, `req`, `json`, `cfg`, `err`, `val`, `args`, `opt`, ...)
and `no_suggestion.go` (banned with no auto-suggestion: `handler`, `data`,
`info`). That code is the source of truth; don't duplicate it here.

Banned segments cause a lint error; this spec documents the resolution
pattern for the segments where the fix isn't a mechanical replacement.

## Banned segments and replacements

| Banned | Context | Use instead |
|--------|---------|-------------|
| `data` | local `[]byte` from `os.ReadFile`, `json.Marshal*` | `b` |
| `data` | local `os.FileInfo`/`fs.FileInfo` | `i` |
| `data` | struct field holding opaque `map[string]any` | `Payload` |
| `data` | struct field holding HTTP request body (`any`) | `Body` |
| `data` | struct field holding text/string content | `Content` |
| `data` | struct field holding Kubernetes Secret `.data` block | `Payload` |
| `data` | type name suffix (e.g. `FooData`) | strip suffix when no conflict; `FooPayload` when `Foo` is taken in the same package |
| `data` | local byte slice parameters (exported funcs) | domain-specific: `payload`, `archive`, etc. |
| `info` | local `os.FileInfo`/`fs.FileInfo` in walk/loop | `i` |
| `info` | type name suffix (e.g. `FooInfo`) | strip suffix or replace with descriptive noun (`Meta`) |
| `handler` | struct type implementing routes | `Server` (see generated-api.md) |
| `handler` | function parameter | `serve` |
| `handler` | function name suffix | drop suffix: `logRequestHandler` → `logRequest` |
| `config` | type name | `Configuration` |
| `config` | package-level data struct | extract to `<name>_option/` package as `Option` |
| `config` | function name | `LoadConfiguration`, etc. |
| `llm` | any identifier | `model` (e.g. `llmPrefix` → `modelPrefix`) |
| `tmp` | field prefix | drop prefix: `TmpFromAccountID` → `FromAccountID` |
| `doc` | local `*goquery.Document` | `d` |

## Type name conflicts

When stripping a banned suffix would produce a name already taken in the same package,
append `Payload` instead of stripping:

```go
// Comment type exists → can't strip to Comment
type CommentPayload struct { ... }

// No conflict → strip cleanly
type Exception struct { ... }  // was ExceptionData
```

## Metadata types

Avoid acronym-based type names. Prefer a plain English noun:

```go
type Meta struct { ... }  // was SDKInfo - holds SDK name/version/packages
```

`Meta` is appropriate when the struct holds metadata about a tool or runtime component
and no collision exists in the package.

## Domain terms

Some names contain a banned segment but refer to a real-world domain concept. Rename to the
most specific English term for that domain artifact:

| Original | Renamed | Rationale |
|----------|---------|-----------|
| `DataFile = "data.tar.gz"` | `ArchiveFile` | Alpine package archive segment |
| `PackageInfoFile = ".PKGINFO"` | `MetadataFile` | Alpine package metadata file |
| `CreateDataTar()` | `CreateArchive()` | creates the archive segment |
| `DataDirectory` (packager field) | `ArchiveDirectory` | directory whose contents become the archive |
| `dataDirectory` / `DataDirectory()` | `storageDirectory` / `StorageDirectory()` | app-level storage path |
| `DataDirectoryEnvironment` | `StorageDirectoryEnvironment` | env var pointing to storage dir |
| `DataHost` (CDN constant) | `AssetHost` | CDN endpoint for media files |

## HTTP handler conventions

See `generated-api.md` for the full Server pattern. Summary:

- The struct implementing `ServerInterface` is named `Server`, never `Handler` or `Router`
- Receiver on `*Server` methods: `s`
- `*http.Request` parameter: `q` (avoids collision with receiver)
- Function parameter that accepts a handler func: `serve`, not `handler`
- Handler methods/functions are named after the resource they serve, without a `handle` prefix: `handleAlerts` → `alerts`, `handleDashboard` → `dashboard`, `handleAddSubmit` → `addSubmit`

## HTML node builder naming

Functions returning `g.Node` (gomponents HTML builders) are named after the component they produce, not the action of building it:

- `alertsTable()` - renders the alerts table
- `addForm()` - renders the add entry form
- `navigationLink()` - renders a navigation link

File name follows the function name: `alerts_table.go`, `add_form.go`, `nav_link.go`.

## Variable shorthands

These are established conventions across the codebase:

| Type | Name |
|------|------|
| `[]byte` from `ReadFile` / `Marshal*` | `b` |
| `os.FileInfo` / `fs.FileInfo` in walk/loop | `i` |
| `*goquery.Document` | `d` |
| `*gzip.Writer` | `z` |
| `*tar.Writer` | `t` |
| `error` | see `conventions.md` Error Handling for full `e`/`f`/`g` escalation |
| serialized markup (`[]byte` of yaml/xml/html) | `m` |
| serialized notation (`[]byte` of json) | `j` |

See `conventions.md` for the full short variable name table.

## Single-character collision fallback

When the suggested single character is already taken in scope (e.g. `m` is the receiver of `*Model`), step through the letters of the full word in order until a free one is found:

- `markup` → `m`, `a`, `r`, `k`, `u`, `p`
- `notation` → `j` (first letter of json, the dominant format), then `n`, `o`, `t`, `a`, `i`
- `error` → `e`, `f`, `g`, `h`

The linter suggests only the first option. Choosing a later letter when there is a collision is left to the developer.

## HTML custom attributes

`data-*` attributes are HTML's custom attribute mechanism. Use `Custom` as the replacement term:

```go
CustomPrefix       = "data-"   // was DataPrefix
KeepCustomAttribute bool        // was KeepDataAttribute
```
