# Coding Conventions

How code reads here. The linters enforce and fix what can be checked
mechanically; these leaves carry the judgment the tooling cannot.

- `style.md` - statement-level style: nesting, chaining, pointers,
  defer placement, string building
- `formatting.md` - what the fix pipeline owns: blank lines, line
  length, call layout
- `files.md` - one identity per file, file naming
- `constants.md` - where constants live, the vocabulary leaf,
  semantic splitting
- `constructors.md` - `New`, `Stub`, `NewEnvironment`, option structs
- `imports.md` - forbidden imports, alias rules
- `interfaces.md` - generic vs domain interfaces, mocks

Sibling specs at the parent level: `../naming.md` for identifier
rules, `../comments.md` for comment discipline, `../testing.md` for
testing philosophy and patterns, `../error-handling/` for the
PanicOnError strategy, `../stdlib-wrappers.md` for the wrapper
tables, `../package-design.md` for when a type earns its own package.
