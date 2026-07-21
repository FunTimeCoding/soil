# Package Taxonomy

When a service tool outgrows the base tree (see `service-tool.md`), additional
packages emerge to separate concerns. This spec covers their roles, dependency
direction, and the criteria for introducing them.

## Package Roles

| Package | Role | Imports from |
|---------|------|-------------|
| `types/` | Pure domain types - validation schemas, metadata, enums. No persistence, no external service deps. | stdlib, shared libs |
| `constant/` | Domain constants, enum values, instances, registries. Purely declarative - data definitions only, no logic. Promoted from `constant.go` when it outgrows a single file. | `types/`, shared libs |
| `model/` | Persistence-aware entities - gorm-tagged structs, JSON field parsing, DB convenience methods. | `types/`, shared libs, ORM |
| `store/` | Data access - CRUD operations on model types. | `model/`, ORM |
| `generated/client/` | oapi-codegen output - typed REST client. Machine output, don't edit. | stdlib |
| `generated/server/` | oapi-codegen output - `ServerInterface`, `HandlerFromMux`, types. Machine output, don't edit. | stdlib |
| `client/` | Domain wrapper for the generated REST client. Used by the CLI. | `generated/client/` |
| `server/` | REST implementation. Implements the generated `ServerInterface`. | `generated/server/`, `store/`, `convert/` |
| `model_context/` | MCP tool implementations. Each method is a handler. | `store/`, `constant/`, `convert/`, `response/` |
| `convert/` | Type filtering shared by `model_context/` and `server/`. | `types/`, `generated/server/` |
| `web/` | HTML rendering (gomponents). Holds `*view.View` on Server. Flat; file-prefix grouping. | `store/`, `constant/`, `model/`, `web/view/` |
| `integration_test/` | Cross-package tests using only the public API. Facet subpackages with shared setup in `base/` (see `test-placement.md`). | all exported packages |

### Shared web packages (soil)

| Package | Role |
|---------|------|
| `web/layout/` | Fluent HTML page builder. `New(identity)` + `With*` + `Clone()` + `Render()`. |
| `web/layout/navigation_item/` | `New(path, label)` / `NewExternal(path, label)` for nav links. |
| `web/palette/` | Command palette. `NewRegistry()`, `Register()`, `NewServe(registry)`. fzf-style fuzzy matching. |
| `web/view/` | HTTP layer wrapping layout. `RenderPage`, `RenderFragment`, `IsExtendedRequest`. |
| `web/theme/constant/` | CSS palette constants (pico.css overrides). Downstream repos can define additional palettes. |

## Dependency Direction

Imports flow downward; never upward or sideways between peers at the same layer.

```
types/
  ↑
constant/   model/
  ↑           ↑
  +---+     store/        generated/client/   generated/server/
  |   |       ↑                ↑                    ↑
  |   +-------+-------+-------+----------+----------+
  |           |       |       |          |
server/  model_context/  web/  client/  integration_test/
                          ↑
                    web/view/ → web/layout/ → web/layout/navigation_item/
```

`server/`, `model_context/`, `web/`, and `client/` are leaf packages -
they import from lower layers but nothing imports from them (except
tests and the run wiring).

Cycle avoidance: if two packages need each other's types, the shared types
belong in `types/` (or an interface in `face/`).

## When to Promote

### `constant.go` → `constant/`

Promote when the file exceeds ~100 lines, or when constants reference types
from a dedicated types package. Inside `constant/`:

- `constant.go` - simple string/int constants, iota enums
- One file per domain concept: `field.go`, `list.go`, `image.go`
- Instance variables (constructed values like template instances, image
  definitions) get `<concept>_<name>.go` files
- Registries (lookup functions, slices of all instances) get
  `<concept>_list.go`, `<concept>_by_name.go`

### No package → `types/`

Introduce when domain types with methods appear that are not persistence
entities. `types/` is a subtree - each type gets its own sub-package per the
one-struct rule:

```
types/
├── template/       # Template struct + validation methods
├── field_meta/     # FieldMeta struct + builder
└── image/          # Image struct + attribute accessors
```

`types/` packages have no persistence deps (no ORM, no store). They define
what the domain looks like; `model/` defines how it's stored.

### No package → `model/`

Introduce when persistence entities need sub-packages (multiple entity types
with receivers). `model/` is a subtree:

```
model/
├── machine/         # Machine struct + Map(), Format()
└── machine_form/    # Form struct + FieldMap()
```

If there's only one entity, `model/` can stay flat (one struct file, one
constructor). Extract sub-packages when a second entity with receivers appears.

### No package → `helper/`

**Avoid.** `helper/` and `util/` are junk-drawer anti-patterns. Every function
has a proper home:

- Query/classification logic on domain data → registry struct (see below)
- Single-consumer utilities → private function in the consumer package
- Thin wrappers (e.g. `Encode()` around a stdlib call) → inline at call site

If multiple sibling packages share a function, that usually means it operates
on domain data and belongs in a registry or a domain-adjacent package - not a
catch-all bucket.

### When to extract a registry

When `constant/` accumulates query functions (lookups, predicates,
classification, splitting) that operate on its declarative data, consider
extracting them into a struct with methods in a dedicated package (e.g.
`<concept>_registry/`). The struct holds references to the data; `constant/`
instantiates it and exports the instance. Consumers call
`constant.Registry.Method()`.

This keeps `constant/` purely declarative and gives query logic a named,
testable home without scattering it across consumer packages.

### No package → `integration_test/`

Introduce when cross-package tests exist. Tests live in facet
subpackages (`client/`, `model_context/`, `web_interface/`, ...)
with shared setup exported from `base/` — layout, naming, and the
unit/integration line are in `test-placement.md`.

## Flat Package Guidelines

### When flat is fine

A package with 20–60 files that share a single struct or a common import
context is fine flat. File-name prefixes provide grouping:

```
server/
├── get_machine.go
├── post_machine.go
├── delete_machine.go
├── get_image.go
├── machine_to_api.go       # conversion helpers
└── write_json.go           # response helper
```

Counter-indicators for splitting: shared struct (Router, Server), shared
helper functions, files under 80 LOC each.

### When to split

Split when a package exceeds ~80 files, or when a distinct subset has no
dependency on the shared struct. Example: pure UI components that don't need
server state could become `web/component/`.

### Naming collisions in flat packages

When a flat package holds instances of multiple domain concepts (image
instances and template instances), use a suffix to disambiguate:

```go
var Debian = image.New("debian", ...)   // image instance
var DebianTemplate = template.New("debian", ...)  // template instance
```

The suffix matches the type name. This avoids collisions like `Debian`
(image) vs `Debian` (template).

## Progression

A service tool typically evolves in this order:

1. **Base tree** - `option/`, `store/`, `server/` (or `web/`), `run.go`
2. **MCP layer** - `model_context/` + tool registration
3. **Constants outgrow** - promote `constant.go` → `constant/`
4. **Types emerge** - extract `types/<concept>/` for non-persistence domain types
5. **Model splits** - extract `model/<entity>/` when multiple entities appear
6. **Query logic extracts** - operations on constant data move to a registry struct
7. **Tests consolidate** - collect cross-package tests in `integration_test/`

Not every service reaches every stage. Promote only when the criteria above
are met.
