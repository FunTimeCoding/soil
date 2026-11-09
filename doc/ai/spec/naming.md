# Naming Conventions

Enforced by the `goanalyze` naming analyzer (`pkg/lint/analyzer/naming/`).
Auto-fixed by `gofix`.

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

When stripping a banned suffix would produce a name already taken in the same
package, append `Payload` instead of stripping:

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

`Meta` is appropriate when the struct holds metadata about a tool or runtime
component and no collision exists in the package.

## Domain terms

Some names contain a banned segment but refer to a real-world domain concept.
Rename to the most specific English term for that domain artifact:

| Original | Renamed | Rationale |
|----------|---------|-----------|
| `DataFile = "data.tar.gz"` | `ArchiveFile` | Alpine package archive segment |
| `PackageInfoFile = ".PKGINFO"` | `MetadataFile` | Alpine package metadata file |
| `CreateDataTar()` | `CreateArchive()` | creates the archive segment |
| `DataDirectory` (packager field) | `ArchiveDirectory` | directory whose contents become the archive |
| `dataDirectory` / `DataDirectory()` | `storageDirectory` / `StorageDirectory()` | app-level storage path |
| `DataDirectoryEnvironment` | `StorageDirectoryEnvironment` | env var pointing to storage dir |
| `DataHost` (CDN constant) | `AssetHost` | CDN endpoint for media files |

## Environment variable names

Singular, even when the value is a comma-separated list:
`CREDENTIAL_REVEALED_FIELD` carries `A,B,C`. The general
avoid-plurals rule applies to env var names too - a list-valued
variable does not earn a plural.

Transport and trust are separate axes and never share a variable:

- `*_INSECURE` selects http instead of https - a locator concern,
  fed to `locator.Environment`.
- `*_UNTRUSTED` skips certificate verification - a client concern,
  fed to a stall_client pool built with `Insecure()`.

An endpoint can need one, the other, or both, and a single flag
cannot express https-with-an-unverifiable-certificate. Where a
client offers these as options, the same split applies:
`WithInsecure` sets the scheme, `WithUntrusted` skips
verification.

## HTTP handler conventions

See `generated-api.md` for the full Server pattern. Summary:

- The struct implementing `ServerInterface` is named `Server`, never `Handler`
  or `Router`
- Receiver on `*Server` methods: `s`
- `*http.Request` parameter: `q` (avoids collision with receiver)
- Function parameter that accepts a handler func: `serve`, not `handler`
- Handler methods/functions are named after the resource they serve, without a
  `handle` prefix: `handleAlerts` → `alerts`, `handleDashboard` → `dashboard`,
  `handleAddSubmit` → `addSubmit`

## Qualifier order

Compound names order big-to-small: the domain noun first, the
narrowing qualifier after - `CertificateKindServer`,
`RegistryGitHub`, `ImagePrefixGitHub`. A name reads as the field
or concept it populates, then its variant. Small-to-big forms
(`GitHubPrefix` populating an `ImagePrefix` field) are the drift
to correct when touched.

## Prepositions

A preposition where a noun should be names the argument list instead
of the thing: `PriceOf(item)`, `UserByName(name)`, `valueAt(series,
second)`. The noun that should answer for itself became an argument,
and the preposition marks where its receiver went. The query moves
onto its noun (`item.Price()`), a lookup's parameter names its key
(`User(name)`), a role becomes a qualifier (`AuthoredPosts(author)`),
and a from-and-until pair becomes one noun, a window. An act is a verb,
and the noun it lands on owns it: `RemoveFromList` is `list.Remove`. Where
the owner is a qualifier, it moves forward: `ListItemsInAisle` is
`ListAisleItems`. A method that never reads its receiver says the same from
the other side: it is a function, or a method of what it reads.

Two analyzers report these, not yet registered: `preposition_name` over
every name the code defines - functions, methods, types, fields, constants,
variables, parameters and locals - and `unread_receiver`. The words and the
places they count are `pkg/lint/constant/preposition.go`: `with`, like
`from`, counts anywhere but first and never after `New`; `per` and the verb
particles (`up`, `out`, `off`, `down`) are not prepositions here. Outside the
first:

- interface implementations, and test functions (named as `testing.md`, Test
  names);
- yes/no states: a bool field, a bool answer taking no argument, and `In`
  opening a bool question (`InStock(name)`);
- moments: a field, parameter or local ending in `At` or holding a time; a
  constant or package variable naming a moment or its column, field or index
  (`ShippedAtColumn`); and a function answering a time, then at most a bool
  and an error, whose word is `at`;
- a parameter or local that is the bare word (`from`, `to`, `since`).

## Ordering variants

Query and store methods that differ only in ordering are separate
named methods, never a boolean direction parameter - Go has no
named arguments, so `Seeds(since, 10, true)` tells the call-site
reader nothing. The name carries the semantics, and the moment goes in the
parameter, never in the name:

- `Versions(since)` - forward walk from a point in time, oldest first
  (streaming consumers: pollers, watermarks)
- `RecentVersions(...)`, `RecentSeeds()` - `Recent` alone marks
  newest-first (display consumers: dashboards, tables)

## HTML node builder naming

Functions returning `g.Node` (gomponents HTML builders) are named after the
component they produce, not the action of building it:

- `alertsTable()` - renders the alerts table
- `addForm()` - renders the add entry form
- `navigationLink()` - renders a navigation link

File name follows the function name: `alerts_table.go`, `add_form.go`,
`nav_link.go`.

## Variable shorthands

These are established conventions across the codebase:

| Type | Name |
|------|------|
| `[]byte` from `ReadFile` / `Marshal*` | `b` |
| `os.FileInfo` / `fs.FileInfo` in walk/loop | `i` |
| `*goquery.Document` | `d` |
| `*gzip.Writer` | `z` |
| `*tar.Writer` | `t` |
| `error` | `e`, then `f`, `g` in the same scope (avoid shadowing) |
| serialized markup (`[]byte` of yaml/xml/html) | `m` |
| serialized notation (`[]byte` of json) | `j` |
| headers | `h` |
| `os.FileInfo` in loops | `i` |
| files | `f` |

When e/f/g run out in deeply nested scopes, short `<verb>Fail` names
(`createFail`, `saveFail`) - prefer short words like "fail" over
"error". A descriptive error name is still a smell: the function
wants refactoring. Eliminate the variable entirely when possible:
`errors.PanicOnError(someCall())`.

Prefer the short word and the concrete noun: `path` over
`outputPath`, `name` over `fileName`, "fail" over "error". Single
letters fit obvious types in local variables and struct receivers
(the type already documents the receiver); function parameters prefer
descriptive names - callers read the signature without seeing the
body. Avoid `-ing` and plural forms in names.

## Single-character collision fallback

When the suggested single character is already taken in scope (e.g. `m` is the
receiver of `*Model`), step through the letters of the full word in order until
a free one is found:

- `markup` → `m`, `a`, `r`, `k`, `u`, `p`
- `notation` → `j` (first letter of json, the dominant format), then `n`, `o`,
  `t`, `a`, `i`
- `error` → `e`, `f`, `g`, `h`

The linter suggests only the first option. Choosing a later letter when there is
a collision is left to the developer.

## HTML custom attributes

`data-*` attributes are HTML's custom attribute mechanism. Use `Custom` as the
replacement term:

```go
CustomPrefix       = "data-"   // was DataPrefix
KeepCustomAttribute bool        // was KeepDataAttribute
```
