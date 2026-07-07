# Console Status Spec

Fluent builder for structured single-line console output with optional detail lines.

## Output Format

A status line has a **header** of pipe-separated bubbles, optionally followed by indented **detail lines**:

```
bubble1 | bubble2 | bubble3
  detail line 1
  detail line 2
```

With indentation (each level = 2 spaces):

```
  bubble1 | bubble2 | bubble3
    detail line 1
```

## Status Builder

`status.New(f *option.Format) *Status`

### Bubble Methods (header values, pipe-separated)

All return `*Status` for chaining.

- `String(v ...string)` - string bubbles (empty strings skipped)
- `Integer(v ...int)` - integer bubbles
- `Integer32(v ...int32)` - int32 bubbles
- `Integer64(v ...int64)` - int64 bubbles
- `Boolean(v ...bool)` - boolean bubbles

### Detail Line Methods

- `Line(format, a...)` - extended line (shown when `ShowExtended`)
- `Lines(v ...string)` - multiple extended lines (each indented with `"  "`)
- `TagLine(tag, format, a...)` - line grouped by tag (shown when tag is in `format.Tags`)
- `DetailLink(url, label, prefix)` - link with format-aware rendering (see Link Rendering below)

`ShowExtended` controls whether `Line`/`Lines` detail lines are included. When off, only the bubble header renders - useful for dense list views. `TagLine` and `DetailLink` are independent of `ShowExtended`; they have their own visibility rules (tags and link mode respectively).

### Debug Methods

- `Raw(a, title)` - `"  {title}: {%+v}"` (shown when `ShowRaw`, magenta when color)
- `RawList(a)` - `Raw(a, "RawList")`
- `RawDetail(a)` - `Raw(a, "RawDetail")`

### Terminal

- `Format() string` - renders the output

## Render Order

1. Header: `[indent][bubble1 | bubble2 | ...]`
2. Tag lines: only those whose tag is in `format.Tags`
3. Extended lines: only when `format.ShowExtended`

## Format Option

`option.New() *Format`

Fluent builder methods (all return `*Format`):

| Method            | Field          | Effect                                            |
|-------------------|----------------|---------------------------------------------------|
| `Color()`         | `UseColor`     | Enable ANSI color                                 |
| `Compact()`       | `UseCompact`   | Use short names/acronyms                          |
| `Extended()`      | `ShowExtended` | Show extended detail lines                        |
| `Raw()`           | `ShowRaw`      | Show debug `%+v` output                           |
| `Tag(v...)`       | `Tags`         | Add tag filters                                   |
| `RemoveTag(v...)` | `Tags`         | Remove tag filters                                |
| `Indent(i)`       | `Indentation`  | Indentation level (each = 2 spaces)               |
| `Filter(k, v)`    | `Filters`      | Add key-value pair filter                         |
| `Copy()`          | -              | Shallow copy (for customization without mutation) |

Query methods: `HasTag(v)`, `HasFilter(k, v)`, `HasFilterKey(k)`.

### Presets

```go
option.Color         = New().Color()
option.ExtendedColor = New().Extended().Color()
```

## Tags

Constants in `status/tag/`:

```
Age, Assignee, Category, Changes, Cluster, Comment, Concerns, Copyable,
Dense, Description, Emoji, Filter, Fingerprint, Graph, Host, Identifier,
Instance, Investigate, Key, Labels, Markdown, Name, Project, Runbook,
Score, State, Status, Timestamp, Type, Usage, Wiki
```

Tags serve two purposes:

1. **Display toggles** - `TagLine` content only renders when its tag is active. Example: `s.TagLine(tag.Comment, "  %s", comment)` only shows when `tag.Comment` is in the format.
2. **Semantic hints** - checked in rendering methods to switch output mode. `tag.Markdown` selects Markdown links, `tag.Copyable` selects plain-text links.

## Link Rendering

`DetailLink(url, label, prefix)` renders a link based on active tags:

- **Default (no tag)** - OSC8 hyperlink as a bubble: `label` (clickable, no color, terminal provides underline). Dense.
- **`tag.Markdown`** - Markdown detail line: `[label](url)` or `prefix: [label](url)`
- **`tag.Copyable`** - plain URL detail line: `url` or `prefix: url`

Parameters:

- `url` - the URL (no-op if empty)
- `label` - short clickable text for OSC8/markdown (defaults to `"Link"` if empty)
- `prefix` - annotation prepended in copyable/markdown mode (omitted in OSC8 mode, ignored if empty)

### Examples

```
# Default (OSC8) - url="https://jira.example.com/PROJ-123", label="Jira", prefix=""
123 | open | Fix auth bug | 3d | Jira

# Markdown - same params
123 | open | Fix auth bug | 3d
  [Jira](https://jira.example.com/PROJ-123)

# Copyable - same params
123 | open | Fix auth bug | 3d
  https://jira.example.com/PROJ-123

# Copyable with prefix - url="https://grafana.example.com/...", label="Grafana", prefix="Prometheus 24h graph"
alert-name | firing | 2h
  Prometheus 24h graph: https://grafana.example.com/...

# Markdown with prefix - same params
alert-name | firing | 2h
  Prometheus 24h graph: [Grafana](https://grafana.example.com/...)
```

## Formattable Interface

```go
// face/formattable.go
type Formattable interface {
    Format(f *option.Format) string
    Meta() *description.Description
}
```

Every domain entity implements this interface.

## Entity Format Convention

### Structure

Each entity has:

- `format.go` - main `Format(f)` method, builds status line
- `format_<field>.go` - private per-field formatter reading `f.UseColor`, `f.UseCompact`, `f.HasTag()`
- `constant.go` - fallback display values (`NoPages`, `NoPrice`, etc.)

### Typical Format Method

```go
func (r *Request) Format(f *option.Format) string {
    s := status.New(f).Integer64(
        r.Project,
        r.Identifier,
    ).String(
        r.formatState(f),
        r.formatTitle(f),
        r.formatAge(f),
    ).RawList(r.Raw)

    s.DetailLink(r.Link, "GitLab", "")

    return s.Format()
}
```

### Typical Field Formatter

```go
func (r *Request) formatState(f *option.Format) string {
    result := r.State

    if f.UseColor {
        if result == OpenAlias {
            return console.Yellow("%s", result)
        }

        if result == constant.ClosedState {
            return console.Green("%s", result)
        }
    }

    return result
}
```

Pattern: check `f.UseColor`, apply ANSI color. Check empty value, return fallback constant (yellow when colored to signal missing data).

### Field Formatter with Fallback

```go
func (b *Book) formatPages(f *option.Format) string {
    if b.Pages == "" {
        if f.UseColor {
            return console.Yellow("%s", NoPages)
        }

        return NoPages
    }

    return b.Pages
}
```

## Domain Format Presets

Each domain defines a base format in its `constant/constant.go`:

```go
var (
    Format      = option.Color.Copy()
    CheckFormat = Format.Copy().Tag(tag.Project)
)
```

CLI entry points copy and customize:

```go
f := option.ExtendedColor.Copy().Tag(tag.Identifier)
```

## Format Threading

`*option.Format` is passed as a function argument through the entire rendering chain. It is never stored on entities - only on `constant` package vars and CLI-level variables. The format flows: CLI entry point -> `entity.Format(f)` -> `entity.format*(f)`.
