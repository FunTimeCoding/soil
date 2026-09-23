# Constructors

- **`New()` in `new.go`** - zero-arg or with params, one per package.
  Zero-arg `New()` is appropriate when call sites use different field
  combinations (callers set fields after construction).
- **A second word behind `New` can mean a package is missing.** When
  the qualifier names a **thing** - `NewStoreTester` - that word is
  the package it is asking for: extract `store_tester/` with a plain
  `New()` rather than qualifying the function. When the qualifier
  names a **flavour** of the same construction it stays:
  `NewEnvironment` reads the same type from the environment,
  `NewMemory` opens the same database in memory, `NewWithStore` takes
  one more collaborator. Thing or flavour is the test. This is the
  diagnostic for the rule above - `New()` per package describes the
  destination, and a thing-shaped qualifier is how you notice you are
  not there yet. Test scaffolding is not exempt: tester packages sit
  beside the tests they serve (`pkg/strings/unit/strings_tester`).
- **`Stub()` in `stub.go`** - returns a zero-value pointer. Used for
  GORM model references (`AutoMigrate`, `Model`, `Delete`) and error
  return paths (when a method fails and needs to return something
  alongside the error). Never use `Stub()` for real construction -
  that's `New()`.
- **`new_environment.go`** - when a package reads from environment
  variables, the environment constructor lives in
  `new_environment.go` and calls `New()` after reading the required
  vars: `func NewEnvironment() *Client { return
  New(environment.Required(constant.HostEnvironment),
  environment.Required(constant.TokenEnvironment)) }`. The base
  `New()` always takes explicit params; `NewEnvironment()` is the
  env-reading wrapper.
- **Option structs** - named after the domain concept, not `Option`
  (e.g. `option.Log`, `option.Build`, `option.Commit`). File named
  after the struct. Constructor in `new.go` returns pointer.

See `../package-design.md` for when types warrant their own packages.
