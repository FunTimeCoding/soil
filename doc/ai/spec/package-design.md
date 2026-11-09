# Package Design

One type with methods per package. This is the core rule for package
scope - it drives when to extract sub-packages and how to organize types.

## The Rule

Each package has at most one type that carries methods (receivers). That
type is the package's identity. Its file is named after the type in
snake_case. Constructor lives in `new.go`.

Any named type counts: a struct, a function type adapting an interface
(`http.HandlerFunc`'s shape), a `type Kind int` with `String()`. Each
carries behaviour, so each gets its own package. A face, a package of
interfaces, holds no adapter beside them; each adapter is a sub-package
of the face named after it.

A receiver **taints** its package: the package holds that one type and no
other struct, exported or unexported. A data struct beside it moves out - to a
bag, the types directory, a subpackage, a sibling or a parent, whichever reads
right. A bag tainted by one stray receiver moves that struct out instead, or
the struct loses the receiver when the method has no callers.

## Bag Packages vs Own Packages

A bag holds several structs and no receivers, grouped by role rather than by
type. The roles: `argument` (MCP tool inputs), `request` (outbound wire
bodies), `response` (wire bodies either direction), `query` (read parameters),
`record` (rows, scan targets), `result` (operation outcomes), `metric`
(Prometheus holders), `call` (mock and tester recordings), `view` (shapes built
for web pages). Inside a bag a redundant prefix or suffix drops:
`go:pkg/tool/goqueryd/store/result.Search`, not `result.SearchResult`. A lone
struct of a role is its own package
(`go:pkg/tool/gokubernetesd/types/apply_result.Result`), never a one-struct bag.

When a type has distinct domain identity within our code - or gains
behavior (formatters, receivers, helpers) - it moves to its own
package named after the type. Examples: `response.Task` → `task.Task`,
`session.History` → `history.History`, `session.Screen` →
`screen.Screen`.

The migration is one-directional: bag → own package.

Decision rule: **our types get own packages, their shapes stay in
bags.**

## Where a Type Goes

- **Moves never cycle**: what something builds or reads stays below it. When a
  receiver struct's fields use bag types the bag itself builds on, the leaves
  go down first - their closure to a leaf bag, the receiver struct above it. A
  package can end as layers: data at the bottom, the receiver struct that
  builds it above, functions using both on top.
- **A leaf follows its reach**: read only inside its head, it is a subpackage
  below the head; read outside, it goes to the types directory. A head holds its
  parts by reference, built with `New`, unless a reader depends on a copy by
  value.
- **`<tool>/types/`** is a grouping directory, never a package: own packages of
  a tool's types not categorized further. A lone struct can go there; a pair
  from one exchange is one package named after the exchange (`register` with
  `Request`, `Response`); a configuration struct is `<name>_option` with
  `Option`. `<tool>/model/` is the same for gorm models, apart from the store
  that migrates them. A library has no tool level: its types directory sits at
  the domain root (`<domain>/types/`).
- **Collisions**: the name takes its parent as a prefix. Field-identical twins
  merge into one type. From the third type sharing a prefix, they move into a
  subdirectory named by it (`issue` under `<domain>/types/github/`).

## Pointer Convention

Convert functions and domain constructors return pointers
(`*server.Tunnel`, not `server.Tunnel`). Slice returns are slices
of pointers (`[]*server.Tunnel`). This applies to:

- `<path>/convert/` functions (netbox, proxmox, habitica, atlassian)
- `New()` and `Stub()` constructors
- Source interface methods (`HabiticaSource`, `SublimeSource`, etc.)

In OpenAPI specs, optional arrays of objects use `nullable: true`
with `allOf` wrapping to generate `*[]*Type`:

```yaml
checklist:
  type: array
  items:
    nullable: true
    allOf:
    - $ref: "#/components/schemas/ChecklistItem"
```

## When to Extract

When a second type with methods appears in a package, extract it and
its helper functions into a sub-package. The sub-package is named after
the type. The type inside drops the package prefix:

```
pkg/tool/gosalt/analyzer/    → type Live struct
pkg/tool/goalertlogd/store/  → type Store struct
pkg/tool/goalertlogd/worker/ → type Worker struct
```

Not `analyzer.Analyzer` or `store.StoreStruct` - the package name
provides context.

## Interfaces

Domain-specific interfaces referencing the extracted struct's types are
defined at the consumer, not in the sub-package - unless the sub-package
owns the domain. The sub-package's mock lives in `mock_*/` next to the
real implementation.

## What Stays

Package-level functions without a receiver are fine alongside the single type
with methods. Constants live in the root's `<path>/constant/` package. Each
function gets its own file.

Tool entrypoint packages (`pkg/tool/go<name>/`) often have no type with
methods at all - just `Main()`, `run()`, and orchestration functions.
This is normal. The rule applies when types with methods do appear.
