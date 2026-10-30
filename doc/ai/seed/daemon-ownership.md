# Daemon ownership

Who owns and closes what in a daemon, how it answers, and how other code
reaches it. The rules are settled and live in the specs: one service per
daemon, built from configuration, the only owner `Run` holds
(`service-tool.md`, Service; `lifecycle.md`; `http-client.md`, Owning;
`database.md`, Store Shape; `testing.md`, Mock Clients); errors answered by
class (`error-handling/README.md`, `external-api.md`); peers guarded, and
waited on only when a surface has no answer without them
(`error-handling/infrastructure.md`, Required and guarded peers, The guard);
one daemon client with a full face and a full mock (`generated-api.md`, The
Daemon Client); a domain face in `<domain>/face/` and its mock beside the
real client (`conventions/interfaces.md`). This seed holds what the specs do
not: the working model, where each soil daemon stands, the order, and the
craft. A consuming repository keeps its own daemons' census and order in a
seed of its own and follows this one for everything else.

## Where the work stopped

Twenty-one soil daemons meet every row of the target (Census); goagentd is
next (Order).

## Working model

Work goes one daemon at a time, in loops. Each loop is a survey, a sketch
agreed before any code, then one change with its tests, and leaves soil and
every consuming repository building and tested.

- The survey reads the surfaces - the REST server, the MCP server, the web
  package - not just `run.go`, and counts the consumers of every client
  touched, by receiver. The logic a route holds, which surfaces bypass the
  service, and how each surface answers errors decide what the service
  becomes. A census line is a hypothesis; the survey decides the loops
- A large daemon takes three loops, a small one folds them into one change:
  - A, ownership - the service builds and closes everything; every surface
    goes through it, holding no client of its own
  - B, honest answers - the upstream client classifies its peer's statuses;
    the service checks what it is asked to write; the spec carries a
    response per class; handlers answer by class. With a peer, P - the
    guard - lands with B, or an outage answers quietly with nobody reporting
    it
  - C, the daemon client - typed operations, a full face and mock, the
    consumers moved onto it
- A can land before B: with the client still unclassified, moving every call
  behind the service changes no answer, and the unchanged tests say so
- Every new test is proven by breaking the code it guards and watching it
  fail, then restoring
- The gate is the build of soil and every consuming repository against the
  local checkout, the affected test suites under `-race` - every consumer's
  when a shared client's lifecycle or face changed, the `local` integration
  suites when they reach the change - `go vet` under every build tag, and
  the five lint stages run from soil's source - golint `--fix`,
  golangci-lint, gofix, goanalyze, goaudit and its `--service`; installed
  binaries can predate the source. Against another repository, golint, gofix
  and goanalyze take `--root`, goaudit takes the root as an argument
- Each loop ends with this seed updated - the census, the order, a lesson
  where one was learned - and consolidated: a lesson merges into the item it
  extends, and what the code or a design doc now holds leaves

## Target shape

| Treatment | The target | Reference |
|---|---|---|
| Ownership | the service builds and closes everything; `Run` holds only the service; the store borrows the database; surfaces hold the service alone | gonetboxd, golinkaced, gomattermostd (optional parts) |
| Honest answers | the client classifies from the response status in its own lowest layer; the service checks writes; the spec inlines a response per class; each handler branches per class | pkg/netbox, pkg/linkace, pkg/chat/mattermost |
| Daemon client | `client.Client` struct; `face.Client` in full, `Close` included; `<path>/client/mock_client` in full; a coverage test against the embedded spec; the CLI on `EmitValue` | gonetboxd, gohabiticad, gomaintlogd |
| Upstream client | errors classified, a transport failure keeping its cause; `Must<Name>` twins only where a caller panics; a full face in `<domain>/face/` and a full mock beside the client; every consumer on them | pkg/kubernetes, pkg/linkace, pkg/chat/mattermost, pkg/prometheus |
| Peers | every peer guarded; the service holds the full face beside its guard and runs each call through `Attempt` in the operation - `guarded[T]` and `attempt` helpers, no adapter types; waited on in `Run` only when a surface has no answer without it | gosentryd, golinkaced (surface peers), goatlasd (worker peers), gomattermostd (both) |
| Worker | one service operation per tick that returns outcomes and an error; the worker keeps the ticker, the gauges and recovery | goatlasd, goalertlogd |
| Outlier transport | one API on two listeners, one client with two constructors | goprocessd |

