# Imports

## Forbidden imports

The `forbidden_import` analyzer refuses these (AST-based, no false
positives on string literals) and names the replacement. New bans go
into the analyzer, not only a doc.

| Forbidden | Use instead |
|-----------|-------------|
| `"flag"` | `"github.com/spf13/pflag"` |
| `"github.com/stretchr/testify"` | `"github.com/funtimecoding/soil/pkg/assert"` |

## Aliases

- **No acronyms** in package names or import aliases.
- **A `constant` package has one name across the tree.** Your own
  domain's `constant` - the one whose domain directory holds the file,
  the deepest when two do - is imported plain. Every other `constant`
  carries its domain, collision or not: `webConstant`,
  `consoleConstant`, `stringsConstant`. The global `pkg/constant` takes
  the module's name: `soil` in soil, `mint` in mint.
- **Other packages alias only when names collide in the same file.** The plain
  name goes, in order, to the standard library; the import sharing the longest
  path with the file; the deeper package; the one already imported plain; the
  first by path. A version-named package (`v1`) never keeps it: `meta` and
  `core` for two colliding `v1`s. Every import that does not keep the plain name
  takes its subsystem - the directory above it, `generated` for an oapi-codegen
  `<path>/server/` - or subsystem plus name when that is taken. A directory
  whose name is banned or no identifier is skipped for the one above it:
  `apimachinery`, not `api`.
- A third-party package whose path yields no usable name keeps the
  alias it has.

gofix and gosourced derive every one of these names from
`pkg/source/imports`; `gofix --realias` applies the `constant` rule.
