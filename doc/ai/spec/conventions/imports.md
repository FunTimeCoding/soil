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
- Alias only when package names collide in the same file.
- Alias by the subsystem name, not the role in the current file:
  - `generative` for `pkg/generative/model_context/server` (when it
    collides with the tool's own `server/`)
  - `generated` for the oapi-codegen `server/` package (when it
    collides with another `server/`)
  - `webConstant` for `pkg/web/constant` (when it collides with a
    local `constant/`)
- **The more local package keeps the natural name** - when two
  packages share a last segment (e.g., both called `server`), the
  tool's own package is imported unaliased; the shared infrastructure
  package gets the alias. In files that only import one of them, no
  alias is needed.