A daemon is done when it meets every row that applies to it. The service has
no face; a face describes one client in full.

## Census

At the target, with what each is the reference for:

- gonetboxd - all three loops, the worked example; its NetBox face (287
  methods) and stateful mock moved home to `pkg/netbox`
- golinkaced - A landing before B; surface peers
- goatlasd - worker peers, collectors taking a guard
- gosentryd - a surface peer's guard
- gohabiticad - a daemon client and CLI on `EmitValue`
- godirectoryd - an upstream face and mock beside the client
- gokubernetesd - MCP only; one guard per cluster, none waited on
- goraidd - a traversal closed by bare names
- gocertificated - a peer resolved at construction; web answers by class
- goatlassiand - two upstream faces moved home; a consumer moved onto the
  daemon client
- goalertlogd - a worker-only peer; a cursor client without `Must` twins; a
  stateful daemon-client mock proven by a round trip through the server
- goalpined - an owner of nothing; an upload staged and renamed into place
- gomaintlogd - a store daemon; a store that fired events moved to the service
- gomattermostd - optional parts; a full 88-method upstream face; a status
  classifier shared with Atlassian; a WebSocket keeping its own reacher; a
  CLI rule (`getUnresolvedThreads`) moved into the daemon
- goqueryd - a peer never waited on, every surface answering without it; a
  process-wide model shared in tests behind a lease; three consumers moved
  onto one client
- godashboardd - optional peers by configuration; one tick over several
  peers reporting each failure
- gogated - a peer some users need and others do not: a directory outage
  answers its users 503 while local users log in; browser pages answering
  by class at their status (`view.RenderPageWithStatus`), a hand-routed
  JSON endpoint by class with the event identifier; a write check moved
  from two surfaces into the service. No C: it is consumed over OIDC
- gotelemetryd - the service as its own surfaces' recorder: `Mount` takes
  the recorder, production passes the service, tests a mock
- gosproutd - a clock and an in-process notifier as option factories,
  which tests set; its file watcher renamed into `worker`, the one package
  name the service rule allows
- goflightd - local commands as option factories (`Command`, `Stream`
  returning `face.Stream`, its mock beside
  `github.com/funtimecoding/soil/pkg/system/run/stream`), so tests drive
  every tick; four worker packages folded into one context-cancelled worker
  type over a tick, the stream's tick blocking until its command ends and
  killed on cancel; deduplication as service state, remembered only after
  the write succeeds
- gofirefoxd - a peer that dials in: the client is the bridge's WebSocket
  handler, so the face includes `http.Handler` and the service hands it to
  `Run` as `Bridge()`; nothing the daemon does brings the extension back,
  so no guard - its absence answers `unavailable`, unreported. The
  extension's refusals classified by the fixed text `bridge.js` throws;
  tests drive a real client through a fake extension over WebSocket

Beside them, upstream faces now full and home: Prometheus, Alertmanager,
Argo CD (`ArgocdClient` in the kubernetes face) and Nextcloud usage
(`UsageClient`); the Ollama and OpenAI embedder clients classify in their
own `refusal` packages.

Ownership done, other rows open: gochromed, gosecretd, gomonitord,
gogithubd, gooutpostd, goraidparsed, goansibled, gosaltd, goterraformd,
goitermd, gosublimed, gojellyfind, gopnsensed, goprocessd, goalertmanagerd,
gobandd, gogitlabd, gopostgresd, goprometheusd, goproxmoxd, gosourced.

## Order

Sized by the re-survey of 2026-10-26 (`goaudit --service` hits, each read).
Leaks first - a client opened and never closed - then shape.

1. goagentd, a small one, missed by the 2026-10-26 re-survey: no service;
   `Run` builds a runner that starts a `claude --print` process in a
   goroutine nothing can stop, so shutdown leaves it running; the checks
   (intent required, already running) sit in the REST handler; a daemon
   client without a face or mock, consumed by the goagent CLI and goatlasd
2. gomemoryd with goclauded, which consumes its client: `Run` builds the
   store and the goqueryd client, waits on goqueryd and then calls it
   unguarded, and reconciles with a hand-held reacher; `RestClient` without
   a face; finished clients handed to `service.New`. REST and MCP already
   answer by class
3. goclauded - `Run` builds about fifteen things: store, pool, the goqueryd
   client, search index, five tickers, gauge, watcher; finished clients
   handed to `service.New`; the largest of all
