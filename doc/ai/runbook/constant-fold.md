# Constant fold

Migrating a repo to the constant placement rule: no `constant.go`
outside a `constant/` package, `constant/` only at `pkg/<domain>/`
or `pkg/tool/<name>/`, no subpackage trees, `pkg/constant` as the
global home. The rule itself lives in `doc/ai/spec/conventions.md`
(Constants section); this runbook is the migration procedure for a
downstream repo that consumes soil.

## The gate

goaudit carries three placement checks: `constant_file` (stray
constant.go), `constant_depth` (constant/ outside a domain or tool
root), `constant_nested` (constant/ contains subpackages). They
are currently parked behind an `if false {}` around the
ConstantPlacement loop in `pkg/tool/goaudit/run_headless.go` until
every downstream repo has zeroed. To census a repo, remove the
`if false` locally, rebuild goaudit, and run `goaudit .` - its
output is the only census; only the tool declares done. No
exemptions, no allowlist: testdata is skipped (Go toolchain
semantics) and `unit_test`/`integration_test` homes inside
`constant/` pass; everything else counts.

## Procedure

1. **Census first.** Run the unparked gate. The finding list is
   the complete work list - nothing outside it is left over,
   nothing inside it is exempt. Prove the census is alive before
   trusting a zero: analyzers need package patterns (`./...` -
   a repo root without Go files reports a silent or panicking
   fake zero), and a deliberately planted stray in a scratch
   package must flag before you believe the empty output.
   Delete the probe after.
2. **Classify before moving.** Read every flagged file and tag it
   by kind - the treatment differs and package-by-package
   flattening applied blindly gets it wrong:
   - **vocab** - plain constants; fold with a qualifier-first
     subsystem prefix on names not already carrying it
     (`http.MethodGet` shape). The fold isn't done until the
     name stands alone.
   - **tree** - a constant subpackage whose name was the
     qualifier; becomes a domain file (`<subsystem>.go`) with the
     qualifier moved into every name.
   - **enum** - typed iota family; type and values move together
     into a type-named file.
   - **table** - curated var (slice/map); rides with its
     vocabulary.
   - **blob** - text asset (CSS, SQL, selectors, templates);
     asset domain file beside constant.go.
   - **constructed** - var built from constructors (option
     presets, registries, compiled regexes, error vars,
     templates). These are constants in everything but Go's
     keyword and fold like the rest. One mechanical limit: if the
     constructors come from the same domain the constant home
     serves, the fold can cycle - fix the arrow direction
     (promote the generic enum the machinery needs to
     `pkg/constant` so machinery becomes self-contained, then the
     domain home imports machinery one-way) rather than
     relocating constants out of constant packages.
   - **misfile** - a type or func living under a constant name;
     it moves OUT to its owning package (or `git mv` in place),
     never into the constant home.
   - **discriminator** - a constant pair (or field-plus-values)
     that duplicates what another field's presence already says
     (a kind string where exactly one kind sets its companion
     field). It dissolves into a derived predicate method
     instead of moving - check serialization consumers before
     deleting the field.
   - **single** - one or two constants; fold to the domain root.
     Cross-domain promotion is a separate scope.
3. **Fold in approval-paced waves**, smallest domain first, one
   presented plan (names before → after) per wave, census count
   dropping between waves. Prefix renames go BEFORE moves -
   collisions block them.
4. **Verify per wave**: build, full lint, tests. Expect the
   duplicate-literal ripple as constants go public - sweep the
   repo per flagged literal, don't iterate one file at a time.

## Rulings that generalize

- Same value, same set (could be one enum): merge. Sets differing
  in any member: separate constants. Same value, different
  meaning: own constant or an honest literal rename - never a
  suppression.
- Enum values never live in vocabulary/leaf packages, even when
  their type does: leaf-typed values move to the constant home,
  which imports the leaf one-way. When the leaf's own machinery
  consumes the values - the cycle case - they move UP the
  hierarchy to the next constant home instead of staying put.
  Values take the type or package qualifier as they move; the
  name must stand alone.
- Export-all applies to internals (analyzer orderings, tuning
  knobs, scoring tables) - in-package residence is the lint gap
  new sessions recreate strays through.
- Error vars keep the `Error*`-first form (`ErrorNotFound`);
  provider-prefixed error names break staticcheck ST1012 - merge
  same-meaning error vars per domain instead.
- Comment preservation is load-bearing: `#nosec` markers, spec
  structure headers, design-rationale comments ride the move
  verbatim. Multi-group `const (...)` blocks are the curated
  idiom - no normalization pass.

## Mechanics and traps

Type-checked AST tooling (gosourced's `move_symbols` and
`rename_symbol` where available, otherwise careful hand edits) is
the safe hands - regex over Go identifiers breaks parens, misses
references, and collides aliases. The recurring traps: a moved
declaration referencing its target package refuses (split by
hand); moving a struct type does not check field visibility
(export fields with the type); single-line `import "..."` forms
dodge alias-detection greps; test files are invisible to
`go build` (typecheck via the linter or `go vet`); alias
naturalization after moves - the more local package keeps the
bare `constant` name, the displaced global takes `library`,
foreign subsystems take their subsystem name. When a
vocabulary's references are qualified and concentrated (a
value-assertion anchor test plus a few consumers), hand-rewriting
the declaration file and sweeping consumers with word-boundary
qualified replacements beats dozens of tool renames - build and
lint prove the sweep. Two clean-up traps after tool moves: moved
symbols can land one-file-per-symbol - consolidate into the
curated `const (...)` block file; and package renames don't touch
file basenames - test files elsewhere that embed the old package
name in their filename need a manual `git mv` (renaming basenames
by string match risks collateral, so the tool correctly refuses
the rabbit hole).
