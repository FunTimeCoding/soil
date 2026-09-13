package constant

import "time"

const (
	SwapTitle    = "Swap"
	TriggerTitle = "Trigger"
	TriggerPath  = "/trigger"
	NotifyTitle  = "Notify"
	NotifyPath   = "/notify"
	StreamTitle  = "Stream"
	StreamPath   = "/stream"
	ExtraTitle   = "Extra"
	ExtraPath    = "/extra"
)

const (
	RemovePath      = "/remove"
	AppendPath      = "/append"
	QuietPath       = "/quiet"
	BranchPath      = "/branch"
	PollPath        = "/poll"
	RedirectPath    = "/redirect"
	RefreshPath     = "/refresh"
	FirePath        = "/fire"
	SlowPath        = "/slow"
	IndicatorMark   = "indicator-slot"
	SlowMark        = "slow-result"
	SlowControlMark = "slow-post"
	IndicatorText   = "working"
	RemoveMark      = "remove"
	AppendMark      = "append"
	QuietMark       = "quiet"
	BranchMark      = "branch"
	PollMark        = "poll"
	FireMark        = "fire"
	FiredEvent      = "fired"
	LoadQuery       = "?term=load"
	DelayedQuery    = "?term=delayed"
	TriggerHeader   = "HX-Trigger"
	RefreshHeader   = "HX-Refresh"
	PollWhen        = "every 500ms"
	FragmentContent = "fragment"
	PageContent     = "full page"
)

const (
	CounterPath    = "/counter"
	RowPath        = "/row"
	FailPath       = "/fail"
	OutOfBandPath  = "/out-of-band"
	SubscribePath  = "/subscribe"
	CounterMark    = "counter"
	RowMark        = "row"
	EchoMark       = "echo"
	ReceiptMark    = "receipt"
	TickMark       = "tick"
	PulseMark      = "pulse"
	SummaryMark    = "oob-summary"
	LoadMark       = "load-region"
	DelayedMark    = "delayed"
	SelectionMark  = "selection"
	SearchMark     = "search"
	TermField      = "term"
	ScopeField     = "scope"
	ConfirmMessage = "Replace this row?"
	FailMessage    = "deliberate failure for the notification path"
)

const (
	MechanicToken         = "mechanic-example-token"
	ExtendedEnvironment   = "MECHANIC_EXTENDED"
	ServerSideEnvironment = "MECHANIC_SERVER_SIDE"
)

const DelayWhen = "load delay:600ms"
const SlowWait = 400 * time.Millisecond