4. The faces and clients off the full shape, the peers off the start-up
   shape
5. The goaudit rule, once a sound check replaces `run_builds_client`
6. The analyzer gaps, once the findings they raise are few, each with a
   case in the `unclosed_resource` testdata

## Off the full shape

- Daemon clients: `RestClient` beside a `client.Client` interface, no face -
  gomemoryd. No `<path>/client/` package - most daemons, many rightly; which
  have a consumer reaching them another way is unsurveyed. gokubernetesd
  serves MCP only
- Faces mirroring the service: gosublimed, goitermd, gojellyfind, gochromed,
  gosecretd, gogithubd; goproxmoxd's `face.Service` and `mock_service`
- Faces narrowed to one consumer: gomonitord, goclauded
- Upstream faces or mocks kept in the consuming daemon: gopnsensed's
  `face.OpnsenseSource` and `mock_source`; goproxmoxd's upstream
  `mock_client` colliding with the daemon client's
- Mocks: gomaintlogd's is not held to its face by a test; 62 methods of
  `github.com/funtimecoding/soil/pkg/netbox/mock_client` still answer `nil,
  nil` - each gains state, the `Failure` check and the per-method one
  through `github.com/funtimecoding/soil/pkg/failure` when a test needs it,
  as the virtual disks and sizing setters did
- Pooled daemons hard-wire their client factory in the service -
  goprometheusd, gobandd, gogitlabd, gopostgresd - so their tests cannot
  reach a mock
- Upstream clients: pkg/linkace, pkg/habitica and pkg/errors/sentry/basic map
  401 with 403 to `forbidden`, so an expired token fails quietly; pkg/gitlab
  does not classify its peer's 503 - checked against GitLab's own status
  table first; pkg/chat/mattermost's `New` panics and dials, so
  gomattermostd's start-up wait never retries
- Peers off the start-up shape: goclauded and gomemoryd wait on goqueryd in
  `Run` but call it unguarded; gomemoryd's `Reachable` probes
  `route:/api/versions`, not `route:/api/reachable`
- Workers whose `Close` is not safe twice, against `service-tool.md`
  (Workers): godashboardd's worker closes its stop channel bare
- Web submits showing `e.Error()` through `form.Redirect` without a class
  check, so an unexpected failure reads as a form refusal: godirectoryd's
- Untested wiring: gogated's admin pages - a served test needs a signed-in
  session from the SSO client; `IsAdministrator` itself is tested.
  gosproutd's worker setup sits under recovery, but only exhausting file
  descriptors makes `fsnotify` fail, so no test reaches it
- Checks in surfaces, service operations without errors: gosproutd's MCP
  checks a decision's understanding, looks up a seed by name and bounds a
  position before calling the service, and `Clear`, `SetPosition`,
  `Reorder` and `Seeds` return no error, so a store failure panics inside
  them - its B loop

## Ports

The role interface and its function type are in `service-tool.md`
(Service); a function type goes in a subpackage where the interface's
package already has a receiver type. Decided 2026-10-27 over three
alternatives: functions name no client; a generated guarded decorator over
the full face is bloat - about 30 guarded calls across the consumers, never
more than two per peer, against ~37 and ~32 methods kept in step; the guard
held outside the shared code sees a failure only after it gave up, and
cannot hold calls off.

## Spec contradiction

The class table in `error-handling/README.md` answers `unreachable`,
`timeout` and `dropped` as a captured 503 `ErrorResponse`; goaudit's
`error_response_on_non_500` allows `ErrorResponse` on a 500 only. Every peer
behind a guard answers `unavailable` (503 `Error`) instead, so the captured
503 shrinks with each guard added.

## Craft

What the specs do not say - how the work goes, not what the code must be.

Surveying:

- Each conversion found a defect nobody was looking for - a leak, a wrong
  answer, a write corrupting records (an edit blanking two fields it was
  not sent, a CLI reporting a delete that never happened, writes answering
  success for a record that does not exist). A survey that reads only
  `run.go` misses surfaces calling the client directly, and tests passing
  nil
