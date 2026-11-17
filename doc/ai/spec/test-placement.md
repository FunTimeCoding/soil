# Test Placement Spec

Where test files live. Testing patterns (mocks, lifecycle HTTP,
stores) are in `testing.md`; fixture layout is in `fixture.md`.

Every test file lives in a `<path>/unit/` or `<path>/integration/`
directory. Source packages carry no `_test.go` files. Analyzer
`<path>/testdata/` trees are exempt — fixture `_test.go` files there are
fixtures, not tests.

Both homes sit directly under a domain root: `pkg/<domain>/`,
`pkg/tool/<name>/`, or the repository's equivalent top-level
grouping (an `internal/<domain>/`). Deeper
subsystems consolidate up — a subsystem's tests live in its root's
homes, not beside the subsystem.

## unit/

One test package per domain root — `pkg/<domain>/unit/`
or `pkg/tool/<name>/unit/`, never per subpackage. Tests from the
whole subtree consolidate up into it. One test binary per domain
instead of one per package: cold runs stop paying link tax, and a
shared-file edit relinks one binary.

A test home holds `_test.go` files only, and every one holds at least one
test. What the tests share lives in a helper package (Helper packages);
goaudit's `stray_helper` flags anything else beside them.

All-or-nothing per root: a single in-package `_test.go` left behind
resurrects that package's test binary.

A daemon's mounted surfaces - REST, MCP, channel, web pages, the guard
battery - run in-process on a dynamic port, so their tests are unit tests
too. `<path>/unit/base/` exports the stack as an object (Helper packages,
The tester object) in plain package source, and helper packages
sit beside the tests they serve as `<path>/unit/<name>/`
(`pkg/strings/unit/strings_tester`). Setup constructors compile untagged
even when a tagged integration test uses them: environment reads happen at
call time, so tags belong on the test files only - unless the package under
test itself builds only under the tag, when its helper package carries the
same tag.

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
   fixture constants go to the domain's `<path>/constant/` package
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
`store_get_document_test.go`, `alertmanager_alert_filter_alerts_test.go`.
Root-package tests use the bare concept or the domain name.

When consolidation makes two test functions, helpers or files collide, the
moving side is renamed for what it observes - an MCP `TestGetNodeNotFound`
beside a REST one becomes `TestGetNodeToolReportsNotFound` - never a
mechanical prefix, which also catches inherited misnomers. A file name the
facet prefix would double drops the repeat.

## Helper packages

A helper package is named for the role it plays in a test:

- `<subject>_tester` builds the subject under test from its parts, or
  asserts on what it produces (`service_tester`, `markup_tester`); the type
  it hands out is named for what it holds, `Tester` when it is a bundle of
  parts
- `<peer>_server` is a fake network peer the subject talks to - type
  `Server`, `New(t, ...)`, and accessors for what it received
  (`proxmox_server`, `issuer_server`, `upstream_server`)
- `mock_<face>` implements a client's whole face, beside the face
  (`testing.md`)
- `<behaviour>_<role>` is an in-process stand-in whose point is one
  behaviour (`failing_writer`, `refusing_source`, `held_peer`)
