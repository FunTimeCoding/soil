# Daemon ownership

What is still open in bringing every daemon to the shape the specs set:
one owner per daemon (`service-tool.md`, The daemon, Services, Stores,
Workers; `lifecycle.md`), answers by class (`error-handling/`), guarded
peers (`error-handling/infrastructure.md`, The guard), one daemon client
with a full face and mock (`generated-api.md`, The Daemon Client), and the
tests that hold them (`testing.md`). How the work is done is
`../runbook/refactor.md`. A consuming repository keeps its own daemons'
open work in a seed of its own. An item leaves this seed when it is done.

## Order

1. A consumer in another repository wraps goclauded's generated client for
   its own short timeout: the typed client builds its pool only through an
   unexported constructor, so it needs a timeout or builder constructor
   first
2. The bubbletea monitor builds gomonitord's client itself in `New`, so its
   mock cannot reach the model; a factory would let a test drive claiming
   and releasing without a daemon
3. Outage tables (`testing.md`, Peer Guard Tests) still owed by goatlassiand,
   gomattermostd, goclauded and goatlasd, each guarding more than one peer, and
   by the pooled goalertmanagerd, gobandd, gogitlabd, gopostgresd,
   goprometheusd, goproxmoxd, gokubernetesd and goterraformd. gobackrestd,
   godashboardd, gogithubd and gopatrold have no REST or MCP surface to tab.
   A provisioning runner never started by a table waits forever on `sync`, so
   gosaltd's and goansibled's tables close it first
4. Request contexts into the guarded services
   (`error-handling/infrastructure.md`, The guard): calls a caller starts
   still pass `context.Background()` to `go:pkg/peer.Guarded` or
   `Guard.Attempt` at each call, or inside the pooled helper, in
   goatlassiand, gobandd, gochromed, godirectoryd,
   gogated, gohabiticad, goitermd, gojellyfind, golinkaced, gomattermostd,
   gonetboxd, gopnsensed, goproxmoxd, gosaltd, gosecretd, gosentryd and
   gosublimed, and in single calls - gobandd's intent record, gohabiticad's
   `Equip`, gomemoryd's profile reads, gosentryd's `DeleteIssue`, gosourced's
   memory retarget and goclauded's document pushes from `complete` and
   `summarize`. The work no caller starts keeps its own: godashboardd's
   refresh, gogithubd's worker, goterraformd's run targets, gomattermostd's
   monitor, notifier and watcher, gosaltd's runner, goalertlogd's recording
   tick and goclauded's reconcile and sweep
5. The daemon as owner in goanalyze: `unclosed_resource` was taught on a
   service owning what it built, so a test seam holding a daemon built from
   finished clients is not flagged. Nor does it see a release the
   constructor registered itself: a tester holding goclauded's test server
   and client, each closed by `t.Cleanup` inside the constructor that built
   it, is flagged unless the tester registers the release again. How the
   analyzer learns which type owns is not worked out; a case goes into the
   `unclosed_resource` testdata when it is
6. `goaudit --daemon` lists no daemon here, but no repository's
   `goaudit.yaml` sets `daemon: true`, so the check runs only when asked
7. goaudit's `error_response_on_non_500` allows 503 too, so its name no
   longer says what it checks; `missing_capture_fail` is met by
   `answer_message.go` too, yet its text still names `capture_fail.go`
8. Selection rides on the MCP session of the legacy protocol versions.
   Protocol version 2026-07-28 drops protocol-level sessions, so moving to
   it needs another home for the choice

## Off the full shape

- One-row lookups onto `go:pkg/relational.FindOne` and `FetchOne`
  (`database.md`, Store Shape): 23 gorm `First` calls outside tests, none
  `Take` - gogated 16, goclauded 3, and one each in godashboardd, goflightd,
  gomaintlogd and gosproutd; mint is at zero. A goanalyze
  restricted call on `First` and `Take` holds the rule once this repository and
  its consumers reach zero
- golinkaced's mock still answers stubs; it takes gosublimed's shape and a
  shared-answers test. `pkg/linkace`'s `FromDaemon` converters and `Format`
  methods have no caller in this module; they stay as exports until a
  consumer outside it is ruled out
- 59 methods of
  `github.com/funtimecoding/soil/pkg/netbox/mock_client` answer `nil, nil` -
  each gains state, the `Failure` check and the per-method one through
  `github.com/funtimecoding/soil/pkg/failure` when a test needs it. Most of
  the Proxmox mock's methods ignore `SetFailure`, and 36 method files of
  `github.com/funtimecoding/soil/pkg/gitlab/mock_client` ignore `Fail` - each
  gains it when a test needs it
- REST and MCP still decide an expected error with no class differently: MCP's
  `go:pkg/generative/mark/response.AnswerMessage` honours
  `go:pkg/errors/classify.Reportable`, while REST's
  `go:pkg/errors/classify.Capture` skips only a caller-left error. No other
  expected error reaches REST today