- A claim put to the user as a decision rests on code read, not on a name,
  a spec's description or a truncated grep; what was not read is said so. A
  peer's status table comes from its source, cloned at the deployed version
  and read, not summarised by a fetch - Prometheus answers 503 for a query
  timeout, so mapping it to `unavailable` opens a guard on a healthy peer;
  one upstream answers a refused token and an ACL refusal both 403, told
  apart only by a field of the body. A closed-source peer is probed live
  with harmless requests: LM Studio answers an unknown model and an
  over-long input 200 (the loaded model, truncated), so a misnamed embed
  model never fails there. A peer's failure as it reaches users comes from
  the deployed daemon's log and the live metrics before a reporting design
  is settled: a fleet of peers dead for weeks, each reported, is a Sentry
  issue per peer on every deploy. A class no source shows is not mapped on
  a guess
- Read the peer client's constructor before deciding where a start-up lookup
  goes: a constructor that already calls the peer has decided it
- A face exposing the client's internals (`Basic()`, `Nested()`) lets a
  surface bypass it; each escape becomes a method on the client, and a
  consumer that used the escape for something the client lacks asks for
  that method. A hand-written client for one of our daemons, kept outside
  it, is that daemon's missing client package
- Count a method's callers by receiver across every module and build tag in
  one pass: a `go/packages` program resolving each selector through
  `TypesInfo.Selections`. gosourced's `find_references` is per symbol and
  per module; a name-level grep only proves a zero
- A CLI that bypassed its daemon to run a rule gets the rule as a daemon
  operation; other callers running the same rule are then bypasses too

Building:

- Closing a leak alone, with a service that only embeds the client, is not a
  conversion. Hold each owned thing in one field; the same guard in two
  fields reads to the analyzer as an owner never closed
- The service never holds its daemon's own worker: `Run` builds the worker
  over the service, and `Workers()` hands out peer guards only - the other
  way round is an import cycle
- An option's defaults reach every test built from it; a test service drops
  what production configures (collections) rather than letting it report. A
  part one mode never uses is optional, nil in that mode, or building it
  demands credentials that mode never needed
- The store saves; the service gets, applies the patch, checks, saves - a
  patch applied inside the store hides the result from the check. A check
  finds the writes already corrupting records: an update setting every
  field from a form that posts some of them sets only what was sent. An
  operation writing before it can refuse leaves the refused input behind;
  stage it and move it into place once every check passed
- A store that fires events owns a side effect; the service fires after a
  write that succeeds
- A move that changes a sentinel to a shared class changes every branch
  testing the sentinel in the same change, or each 404 turns into a 500; the
  class made visible can fail an analyzer the sentinel hid from. An empty
  list is not `not_found` - goanalyze's class table refuses a list surface
  un-mapping it to 200; where surfaces disagree on the empty case (REST 200,
  MCP text), the operation returns the data with a flag. Text dates
  becoming times keep each filter's string semantics: a date-only bound
  includes its day on `from` and excludes it on `to`, which midnight
  reproduces. An aggregate over a time column has no declared type, so
  sqlite answers it as text that will not scan into `time.Time`; read the
  newest row's typed column instead. A store no test reached hides this
- A handler joining URL segments into a path is a traversal until each
  segment is checked as a bare name - an encoded `%2e%2e` passes
  `http.ServeMux`'s cleaning. A root fixed in a constant hides it from tests.
  A value formatted into a CSS selector or a script is the same hole: check
  it as the narrowest shape it can be (a bare number)
- A write whose confirmation timed out against an external party is neither
  success nor refusal: it may have happened. It answers a captured fault
  whose message says so and asks to check before retrying, since a blind
  retry can repeat it; a create that cannot find its new identifier answers
  the same, never success with an empty one. No error class is needed when
  no caller branches on it - the message carries it
- A client that panics on every failure, recovered into plain errors at the
  service, answers every failure as one captured 500. Returning classified
  errors from the client is the conversion; a recover stays only around a
  dependency that still panics, turned into a captured error
- An upstream client's own `Close` can leak (a WebSocket and keep-alive left
  running) and, once fixed, race a goroutine still reconnecting - lock the
  shared state and refuse a refresh after close
- A library that turns a status into its own error type has thrown the
  status away by the time a method sees it; classify inside the transport
  the library calls, keeping the response beside the error where the
  library reads it (`DoGetFallback` retries on the status). go-openapi
  keeps the code on its errors - a `Code()` on declared responses, a `Code`
  field on `APIError` - so its clients classify at each call site instead.
  A status meaning the caller's mistake on some calls only carries its own
  refusal on those requests (`request.WithRefusal`)
- An oapi-codegen client decodes a non-JSON answer by its content type: a
  YAML body declared as a string fails to unmarshal into it, so that
  operation reads the raw body from the plain client
