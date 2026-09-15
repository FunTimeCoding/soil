# Constants

## Placement

Constants live in the subsystem's `constant/` package (e.g.
`pkg/tool/gobuild/constant/`), in `constant.go`. Large non-vocabulary
assets (CSS, JS, SQL, HTML string blobs) go to domain files beside it
in the same package - `style.go`, `script.go`, `query.go` - keeping
`constant.go` readable as the vocabulary home.

Don't create package-local `constant.go` files - add to the nearest
`constant/` package and export the constant, even when only one
package uses it today. Deeper subpackages share their subsystem
root's `constant/` rather than growing their own: `constant/` lives
only at `pkg/<domain>/` or `pkg/tool/<name>/`, with `pkg/constant` as
the global home. Subsystem vocabularies inside a shared `constant/`
take a per-subsystem domain file.

Prefix a symbol with its subsystem qualifier when collision or
ambiguity forces it - `systemd.go` with `SystemdCommand`, `lite.go`
with `LiteDialectName` (qualifier-first, the `http.MethodGet` shape);
distinctive collision-free names keep their shortness. Either way the
name must stand alone, readable without the file around it.

Group with `const (...)` blocks per semantic domain - multiple groups
with comment headers are the idiom for large vocabularies. When an
iota enum is defined with a named type, the type and its const values
move together into a type-named file (`variable_kind.go` for
`VariableKind`) - separating them across packages creates circular
imports.

`constant/` packages carry no build tags - data is inert (valid while
they import nothing tag-gated).

## The vocabulary leaf

A type that carries behavior (methods, named transforms, a
constructor) doesn't live in `constant/` - it descends into a leaf
package (`types/<name>` at the subsystem root, or a concept package
like `pkg/time/day`) that imports nothing above stdlib, other leaves
(local or library - a downstream repo imports soil leaves freely),
and constant homes. `constant/` then speaks the vocabulary
declaratively - tables typed by the leaf, importing it one-way - and
machinery acts on both. Three layers, one-way imports: `types/<name>`
(vocabulary) <- `constant/` (tables) <- machinery.

Record types never live in `constant/`, methodless or not - the type
gets its leaf, the table stays in the constant home referencing it;
the only type declarations inside `constant/` are enum-shaped named
types over basic kinds (enforced by the `constant_declaration`
analyzer). Enum values never live in the leaf, even when their type
does: leaf-typed values belong to the constant home; when the leaf's
own machinery consumes the values (the cycle case), they move up the
hierarchy to the next constant home instead. Constructor cycles
dissolve in the mirror direction - the constructor descends to a leaf
(`day.New`) so `constant/` builds values without importing machinery.

`constant/` stays behavior-free: zero func declarations;
function-valued table fields are fine when they reference named leaf
functions.

## Reuse and naming

- **Reuse existing constants** - `web/constant.ListenPort`,
  `argument/constant.Name`, `constant.ParameterQuery`. Never hardcode
  strings that already have a constant; the `string_constant`
  analyzer flags them. MCP parameter names shared across tools live
  in `generative/constant`.
- **Constant value naming** - abbreviations and acronyms are avoided
  in constant values, not just names. `Link = "link"` not
  `Link = "url"`. The value should be the honest, full-word form.
- **Semantic constant splitting** - when the same string serves
  different layers (DB column, HTML form field, domain key), use
  separate constants even if the values are identical today:
  `NameFieldKey = "name"` (field system), `NameColumn = "name"`
  (GORM), `NameParameter = "name"` (web routes/forms). They could
  diverge independently.
- **Propagate generic constants toward soil** - when a constant is
  used across multiple tools or packages (e.g. `FormMethod =
  "method"` for HTML forms), it belongs in soil's shared vocabulary
  (e.g. `web/constant`), not duplicated in each consumer.
