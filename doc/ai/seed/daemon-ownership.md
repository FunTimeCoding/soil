# Daemon ownership

Who owns and closes what in a daemon, how it answers, and how other code
reaches it. The rules are settled and live in the specs: one
`daemon.Daemon` per daemon, built from configuration, the only owner `Run`
holds, its use cases in services split by noun, its stores split along
consistency boundaries (`service-tool.md`, The daemon, Services, Stores,
Workers; `lifecycle.md`; `http-client.md`, Owning;
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

The daemons listed at the target in Census meet every row of it; goclauded's
three loops are in, its `Run`-built workers left (Order).

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
  - A, ownership - the daemon builds and closes everything; every surface
    goes through the services it is handed, holding no client of its own
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
  golangci-lint, gofix, goanalyze, goaudit and its `--daemon`; installed
  binaries can predate the source. Against another repository, golint, gofix
  and goanalyze take `--root`, goaudit takes the root as an argument
- Each loop ends with this seed updated - the census, the order, a lesson
  where one was learned - and consolidated: a lesson merges into the item it
  extends, and what the code or a design doc now holds leaves

## Target shape

| Treatment | The target | Reference |
|---|---|---|
| Ownership | the daemon builds and closes everything; `Run` holds only the daemon; each store borrows the database; surfaces hold the services they call alone; the daemon keeps no forwarding method | gonetboxd, golinkaced, gomattermostd (optional clients) |
| Honest answers | the client classifies from the response status in its own lowest layer; the service checks writes; the spec inlines a response per class; each handler branches per class | pkg/netbox, pkg/linkace, pkg/chat/mattermost |
| Daemon client | `client.Client` struct; `face.Client` in full, `Close` included; `<path>/client/mock_client` in full; a coverage test against the embedded spec; the CLI on `EmitValue` | gonetboxd, gohabiticad, gomaintlogd |
| Upstream client | errors classified, a transport failure keeping its cause; `Must<Name>` twins only where a caller panics; a full face in `<domain>/face/` and a full mock beside the client; every consumer on them | pkg/kubernetes, pkg/linkace, pkg/chat/mattermost, pkg/prometheus |
| Peers | every peer guarded; the service calling the peer holds the full face beside its guard, both built by the daemon, and runs each call through `Attempt` in the operation - `guarded[T]` and `attempt` helpers, no adapter types; waited on through the daemon's `Wait` only when a surface has no answer without it | gosentryd, golinkaced (surface peers), goatlasd (worker peers), gomattermostd (both) |
| Worker | one service operation per tick that returns outcomes and an error; the daemon builds the clock and hands it out through `Workers()`, after its guards; the worker keeps the ticker, the gauges and recovery | goatlasd, goalertlogd, `github.com/funtimecoding/soil/pkg/worker` |
| Outlier transport | one API on two listeners, one client with two constructors | goprocessd |

A daemon is done when it meets every row that applies to it. Neither the
daemon nor a service has a face; a face describes one client in full. The
layout's rollout moves each owner to `<path>/daemon/` and splits the large
ones by noun; `goaudit --daemon` lists the daemons still on
`<path>/service/`.

## Census

At the target, with what each is the reference for:

- gonetboxd - all three loops, the worked example; its NetBox face (287
  methods) and stateful mock moved home to `pkg/netbox`
- golinkaced - A landing before B; surface peers
- goatlasd - worker peers, collectors taking a guard
- gosentryd - a surface peer's guard
- gohabiticad - a daemon client and CLI on `EmitValue`
- godirectoryd - an upstream face and mock beside the client; on
  `<path>/daemon/`, the daemon holds the guard and the SSO client and hands the
  client to web; its REST and MCP answer tests, once over hand-wired services,
  go served, and the form test seeds through its own form. The leak test cannot
  see the SSO client left open
- gokubernetesd - MCP only; one guard per cluster, none waited on; on
  `<path>/daemon/`, the port forwards a registry the daemon builds and hands
  out after the guards, launching each forward and stopping every one it still
  holds at close, in `<path>/port_forward/registry/` beside the forward it
  holds; the cluster, resource, port forward and format helpers at the
  root, the result and request types in `<path>/types/`
- goraidd - a traversal closed by bare names; on `<path>/daemon/`, the watcher
  built over the service's event handler, its watched paths and its start-up
  steps, which run in order before watching, each returning its errors so one
  that fails leaves the next running; the daily cleanup a clock of its own at
  05:00; the report result type in `<path>/types/report/`. The leak test cannot
  see the SSO client left open, which starts nothing until used
- gocertificated - a peer resolved at construction; web answers by class; on
  `<path>/daemon/`, the godashboardd recipe: the daemon holds the SSO client and
  hands it to web, its forge guard now started through `Workers()`, and the
  publish tests are service tests built from parts. The leak test cannot see the
  SSO client left open, which starts nothing until used
- goatlassiand - two upstream faces moved home; a consumer moved onto the
  daemon client; on `<path>/daemon/`, the state the web read off a worker - the
  plate and its live notifier - is the service's, refreshed by a plain clock
  that notifies only on change, and every surface test is served
- goalertlogd - a worker-only peer; a cursor client without `Must` twins; a
  stateful daemon-client mock proven by a round trip through the server; on
  `<path>/daemon/`, the last poll and the live notifier the surfaces read off a
  worker are the service's - `Record` stamps the attempt and notifies, failed or
  not - and the worker is a gauged clock over that one function value
- goalpined - an owner of nothing; an upload staged and renamed into place;
  the first on `<path>/daemon/` - a daemon with no `Close`, the service built
  from its parts, served tests mounting the daemon built from the option
- gomaintlogd - a store daemon; a store that fired events moved to the service.
  On `<path>/daemon/`: the daemon opens the database and hands the service and
  the web one notifier, which the web's live stream reads - a served test
  hears a write arrive on it
- gomattermostd - optional clients; a full 88-method upstream face; a status
  classifier shared with Atlassian; a WebSocket keeping its own reacher; a
  CLI rule (`getUnresolvedThreads`) moved into the daemon. On
  `<path>/daemon/`, the optional parts reach the service through `With…`, and
  `Close` flushes what the watcher drained before the goclauded client closes.
  That flush relies on a closed `peer.Guard` still admitting `Attempt`: the
  lifecycle closes the workers, the roster guard among them, before the
  daemon's `Close` runs. A guard that refused once closed would drop the
  drained events, so the flush moves before anyone changes that
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
  name the service rule allows; on `<path>/daemon/`, the watcher is built
  over the service's sync as a function value, and its `Close` waits for its
  goroutine - which the leak test cannot prove, since goleak retries until a
  late goroutine has exited, so a test holds a sync in flight instead
- goflightd - local commands as option factories (`Command`, `Stream`
  returning `face.Stream`, its mock beside
  `github.com/funtimecoding/soil/pkg/system/run/stream`), so tests drive
  every tick; four worker packages folded into one context-cancelled worker
  type over a tick - since lifted into the shared
  `github.com/funtimecoding/soil/pkg/worker`, whose `Close` waits for a
  running tick and whose once form runs start-up work - the stream's tick
  blocking until its command ends and
  killed on cancel; deduplication as service state, remembered only after
  the write succeeds. On `<path>/daemon/`, the daemon builds the clocks, the
  wireless one only when wdutil answers, and a served test breaks the store by
  dropping a table through its own connection
- gofirefoxd - a peer that dials in: the client is the bridge's WebSocket
  handler, so the face includes `http.Handler`. On `<path>/daemon/`, one mount
  method per listener: `MountBridge` serves the client on the bridge port beside
  `Mount`, and the `Bridge()` accessor is gone. The leak test connects a fake
  extension before closing, since nothing runs until one dials in. Nothing the
  daemon does brings the extension back,
  so no guard - its absence answers `unavailable`, unreported. The
  extension's refusals classified by the fixed text `bridge.js` throws;
  tests drive a real client through a fake extension over WebSocket
- goagentd - a process the daemon owns: the runner is the daemon's one worker,
  handed to the service, and a `face.Drainer`, so shutdown waits out the grace
  period and then kills the agent; it launches through an option factory
  returning `face.Stream` with the stderr writer passed in, so tests drive every
  branch on `mock_stream` without shelling out. An interrupt arriving before the
  launch returns is remembered and kills on attach - the tests calling
  `Interrupt` right after `Run` reach that window. Its own transport classifier
  went; `pkg/errors/connection` does that job for every client
- gomemoryd - a store that is the truth and a peer that is its projection:
  writes save and answer success whatever goqueryd does, a worker reconciles
  every memory, forgotten ones included, through the guard and once at start,
  and goqueryd's hash dedup and idempotent delete keep that cheap; profile
  leaves goqueryd's tiers out with a notice rather than failing. A daemon
  client moved three consumers off the generated client, one of them a CLI
  in another repository. On `<path>/daemon/`, the daemon holds the hidden tag
  and hands it to the service and to REST, whose profile leaves hidden
  memories out of every tier - a served test proves it through the daemon
  client; the reconcile clock is built in `New` and handed out after the
  guard; the profile's text shapes in `<path>/format/`. Untested: the daemon
  wiring `Reconcile` into the clock, since the goqueryd mock's records are
  read unlocked while a tick runs
- goprometheusd - a pooled peer guarded per instance: the client factory on
  the option, each pool entry a `peered` pair of client and guard built and
  closed together, the pool filled in `New` so `Workers()` hands every guard
  to the lifecycle; operations run through a `guarded` helper with the
  instance's name. Its upstream already classified, so no B work. The pooled
  shape on `<path>/daemon/`: the daemon holds the pool, the inventory and the
  selection, `Mount` hands the last two to the MCP surface, `peered` sits at
  the daemon's root; the service tester builds the pool of mock-and-guard
  pairs itself, and the daemon tester builds the daemon from the option
- goalertmanagerd - two peers behind one instance, each guarded: the pool
  entry holds the Alertmanager and Prometheus clients and a guard for each,
  so a Prometheus outage no longer reads as Alertmanager's. The Alertmanager
  client stopped holding a Prometheus client: it had fetched the rules
  through one built inline as its constructor's argument and never closed,
  and cached them for the process's life, so a new rule never reached
  enrichment and a silence for it was refused until a restart. Rules are now
  an argument to `Alerts` and `SetSilence`, fetched by each caller from a
  Prometheus client it owns - 13 call sites in 7 packages, counted by
  receiver, moved new-beside-old through a transitional variadic so the
  consumers kept building. goalertlogd dropped Prometheus instead of
  guarding it: rules only set a runbook it never stored. The client's nil
  filter and nil option went too - an empty filter sends the same query, but
  an empty option is not no option: `All: false` drops suppressed and
  unrated alerts, so a caller that meant "everything" passes `All: true`.
  On `<path>/daemon/` in goprometheusd's pooled shape; its two service
  testers - mocks, and a real Alertmanager client over a fake server - share
  one pool builder
- gopostgresd - a pooled peer where who wrote the SQL decides whose move a
  refusal is: the client factory on the option over `pkg/postgres`, a guard
  per instance, explain and the default schema in the service, catalogue
  names as query arguments. `pkg/postgres/refusal` classifies a server answer
  by SQLSTATE class from errcodes.txt at the deployed version, transport left
  to `connection.Classify`; the caller's own query answers their mistake
  quietly, a refusal of our catalogue SQL is captured, a statement timeout is
  quiet on theirs and captured on ours, and a caller who left is not captured
  on MCP. REST, MCP and the daemon client answer columns and positional rows,
  and the CLI on the daemon client prints through `terminal.EmitValue` - the
  wire shape changed, so the daemon and the CLI deploy together. On
  `<path>/daemon/` in goprometheusd's pooled shape, `Mount` handing the
  inventory to the REST server as well. `caller_left`, an error class, moved
  to `pkg/errors/caller_left` when a second daemon needed it
- gogitlabd - a pooled peer with a board: the client factory on the option, a
  peered client and guard per instance filled in `New`, every operation
  through `guarded` or `attempt`. The service owns the board and its
  notifier; `Refresh` keeps a failing instance's rows and returns one failure
  per instance, which the worker's tick reports by class, so an instance that
  is away never hides another's refusal. `pkg/gitlab` no longer dials in
  `New`, keeps its caches under a mutex and classifies through the shared
  `refusal.Status`, 412 as `conflict` beside it and 401 captured as the
  daemon's own credential. MCP answers by class through `answerWith`; a
  malformed number on a web page is a validation refusal, rendered 500 until
  `view.Recovery` answers by class. MCP and web only, so no daemon client: the
  gogitlab CLI reads GitLab directly and runs no rule gogitlabd runs. On
  `<path>/daemon/` in goprometheusd's pooled shape, plus a gauged clock: the
  worker takes `Refresh` and `Entries` as function values - one entries
  function feeding both the gauge and the faster interval while a pipeline
  runs - is built once in `New` over the metric registry, since its gauge
  registers on construction, and is handed out after the guards; its `Close`
  is safe twice and waits out a tick in flight, so no `Refresh` meets a closed
  pool. The service keeps the inventory as a part, `Refresh` walking
  it; the daemon holds the web server and hands `Run` its recovery. Untested:
  the daemon wiring `Refresh` and `Entries` into the clock
- gobandd - a pooled peer whose requests the daemon writes: on `<path>/daemon/`,
  the engine factory on the option answering `validation` for an instance
  without a host or password, where `band.New` would exit the process; a
  `peered` engine and guard per host, every engine operation through `guarded`;
  NetBox optional, set on the service by `WithNetbox`, its write after the power
  change outside the engine's guard. The upstream face is `BandClient` in
  `pkg/band/face`, `Reachable` an authenticated settings read so a dead
  credential reads as an outage, its mock beside the client. pkg/band
  classifies: transport through `connection.Classify`, a status without a fault
  through the shared table, the 401 after the digest round captured. DSP0226
  puts sender faults on 400 and receiver faults on 500, but the daemon writes
  every request, so a sender fault is ours, not the caller's: it stays captured,
  and only subcodes that name the class map - `wsman:AccessDenied` to
  `forbidden`, `wsman:AlreadyExists` and `wsman:Concurrency` to `conflict`,
  `wsman:QuotaLimit`, `wsa:EndpointUnavailable` and `wsman:TimedOut` to
  `unavailable`, by local name since the prefix varies. REST answers 400, 403,
  409 and 503 by class
- gocredentiald - a local file the daemon keeps answering through; on
  `<path>/daemon/`, the service built from its parts, its mirror face and the
  mock of the service retired, and the surface tests served. A change on disk it
  cannot read had panicked mid-request; the copy loaded before it keeps
  answering reads, every call retries the reload, a write is refused as
  `conflict`, and Sentry hears once per change. `pkg/keepass` returns its
  errors: `New` beside a `MustNew` for start-up, a failed save no longer leaves
  the protected entries locked in memory or a temporary file behind, and a
  backup that fails refuses the write. A failed write is discarded by reloading
  from disk. It owns nothing running, so it has no `Close` and no leak test

Beside them, upstream faces now full and home: Prometheus, Alertmanager, Argo CD
(`ArgocdClient` in the kubernetes face), Nextcloud usage (`UsageClient`),
GitHub, GitLab (`GitLabClient`), aptly (`AptlyClient` beside the `Repository`
role), Proxmox (`ProxmoxClient`, its access and cluster reads joining the
machine, container, node and snapshot roles) and SSH (`SSHClient`, the snippet
role goproxmoxd had held), Docker Hub (`DockerHubClient`, its refusals classed
at the client - an outage or rate limit `unavailable`, a request of ours that
Hub rejects left unexpected); the Ollama and OpenAI embedder clients classify in
their own `refusal` packages. GitHub's and GitLab's faces and trivial mocks were
generated from the client's AST, after a by-receiver count across both
repositories and every tag - one counting calls through the face and inside the
package too, or a twin only the face or the package itself calls reads as dead -
removed the Must twins nobody called (8 and 31; 17 in pkg/proxmox).
`pkg/postgres` is the client for the Postgres instances gopostgresd inspects,
not a second way to open our stores - those stay on `pkg/relational`. Its query
returns ordered columns and positional rows, so a duplicate column name and the
column order survive.

Ownership done, other rows open: gochromed, gosecretd, gomonitord, gogithubd,
gooutpostd, goraidparsed, goansibled, gosaltd, goterraformd, goitermd,
gosublimed, gojellyfind, gopnsensed, goprocessd, goproxmoxd, gosourced. Of
these, gopnsensed, gomonitord, gosublimed, goitermd, gojellyfind and gochromed
are on `<path>/daemon/`, the last five with their mirror face retired - each
surface takes the service. gogithubd is too: its clock had held a two-method
face the service filled, `Owner` among them, and closed its stop channel bare;
the daemon builds it on `github.com/funtimecoding/soil/pkg/worker` over `Latest`
and the option's owner, so its `Close` is safe twice. goraidparsed is too: its
REST had held a one-method face the service filled, and its tests a fixed answer
in the service's place; the Python top-stats parser is now the factory's
`TopStatsParser` and the service holds it, so linking, counting, finding and
moving the report are tested served over a parser that writes what Python would.
It owns nothing, so it has no `Close`. A report's name carries the minute it was
made, so two runs on one day no longer overwrite each other. goproxmoxd is too:
both pools come from factories on the option - `BuildHypervisor` answering
`validation` for an instance without a host or credentials, where `proxmox.New`
would panic, and `BuildShell` dialling SSH lazily per instance - its mirror face
and `mock_service` retired, and the snippet operations its surfaces ran on an
SSH client of their own are the service's. A snippet name joined into the
snippet directory was a traversal; it is checked as a bare name. The SSH mock
listed every file under its full path, wherever it lay, so it now lists one
directory by base name as the client does. The poller reaches the API client
through the pool's lookup until it becomes a clock over a service `Refresh`; the
API pool has no guard yet, and its surfaces answer the new classes through the
old branches (next: P and the per-class spec). pkg/proxmox classifies in a
transport under go-proxmox, which keeps only `res.Status` for a 500 and folds
401 and 403 into `ErrNotAuthorized`. PVE (pve-manager `HTTPServer.pm`
`rest_handler`) answers a `PVE::Exception`'s code and 500 for any other `die`,
its message in the status line, and never 404 for a guest: "does not exist"
(pve-guest-common `AbstractConfig.pm` 58, and the snapshot lookups) reads as
`not_found`, "is locked" (340) as `conflict`, 403 (`raise_perm_exc`) as
`forbidden` still matching `ErrNotAuthorized`, so a token without `Sys.Audit`
keeps its node. A 400 is `validation` only on create, update and clone, which
carry the caller's configuration; elsewhere it is ours, captured. 401 passes to
go-proxmox's session re-create. The texts are fixtures, so only a live call
would show a PVE upgrade rewording them. 595, 596 and 599 - AnyEvent::HTTP's
statuses when PVE cannot reach another cluster node - stay captured until a
second node exists. pkg/ssh's file operations return errors
classified at their lowest layer - a missing path `not_found`, a refused one
`forbidden`, sftp's lost connection `dropped` - tested against an in-process SSH
server; its lazily opened sftp session had no lock while REST and MCP shared one
pooled client, a race now closed. Its shell pool is the only pool in the fleet
holding a connection between calls - the others hold HTTP clients that redial
per request, or pgx's self-healing pool - so only goproxmoxd calls
`github.com/funtimecoding/soil/pkg/inventory/pool`'s `Forget`: a snippet
operation answering `dropped` forgets that shell, the error still reaching the
caller and nothing retried, and the next call redials. `Forget` drops a client
only while the pool still holds that one, so a second caller holding the dead
client never closes one a third call rebuilt; the pool's client type is
`comparable` for it, and its `Set` went with goproxmoxd's seam. gooutpostd is
too: it owns nothing, so it has no `Close` and no leak test, and the service
takes its reader. A failed `systemctl` or `launchctl` listing had answered 200
with no services, which goatlasd read as a host with nothing deployed; it is an
error, captured and answered 500, and the lookups after it still tolerate a
nonzero exit. Its client is named for the operations, `ListServices` and
`GetHost`, and a coverage test holds it to the spec. goprocessd is too: its
supervisor is the daemon's one worker - `Start` spawns the Procfile, `Close`
stops every process - and `Run` waits on interrupt, terminate, hangup or the
daemon's `Done`, so a closed terminal stops the processes and the daemon still
ends when they have all exited. Its second listener, the local socket, is
`MountLocal` beside `Mount`. Each process's log writer had run a goroutine
nothing stopped, a dropped process's included; the writer now flushes a partial
line on a timer and runs no goroutine. Its leak test starts real processes, so
it lives under `<path>/integration/`. gosourced is too: the daemon builds and
closes its memory peer and guard; the index cache and the turn scheduler start
nothing and close nothing, so they stay the service's own fields. Its
subpackages moved out of `<path>/service/` - value types to `<path>/types/`,
the sink, turns, decoration, matcher, report and revision to the root. Its
250-file service splits by noun in a loop of its own.

goterraformd's downstream guards and goansibled's inventory guard have no
recovery test. Tests prove that `Workers()` returns them first and that an
outage is held off and kept out of Sentry, but nothing proves the probe loop
brings a peer back after a hold-off. Both daemons build their guards at the
default backoff, so such a test would wait out `ProbeMinimum`.

gosaltd's leak test cannot see the master client's `Close`. Its master never
answers, so the client never holds a connection to leak, and the daemon has no
surface that would tell a test when a login against a live master happened.

## Order

Sized by a re-survey (each goaudit hit read).
Leaks first - a client opened and never closed - then shape.

1. goclauded - A and B are in: the service builds the database (the store
   borrowing it), the search index's connection and both peer clients from
   the option and closes them; both peers guarded, neither waited on; writes
   that push to goqueryd save and answer success, a worker reconciling
   summaries and completions every ten minutes; the start-up warm and search
   catch-up run as once-workers on `github.com/funtimecoding/soil/pkg/worker`,
   the catch-up stopping between files on cancel; a stopper worker, registered
   last, releases held channel polls before the servers stop. C's first
   change is in: `connector` renamed `client`, one method per operation
   under its own name with a context and classified refusals (the label and
   pulse helpers and `Subscribe` beside them, the label claim's refusal read
   as `conflict`), `face.Client` and a stateful mock beside the client, a
   coverage test, every daemon consumer on the face; `NewOptional` answers
   the face, so an absent goclauded is a nil interface rather than a nil
   pointer inside one. The operation `deleteSessionById` became
   `deleteSessionByIdentifier`. C's second change is in: the goclaude CLI
   holds two clients, `New` and `NewLong` for the channel's held polls, and
   answers a refusal with `Exitln`, its hooks staying silent. The first
   change's generated methods took the 200 body as the only success, so the
   two operations that also answer 204 for "nothing yet" - `getUsage` and
   `getChannelCallsign` - failed on it, and the callsign wait gave up after
   three released holds; they return the Find shape now. `getResolve`'s 409
   carries the ambiguous matches, which the client returns beside the result
   instead of refusing. A generated method checks every 2xx its operation
   declares, and every non-2xx body the caller reads. Left:
   - the tickers, gauge and watcher built in `Run`, which move to the daemon
     at its layout loop (`run_builds_clock`, `run_builds_client`)
   - `Usage` answers a store failure as no data (204), the swallow B was
     meant to end
2. The faces and clients off the full shape
3. The analyzer gaps, once the findings they raise are few, each with a
   case in the `unclosed_resource` testdata

## Off the full shape

- Daemon clients: no `<path>/client/` package - most daemons, many rightly;
  which have a consumer reaching them another way is unsurveyed. gokubernetesd
  serves MCP only
- Faces mirroring the service: gosecretd
- `github.com/funtimecoding/soil/pkg/tool/gochromed/face.ChromiumSource` has
  no `Close`, so its owners close it through `face.Closer` - gochromed's own
  service and its consumers
- Upstream faces or mocks kept in the consuming daemon: gopnsensed's
  `face.OpnsenseSource` and `mock_source`
- Mocks: gomaintlogd's is not held to its face by a test; 62 methods of
  `github.com/funtimecoding/soil/pkg/netbox/mock_client` still answer `nil,
  nil` - each gains state, the `Failure` check and the per-method one
  through `github.com/funtimecoding/soil/pkg/failure` when a test needs it,
  as the virtual disks and sizing setters did
- gopostgresd: a refusal that surfaces while rows are read is untested (needs
  a live database)
- REST and MCP disagree on an expected error with no class: MCP's
  `response.AnswerWith` honours `classify.Reportable`, REST's
  `classify.Capture` captures it anyway. Only gopostgresd's caller-left error
  meets it on REST, and Sentry has seen one "context canceled" ever (an embed
  cancelled at shutdown), so REST keeps capturing until a second surface
  needs otherwise; the fix then is in `Capture`, not in a daemon
- Upstream clients: pkg/linkace, pkg/habitica and pkg/errors/sentry/basic map
  401 with 403 to `forbidden`, so an expired token fails quietly; client-go
  answers GitLab's 404 with a sentinel, so pkg/gitlab's not-found carries "Not
  Found" rather than GitLab's message, and it retries a 429, a 5xx and a refused
  dial five times with backoff - about 12 s for a refused dial - before the
  error reaches a caller; gogitlabd builds with one retry, its guard owning
  anything longer, and the CLIs keep five. pkg/gitlab retries a 5xx only for
  an idempotent method, so a write meeting one answers at once instead of
  running twice - gocommit's commit, gocertificated's publish, gogitlabd's
  pipeline and branch creates; pkg/chat/mattermost's `New` panics and
  dials, so gomattermostd's start-up wait never retries; pkg/gitlab's `Jobs`
  deletes old pipelines from inside a read when `AutoCleanup` is set
- No caller context: pkg/gitlab keeps `context.Background()` and no `Forge`
  method takes a context, so a caller who left keeps its GitLab call
  running, and gogitlabd's guards attempt under `context.Background()`;
  threading one through its 120 methods is a loop of its own, not yet
  scheduled
- A downstream daemon's 500 is captured twice, by the daemon and again by
  the caller reading it as `unexpected` - fleet-wide, every
  daemon-to-daemon client
- Untested or unused: gomemoryd's profile search branch has no failure test
  of its own - goqueryd's mock fails every method at once, so completions
  fail first; gomemoryd's `importer.Import` has no caller but its test
- Workers whose `Close` is not safe twice, against `service-tool.md`
  (Workers): `github.com/funtimecoding/soil/pkg/ticker`'s `Close` panics
  before `Start` and on a second call, and does not wait for a tick in
  flight; nothing in this repository builds on it any more
- A Mattermost that accepts the socket and stalls the handshake holds a
  hold's `Close` up to gorilla's 45 s `HandshakeTimeout`; gomattermostd's
  watcher still looks up an event's author over the client inside the
  socket's callback
- Web submits showing `e.Error()` through `form.Redirect` without a class
  check, so an unexpected failure reads as a form refusal: godirectoryd's
- Untested wiring: gogated's admin pages - a served test needs a signed-in
  session from the SSO client; `IsAdministrator` itself is tested. gosproutd's
  worker setup sits under recovery, but only exhausting file descriptors makes
  `fsnotify` fail, so no test reaches it. MCP session end clearing the
  selection - `WithSessionEnd` in goalertmanagerd, gobandd, gogitlabd,
  gokubernetesd, gopostgresd, goprometheusd, goproxmoxd and gosourced: the hook
  is tested in `github.com/funtimecoding/soil/pkg/generative`, and nothing
  outside the daemon sees a selection cleared
- Checks in surfaces, service operations without errors: gosproutd's MCP
  checks a decision's understanding, looks up a seed by name and bounds a
  position before calling the service, and `Clear`, `SetPosition`,
  `Reorder` and `Seeds` return no error, so a store failure panics inside
  them - its B loop

## Ports

The role interface and its function type are in `service-tool.md` (The daemon);
a function type goes in a subpackage where the interface's package already has a
receiver type. Decided over three alternatives: functions name no client; a
generated guarded decorator over the full face is bloat - about 30 guarded calls
across the consumers, never more than two per peer, against ~37 and ~32 methods
kept in step; the guard held outside the shared code sees a failure only after
it gave up, and cannot hold calls off.

## Spec contradiction

The class table in `error-handling/README.md` answers `unreachable`,
`timeout` and `dropped` as a captured 503 `ErrorResponse`; goaudit's
`error_response_on_non_500` allows `ErrorResponse` on a 500 only. Every peer
behind a guard answers `unavailable` (503 `Error`) instead, so the captured
503 shrinks with each guard added.

## Spec gaps

`service-tool.md`'s daemon takes the option, a logger, a reporter and a metric
registry. A daemon whose web records domain telemetry builds that web in `New`
for its recovery, so its `New` takes the recorder too and its `Mount` takes
only the mux; the spec's list and its `Mount` signature have not caught up.

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
  per module; a name-level grep only proves a zero. A method gaining an
  error return still compiles wherever a caller discarded its result -
  errcheck finds those, the build does not
- A CLI that bypassed its daemon to run a rule gets the rule as a daemon
  operation; other callers running the same rule are then bypasses too

Building:

- Closing a leak alone, with a service that only embeds the client, is not a
  conversion. Hold each owned thing in one field; the same guard in two
  fields reads to the analyzer as an owner never closed
- Who builds a worker follows what it is, not who calls whom: most workers
  call both ways - a runner a service launches into reports back to it.
  The daemon builds every worker: a clock over one service operation per
  tick, and a worker holding state or a connection, which reaches back
  through function values it was given, since holding the daemon or a
  service that holds it is what makes the import cycle; neither carries
  the work (`service-tool.md`, Services and Workers)
- An option's defaults reach every test built from it; a test service drops
  what production configures (collections) rather than letting it report. A
  client one mode never uses is optional, nil in that mode, or building it
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
- A step a service operation runs that can panic - a file walk over a
  directory that may be missing - is checked before it runs, inside a guard
  or not. `Attempt` releases a trial whose call panics and lets the panic
  go on, so the guard keeps admitting trials, the only way out of an
  outage the peer served itself; but an MCP tool has no recovery of its own
  (`../spec/error-handling/mcp.md`), so the panic answers an HTML 500 the
  model cannot read, and REST a captured 500
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
  harmless. A client too costly to build per test (goqueryd's ONNX reranker)
  stays owned by the daemon in production; the test package shares one
  behind a lease whose `Close` is a no-op, and `TestMain` closes it after
  `m.Run()`, the one point every test has finished
- A test seeding a store the daemon will own seeds it through its own
  connection on the option's `LitePath`, then builds the daemon there. A
  test seam building the daemon from finished clients (`NewFor`) builds one
  that cannot close, and its holder is an owner goanalyze flags - a service
  test builds the service from its parts instead, owning what it built; and
  a fleet double above the upstream face cannot reach a daemon built from
  the option - the fixtures become the face's seeded data, through one
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
  or code nothing needs: an owner's once-guard over workers each safe twice,
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
- A not-found test through a daemon client cannot fail on a missing route:
  the route's own 404 reads as `not_found` too. The double records what it
  was asked and the test asserts the ask - gopnsensed's unknown-host test
  passed with the REST mount removed until it did
- A quiet class reads the same guarded or not, so an outage test cannot prove
  the guard; what only the guard does - fail the next call at once after an
  `unavailable`, without asking the peer - can. The double's failure is what the
  client produces: a dial failure arrives as a `connection.Failure`, which the
  guard counts as an outage, where a bare `unreachable.Format` passes through as
  a refusal and the guard never opens. A hand-built `net.OpError` classifies
  only around a real errno
  (`os.NewSyscallError("connect", syscall.ECONNREFUSED)`); one wrapping a plain
  error answers 500. A guard's outage report waits out `Lasting`, so a fast test
  counts the peer's calls instead of the reporter's events. The guard answers
  `unavailable` on the very call that opens the outage, not only on the calls
  after it, so a tick returning failures through `Attempt` hands its worker
  nothing reportable. A guard comes back only through its probe loop, on a real
  timer from `ProbeMinimum` that no option reaches, so a test that needs the
  peer back - nothing lost across a failed tick - fails it with a fault the
  guard passes through untouched, and leaves the outage to the next-call test
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
  `store.New(lite.New(...))` - is never assigned; the telegram environment is
  a live hit
- An owner built in `Main` and handed to `Run` crosses a function boundary
  the analyzer does not follow; no hit remains

It also sees one borrow as ownership: `checkedValue` follows `v, e := call()`
only from a map index or a type assertion, so a helper that unpacks a pooled
entry - `p, e := s.pool.Client(i)` then `return p.Client(), nil` - is read as
handing out a fresh client, and every caller is flagged unclosed. gogitlabd
calls `p.Client()` inline at each site instead.

Its owner tracking predates the layout: it was taught on a service that
owned what it built. Under the layout no service holds a closer and the
obligation sits on `daemon.Daemon`, so the craft's test-seam line - the
holder of a daemon that cannot close is an owner goanalyze flags -
depends on the analyzer knowing which type owns. How it learns that is
not worked out; a follow-up, with a case in the `unclosed_resource`
testdata when it is.

## goaudit's layout rules

goaudit holds a daemon to `doc/ai/spec/service-tool.md` from the day it
has `<path>/daemon/`, on every run: `missing_service`, `mixed_service`
(a subpackage declaring `type Service` beside code in
`<path>/service/`), `service_imports_mapper`, `run_builds_clock`,
`run_builds_client`, `daemon_not_closed`, and `missing_service_mount`
looking in `<path>/daemon/`; `surface_holds_daemon` runs everywhere.
`--daemon` or `daemon: true` arms only `missing_daemon`, which lists the
daemons not yet converted. goanalyze's `forwarding_function.CheckDaemon`
flags a daemon method that only forwards to the same method on a
service. The checks do not see a client built inside the `Mount`
closure or a worker `Close` unsafe twice.