- A full face over a large client is generated work: a program reading the
  client's AST writes the face and trivial mock files as new files, refusing
  to overwrite stateful ones; Must twins copy the client's body; where a face
  exists, only the missing methods go to a scratch file for Edit. Its zero
  values cover the basic types it lists - a struct result needs `{}`, a float
  `0` - and a body returning only the failure must carry no import block:
  golint turns an empty one into `import ()` or a file that no longer
  parses, so grep for `^import ()` after it runs. A mock's existing failure
  field names the generated bodies (`c.failure` where `Fail(e)` sets it)
- Moving a face or mock home is gosourced work - `move_package` then
  `rename_package` for the mock, `move_symbol` then `rename_symbol` for the
  face
- A face gaining `Close` makes every test holding a double of it its owner:
  one helper closing on cleanup (`NewSource(t)`, goqueryd's `newEmbedder(t)`
  and `newReranker(t)`); the service closing the same double again is
  harmless. A part too costly to build per test (goqueryd's ONNX reranker)
  stays owned by the service in production; the test package shares one
  behind a lease whose `Close` is a no-op, and `TestMain` closes it after
  `m.Run()`, the one point every test has finished
- A test seeding a store the service will own seeds it through its own
  connection on the option's `LitePath`, then builds the service there. A
  test seam building the service from parts (`NewFor`) builds one that
  cannot close, and its holder is an owner goanalyze flags; and a fleet
  double above the upstream face cannot reach a service built from the
  option - the fixtures become the face's seeded data, through one
  translator proven first by loading every fixture back field by field.
  Size a test migration by its doubles, not its files
- A seeder gaining fields becomes a second seeder beside it
  (`SeedSizedDeviceType` beside `SeedDeviceType`) where tests call the
  first; a field the upstream does not model travels as the custom field it
  is
- Keep a client's existing signatures where its consumers depend on them; a
  CLI on `EmitValue` loses its column formatting
- Fixtures in soil carry neutral values - NATO names, generic fields -
  never a consumer's sites, hosts or field names

Testing:

- A break proof backs the file up under a name unique in the backup
  directory - two packages' `get_owner.go` overwrite each other - breaks the
  code path, not its inputs, keeping it compiling, watches the test fail,
  restores and compares against the backup. A break that makes the code
  block - a reader never killed - hangs to the package timeout and names no
  test, so a test of a blocking operation runs it in a goroutine and fails
  by name after a bound; a break that panics aborts the package and hides
  every other break's verdict, so breaks run one at a time. An Edit that
  missed its anchor leaves the code unbroken and passes as a proof; the
  break script compares the target with its backup and refuses to run on a
  match, and reads a build failure as no verdict rather than a failed test.
  A break the test survives names a branch no test reaches - add the test -
  or code nothing needs: a service once-guard over parts each safe twice,
  removed. Staging while a break stands stages it; name the files. A mock
  with one `Failure` for every method fails a sequence at its first step, so
  a break swallowing a later step's error survives; the mock fails one
  method by its method value (`c.Fail(c.DeviceTypes, e)`), which a rename
  follows and which refuses a method taken through the face
- A filter test needs a row the filter must drop and one it must keep: a
  filter dropping everything passes a test asserting only the drop. It
  counts the rows directly: gogated's `GetLoginSession` deletes an expired
  row as it reads it, so a cleanup test reading through it passed with the
  cleanup removed
- A quiet class reads the same guarded or not, so an outage test cannot prove
  the guard; what only the guard does - fail the next call at once after an
  `unavailable`, without asking the peer - can. The double's failure is what
  the client produces: a dial failure arrives as a `connection.Failure`,
  which the guard counts as an outage, where a bare `unreachable.Format`
  passes through as a refusal and the guard never opens. A hand-built
  `net.OpError` classifies only around a real errno
  (`os.NewSyscallError("connect", syscall.ECONNREFUSED)`); one wrapping a
  plain error answers 500. A guard's outage report waits out `Lasting`, so a
  fast test counts the peer's calls instead of the reporter's events
- A mock answering `nil, nil` where the client never does hides a nil
  dereference; an unseeded query answers an empty result. A double replaced
  by a shared mock keeps its old default by seeding it: a query double
  answering 100 to every query, replaced by a mock answering 0, flips which
  alerts a threshold filters. A mock that only ever replaced the client
  hides the client's own bugs: gofirefoxd's `CreateTab` decoded into a nil
  pointer and failed every call after Firefox had opened the tab; one round
  trip through the real client found it. A browser-scraping client is tested
  without a browser through a page serving a captured fixture's HTML as its
  body; its `Evaluate` tells a click (nil target) from a value poll, or a
  test meant for one branch fails earlier in another - two write branches
  answering alike each need their own test, which a break of the other
  shows. The daemon-client methods carry the spec's operation names, so the
  coverage test checks names exactly instead of mapping them
- Moving a peer from the guard's `Do` to `Attempt` moves the reporting: the
  caller that got nothing back now gets the error and reports what
  `classify.Reportable` admits, or every non-outage failure goes silent. A
  tick collecting from several peers returns each failure on its own:
  `errors.Join` classifies as its first matching branch, so one outage hides
  a real failure beside it
- A face test's break removes a mock method no other mock file calls; a Must
  twin calling its plain method fails the mock's own build first. A
  round-trip test makes a renamed method a build failure first; the coverage
  test's own break is a spec operation with a server stub and no client
  method. The round trip runs against the daemon's real mount, so a
  consumer's hand-built fake upstream decoding the same replies goes with
  the code it tested; what the consumer decides moves onto the mock
- A failure through a transient status on an idempotent request waits out
  the requester's retries; a write answers at once
- A test reaching an unexported MCP tool goes through the protocol; tests
  reaching a real upstream run only under `//go:build ci`. Wiring in `Run` -
  a middleware passed to the server - is untested, as in every reference

Tools:

- gosourced covers renames, moves and imports, and refuses every operation
  while the module does not type-check; it loads a consumer with soil, so
  keep the consumer type-checking between steps of a cross-repository
  change. It renames struct methods, not interface methods - Edit the face
  first - and moves a type with methods only through `extract_type`,
  refusing back-references
- Edit changes existing files, one anchor per site, and an anchor must be
  unique: in an OpenAPI spec the components reference the same schemas as
  the operations, so anchor on the operation's own request schema. A
  `replace_all` of a name rewrites it inside string literals too. A program
  may write new files only. Rewriting YAML through a library re-indents
  every list
- gofix removes import aliases where nothing collides and renames `err*`
  helpers; a file named after the old identity then fails goanalyze. golint
  sorts imports and can leave a half-merged `var` block for gofix to settle
- goanalyze keeps `<path>/constant/` packages free of behaviour and every
  package to one receiver type: value checks go in the service over constant
  slices, a patch's apply in the service as a function. A constant whose
  value is a common word collides under `string_constant`, and a test value
  equal to any constant's value collides too. goaudit allows
  `<path>/constant/` only directly under `pkg/<domain>` or
  `pkg/tool/<name>`, so a test package's values
  stay literals

## Analyzer gaps

The analyzer misses four shapes, so a clean lint is not a clean tree:

- A consuming chain on a fresh owner - `gitlab.NewEnvironment().MustProjects()`
  is never assigned, so never a candidate. Hits: eight soil examples,
  Ollama's `reachable_skip` twice
- An owner captured by a closure - `escapes.go` treats any function literal
  referencing the owner as a hand-off; no hit remains
- An owner built inline as another constructor's argument -
  `store.New(lite.New(...))` - is never assigned; goclauded's `Run` and the
  telegram environment are live hits
- An owner built in `Main` and handed to `Run` crosses a function boundary
  the analyzer does not follow; no hit remains

## goaudit rule, parked

A disarmed rule sits in
`github.com/funtimecoding/soil/pkg/tool/goaudit/scan/check_service_shape.go`,
armed by `--service` or `service: true` in a repository's goaudit
configuration.

| Check | Verdict |
|---|---|
| `missing_service` | sound |
| `service_not_closed` | sound |
| `run_builds_client` | most hits real on the 2026-10-26 count; false positives are a daemon's own workers (`watch.New` in gosecretd and kind) and plumbing (`notifier.New`, `reacher.New`); it knows a worker by its package being named `worker`, so a second ticker takes the same type over a tick function (goflightd) rather than a package of its own |

It does not see a client built inside the `Mount` closure or a worker
`Close` unsafe twice. The candidate replacement is
`run_owns_resource` - `Run` defers `Close` on anything but the service -
which reads what `Run` owns instead of what it calls; it would not see a
client never closed at all, so it complements `run_builds_client` rather
than replacing it.
