# Test Placement Spec

Where test files live. Testing patterns (mocks, lifecycle HTTP,
stores) are in `testing.md`; fixture layout is in `fixture.md`.

Every test file lives in a `unit/` or `integration/`
directory. Source packages carry no `_test.go` files. Analyzer
`testdata/` trees are exempt — fixture `_test.go` files there are
fixtures, not tests.

Both homes sit directly under a domain root: `pkg/<domain>/`,
`pkg/tool/<name>/`, or the repository's equivalent top-level
grouping (an `internal/<domain>/`). Deeper
subsystems consolidate up — a subsystem's integration facets live
at `<root>/integration/<facet>/`, not beside the subsystem.

## unit/

One test package per domain root — `pkg/<domain>/unit/`
or `pkg/tool/<name>/unit/`, never per subpackage. Tests from the
whole subtree consolidate up into it. One test binary per domain
instead of one per package: cold runs stop paying link tax, and a
shared-file edit relinks one binary.

Test homes carry normal package names, so regular files coexist
with tests: helpers live beside the tests as plain `.go` files,
and every `_test.go` file holds at least one test.

All-or-nothing per root: a single in-package `_test.go` left behind
resurrects that package's test binary.

### Black-box only

`package unit`, exercising the public API — the test home is a
separate package from everything it tests, so unexported symbols
stay out of reach. When a test needs white-box access, change the
code, not the test boundary — in order of preference:

1. Use an existing accessor.
2. Export the symbol, or add the missing constructor.
3. Whole-struct compares over the exported surface: `assert.Exported`
   (unexported fields are invisible to the comparator).
4. Test-only fixtures in package source move to the test home;
   fixture constants go to the domain's `constant/` package
   (`Fixture`-prefixed names, or a `sample.go` domain file for
   data blobs). A `constant_test.go` is never a constant home -
   a file by that name holds constant value assertions
   (usage-anchor tests), nothing else.

Never suppress analyzer findings: a `struct_literal` finding on a
test literal means the type needs a constructor. Expected sides of
asserts stay literal, never constants.

Assert helpers mirror the assert shape: expected value first,
right after `t`. When a helper checks several fields at once, the
expected side is one typed struct owned by the tester package,
not a parameter list.

### Naming

Files: subpackage path joined with underscores, plus the concept —
`store_get_document_test.go`, `types_native_event_test.go`.
Root-package tests use the bare concept or the domain name.

Colliding test function names get the subpackage camel prefix
(`TestPort` in a `console/server` file becomes
`TestConsoleServerPort`). A collision the rule manufactures is
resolved by naming what the test actually tests — which also
catches inherited misnomers.

## The line

`unit/` is in-process; `integration/` crosses a process
boundary or needs the environment.

In-memory sqlite stores, fixture-file parsers, and in-process HTTP
servers on dynamic ports stay unit. Shelling out (the go toolchain,
local binaries) or requiring services, credentials, or a display is
integration. Environment-dependent integration tests additionally
carry a build tag (`//go:build local`, `ci`); hermetic ones run
untagged.

## integration/

Facet subpackages, one per surface:

```
integration/
├── base/                  # shared setup: exported, non-test code
├── client/                # generated REST client workflows
├── model_context/         # MCP tools
├── model_context_tester/  # optional per-facet helper package
├── web_interface/         # HTML pages and forms
└── worker/                # poll cycles
```

`base/` exports the stack constructor (`New(t) *Server` with
accessors and `Close()`) as plain package source — the
`pkg/tool/goalertlogd/integration/base` shape. Setup
constructors compile untagged even when the tests they serve are
tagged: environment reads happen at call time, so tags belong on
the test files only.

Analyzer tests live in their own subpackage with their `testdata/`
beside them — relative fixture paths keep working because each
subpackage binary runs in its own directory.

## Interactions

- A test package may import a package with the same name as itself
  (the package's own name is not an identifier in its scope), so
  generated-vs-home clashes like a `client` test package importing
  `generated/client` need no alias — golint de-aliases them.
- golint's stub-test generation is wired off (`stubTest` is `false`
  in the `lint.Lint` call chain); source packages without tests do
  not get stubs demanded back.
- CWD-relative paths in tests anchor via `git.FindDirectory` or
  `fixture.Path`, never parent-depth walking — moves between
  directory depths must not break them.