- Upstream clients: client-go answers GitLab's 404 with a sentinel, so
  pkg/gitlab's not-found carries "Not Found" rather than GitLab's message;
  pkg/gitlab retries a 429, a 5xx and a refused dial five times with backoff -
  about 12 s for a refused dial - before the error reaches a caller (gogitlabd
  builds with one retry, its guard owning anything longer);
  pkg/chat/mattermost's `New` panics and dials, so gomattermostd's start-up wait
  never retries, and its face takes no context, so a caller's context stops at
  the guard - no hang-up or trace reaches a post; pkg/gitlab's `Jobs` deletes
  old pipelines from inside a read when `AutoCleanup` is set
- A downstream daemon's 500 is captured twice, by the daemon and again by
  the caller reading it as `unexpected` - every daemon-to-daemon client,
  most often on clocks over a peer (goclauded's `Reconcile` against goqueryd,
  `PollMemory` and `DeliverImpressions` against gomemoryd). How many
  captures one 500 makes there is unmeasured
- `github.com/funtimecoding/soil/pkg/ticker`'s `Close` panics before `Start`
  and on a second call, and does not wait for a tick in flight; nothing in
  this repository builds on it
- `github.com/funtimecoding/soil/pkg/worker`'s `NewEvery` wraps a
  `func() []error`, so `Close` cannot cancel a tick in flight the way `New`'s
  can. `NewEvery` retires into `New` daemon by daemon, each operation taking
  `x` and passing it to the peer it reaches - not wrapped at the wiring. What
  is left in a repository: `grep -rn 'worker.NewEvery(' --include=*.go pkg`
- A Mattermost that accepts the socket and stalls the handshake holds a
  hold's `Close` up to gorilla's 45 s `HandshakeTimeout`; gomattermostd's
  watcher looks up an event's author over the client inside the socket's
  callback
- godirectoryd's web submits show `e.Error()` through `go:pkg/web/form.Redirect`
  without a class check, so an unexpected failure reads as a form refusal
- Surfaces still deciding on an error: goproxmoxd's MCP `captureDetail` tests
  `proxmox.ErrNotAuthorized` and `proxmox.ErrTimeout`, and gochromed's
  classifies transport, where the client should (`external-api.md`, Translate at
  the boundary); gobandd's REST `captureDetail` replaces every message but an
  `unexpected` one with a generic text, where the message is the error's own
  (`README.md`, Answering an error). The `captureDetail` of goalertmanagerd,
  gobandd's MCP, gokubernetesd, gonetboxd's MCP and REST, gopostgresd,
  goprometheusd and gosourced only forwards to the answer helper - a name left
  over
- gosourced's MCP handlers answer through its own `captureFail`, which answers
  by class and decides one case itself: a `load_failure` is answered quietly
  while the broken package is fresh and captured once nothing in it changed for
  `go:pkg/tool/gosourced/constant.LoadSettled`. The class table has no
  load-failure class, and a class's capture is fixed, so "quiet while edited
  toward, captured once settled" cannot be said there yet - that is what folding
  `captureFail` into `answer` waits on
- gocertificated's `policy.Validate` checks a caller's distinguished name with
  `PanicOnEmpty`, so a request leaving out a required field such as the country
  is a captured panic rather than a `validation` answer
- gosproutd's MCP checks a decision's understanding, looks up a seed by name
  and bounds a position before calling the service, and `Clear`,
  `SetPosition`, `Reorder` and `Seeds` return no error, so a store failure
  panics inside them - its B loop
- goansibled's `unit` package builds its fixture with `git init`, `clone` and
  `push` (`ansible_tester.NewFixture`), which crosses the line
  `test-placement.md` draws for `<path>/unit/`

## Untested

- goterraformd's downstream guards and goansibled's inventory guard: nothing
  proves the probe loop brings a peer back after a hold-off. Both build
  their guards at the default backoff, so such a test waits out
  `ProbeMinimum`
- Leak tests that cannot see a client left open: gosaltd's master client
  (the master never answers, so no connection is held), and the SSO client
  in godirectoryd, goraidd and gocertificated (it starts nothing until used)
- Clock wiring: gogitlabd's `Refresh` and `Entries`, gomemoryd's
  `Reconcile` (the goqueryd mock's records are read unlocked while a tick
  runs)
- gopostgresd: a refusal surfacing while rows are read (needs a live
  database). gomemoryd: its profile search branch has no failure test of its
  own, and `go:pkg/tool/gomemoryd/importer.Import` has no caller but its test
- goproxmoxd: two handlers are tested per class, and `error_answer` checks
  the shape of the other branches but not that each is present. PVE's
  refusal texts are fixtures, so only a live call would show an upgrade
  rewording them; 595, 596 and 599 stay captured until a second node exists
- gogated's admin pages (a served test needs a signed-in SSO session);
  gosproutd's worker setup under recovery (only exhausting file descriptors
  makes `fsnotify` fail); MCP session end clearing the selection in the
  daemons using `WithSessionEnd` (the hook is tested in
  `github.com/funtimecoding/soil/pkg/generative`)
- goaudit does not see a client built inside the `Mount` closure, or a
  worker `Close` unsafe twice