- `base` builds the daemon stack
- `command_tester` fronts a command-line tool's tests: `Run(t, client,
  arguments...)` over `go:pkg/terminal/terminal_tester.Run`, holding the
  home's argument lists and seeded mock

`fake_` tells a reader nothing the role does not: a fake peer is a server,
a stand-in is named for its behaviour. An accessor names what the test
wants and fails the test when it is absent - `s.Last()`,
`outage_tester.Guard(t, workers, name)` - never an index into a slice.

Every helper lives in a named package, even one with a single caller, and
never in a `_test.go`. A helper that needs `t` and its own package's tester
is a method on that tester, which holds `t`; its methods never take `t`
themselves (`tester_receiver`). Where every caller only fails on the error a
method returns, the method fails the test itself and returns the value alone
(`go:pkg/tool/goclauded/unit/store_tester.Tester.Announce`); a test wanting
the error calls the subject directly.

Placement follows reach. A helper serving one domain's tests lives in that
domain's `<path>/unit/<name>/`; one serving only tests that cross the line
lives on the integration side, in `<path>/integration/<name>/` or the
subpackage it serves (`go:pkg/tool/gofix/integration/fix_tester`), and may
wrap a unit sibling of the same name
(`go:pkg/tool/goqueryd/integration/store_tester` over
`go:pkg/tool/goqueryd/unit/store_tester`); one serving other domains' tests
lives beside the source package it stands for (`pkg/web/upstream_server`,
`pkg/web/authorization/authorization_tester`). Production API wins over a
test-side copy: look for it before writing a helper
(`locator.NewOrigin`), and for a helper package already playing the role.

### The tester object

A home's `base` and `service_tester` are objects holding `t`, even when they
hold a single function, so one shape reads across homes.
`go:pkg/tool/gohabiticad/unit/base` and
`go:pkg/tool/gobandd/unit/service_tester` are the patterns.

- `New(t, o ...Option)` builds the default parts, and a test swaps one with a
  `With…` option, written when a caller needs it. Options add: a part most
  callers want is still opt-in, named by its callers. A part every caller
  names, with no sensible default, stays a positional argument. A
  `NewWith…` beside `New`, or a hook editing the daemon's option, gives way
  to options.
- Parts are exported fields, for reading and for driving the fake
  (`s.Source.Add(…)`, `s.Reporter.Events()`), never for configuring.
- `Service()` and `Daemon()` build on first call and hand back that instance
  after, so the tester's own helpers and the test share one subject. A
  tester hands out its subject through them, never as a bare return value,
  a field built at `New`, or a fresh build per call.
- A tester built on another takes that tester's options and passes them
  through, so each option is written once, beside the part it sets;
  `base` embeds its `service_tester`.
- `base` holds the daemon stack and hands out what the daemon serves: the
  mock as `Source`, `Daemon()`, `Served()` mounted through
  `model_context_server.Mount`, and `Tools()` or `GeneratedClient()` for the
  surfaces it has. `client_tester` and `outage_server` build on it. A
  daemon whose `Mount` takes more than reporter, recorder and mux, or
  passes its own recorder, keeps `model_context_server.New` and its
  closure.
- Wrappers over the stack - `model_context_tester`, `web_service_tester`,
  `web_interface_tester`, `daemon_tester` - fold into `base`, their helpers
  becoming its methods, so a test holds one object. A second object with
  methods keeps a package named for what it holds
  (`go:pkg/tool/goclauded/unit/session_tester`, handed out by
  `base.NewSession()`), as a package holds one receiver type
  (`type_receiver`).
- A multi-instance home builds one mock per name: `WithInstances(names
  ...string)`, `alfa` alone by default. Tests reach a mock through accessors
  that fail the test naming a missing instance - `Alfa()` and `Bravo()` for
  the fixture pair, `Instance(name)` for any other - handing out the mock, or
  a `Clients` struct where an instance has several
  (`go:pkg/tool/goproxmoxd/unit/service_tester`).
- Fixtures filling a tester's parts are its methods; a fixture the subject
  reads as it starts is an option instead, as a method would run too late. A
  builder preparing a fixture returns the tester, not its directory.

## The line

`<path>/unit/` is in-process; `<path>/integration/` crosses a process
boundary or needs the environment.

In-memory sqlite stores, fixture-file parsers, and in-process HTTP
servers on dynamic ports stay unit. Shelling out (the go toolchain,
local binaries) or requiring services, credentials, or a display is
integration. Loading code with `go/packages` runs `go list`, so any test
that loads packages crosses the line - analyzer tests and the source tools
alike. Environment-dependent integration tests additionally carry a build
tag naming the environment (`//go:build local`, `ci` and `browser` among
them); hermetic ones run untagged. Only `local` reaches the gate — `task test`
passes that tag alone, so a test under any other tag runs solely
when invoked explicitly. Choosing the tag chooses whether anything
will notice the test failing.

## integration/

Only what crosses the line: tests tagged for an environment, tests that
launch a process, and tests that load code through `go/packages`. Each sits
in a subpackage named for what it reaches - `toolchain`, `repository`,
`postgres` - or in the root `<path>/integration/` package when the whole
domain is what it reaches.

Analyzer tests live in their own subpackage with their `<path>/testdata/`
beside them — relative fixture paths keep working because each
subpackage binary runs in its own directory.

## Checking the line

A misplaced test runs the same either way, so placement drifts silently.
Two counts check it, from the repository root.

Untagged tests under `<path>/integration/` - each must launch a process or
load code:

    find pkg -path '*/integration/*' -name '*_test.go' -not -path '*/testdata/*' -exec grep -L '^//go:build' {} +

Tests under `<path>/unit/` that cross - run every unit package with an
empty `PATH`, where anything starting a program by name fails:

    go test -count=1 -json -exec '/usr/bin/env PATH=/nonexistent' $(go list ./... | grep '/unit$')

Read the second count as passes, not failures: a test that panics ends its
binary, and every later test goes unreported rather than failing. Compare
what passes against a normal run, and rerun a package that lost tests one
test at a time with `go test -exec ... -run`. Never probe a copied test
binary - `runtime.RunningFromBinary` reads `os.Args[0]` and fails anywhere
outside `go-build`. A program started by absolute path does not show.

## Interactions

- A test package may import a package with the same name as itself
  (the package's own name is not an identifier in its scope), so
  generated-vs-home clashes like a `client` test package importing
  `<path>/generated/client/` need no alias — gofix de-aliases them.
- golint has no stub-test generation; a source package without
  tests is never asked for one.
- CWD-relative paths in tests anchor via `git.FindDirectory` or
  `fixture.Path`, never parent-depth walking — moves between
  directory depths must not break them.
