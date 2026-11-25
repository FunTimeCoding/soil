# Testing Spec

Testing philosophy and integration patterns for service tools. Where
test files live is `test-placement.md`.

- `double.md` - mock clients, how a double answers, fake upstreams
- `served.md` - served surfaces, HTTP clients in tests, the guard battery
- `outage.md` - peer guard outage tables and what proves a guard

## Philosophy

- **Unix approach: silence on success, noise on failure.** No
  progress logs (`t.Log("doing X...")`), no success logs
  (`t.Logf("✓ X succeeded")`). Only `t.Fatalf()` for must-succeed
  setup, `t.Errorf()` for validation failures.
- **Assert package** - `pkg/assert` for every assertion, over raw `if`
  checks, argument order `(t, expected, actual)`. The `forbidden_import`
  analyzer refuses testify.
- **No defensive programming in tests** - "should work or blow up".
  Extract logic into helpers that panic instead of returning errors;
  helpers return only success values.
- **A new test is proven by breaking what it guards** - break the code
  path, not its inputs, keep it compiling, watch the test fail, restore.
  `go test -overlay` breaks without touching the tree; `gobreak` runs a file
  of breaks. One break at a time: a panicking break aborts the package and
  hides every other result. Run the suite green first, then read which
  assertion fails - a red from a build error, a setup step or a test already
  failing proves nothing. `no tests to run`, or a break whose pattern no
  longer matches, reads green while breaking nothing. A break the test
  survives names a branch no test reaches, or code nothing needs. Break, run
  and restore go one after another, never as parallel tool calls: a run
  beside its edit can read the tree before the edit lands.
- **A split makes wiring a thing to break** - splitting a type turns wiring it
  held implicitly into constructor arguments that can be handed wrong, and the
  tests written before the split could not have made that mistake. Break each
  new hand-off once.
- **A filter test keeps a row and drops a row**, and counts the rows
  directly - a filter dropping everything passes a test asserting only
  the drop, and a read that deletes what it filters hides a missing
  cleanup.

## Test names

A test name states the subject and what holds of it: the thing under test
first, then its outcome in plain software words - `TestNpxBlocked`,
`TestGitReadersAllowed`, `TestRunFailsOnUnknownExpect`. No articles
(`test_article` flags them), no `Is` or `Are`, and no verbs that give code a
voice or a mind - answers, refuses, says, keeps, sees. It is chopped: no tail
the outcome already tells apart - `TestOpenAIServerListingModelsReachable`
beside `TestUnauthorizedOpenAIServerUnreachable`. Two tests of one behaviour at
different layers carry the layer. Prepositions stay when they carry meaning;
`preposition_name` doesn't read test names.

## Service First

Behaviour is tested where it lives: the service, called directly. A
surface - MCP tool, REST route, web handler - parses its input, calls the
service and maps the error, so a test through it repeats what the
service test already holds. A surface earns a test only for what it alone
decides: an error it classifies differently (a not-found answered rather
than captured), a parse no other surface shares, a form round-trip a
reader would not notice breaking. Logic two surfaces share moves into the
service, where one test covers both. The guard battery is the standing
exception - the auth contract belongs to the mount, not the service.

## Store Testing

Gorm stores test against in-memory sqlite by default:

```go
m := lite.NewMemory()
t.Cleanup(func() { relational.CloseMapper(m) })
migrate.AutoMigrate(m)
s := store.New(m)
```

The store borrows the mapper, so the test closes it, as the daemon does
in production (`database.md`). Raw `go:database/sql` stores use
`store.New(connection.NewMemory())` the same way. Use the file openers
with `t.TempDir()` only when the test genuinely needs a file on disk. Never call
`gorm.Open` directly - the `restricted_call` analyzer confines it
to `pkg/relational/` (see `database.md`).

## Poll Cycle Testing

Test poll logic directly by calling `Poll()` - no goroutine or ticker needed:

```go
w := worker.New(v.Poll, l, r, time.Minute)
w.Poll()
records := s.ByName("HighMemory")
assert.Count(t, 1, records)
```
