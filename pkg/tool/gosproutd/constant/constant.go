package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gosproutd",
	"Local seed priority tracker",
	"gosproutd",
).WithInstructions(
	"Local seed priority tracker. Scans a seed directory for markdown files, tracks them with priority ordering. Use list_seeds to see priorities, reorder_seed to change them.",
)

const (
	SeedDirectoryEnvironment = "SEED_DIRECTORY"

	ListSeeds   = "list_seeds"
	ReorderSeed = "reorder_seed"

	DashboardTitle = "Dashboard"
	DashboardPath  = "/"
	SeedsTitle     = "Seeds"

	Seeds       = "seeds"
	SeedsRecent = "seeds-recent"

	SortParameter = "sort"
	SortModified  = "modified"

	MaximumTurnLength = 300
	MaximumTurnCount  = 3

	PushDecision   = "push_decision"
	DecisionStatus = "decision_status"
	AddTurn        = "add_turn"
	ClearDecision  = "clear_decision"
	BumpDecision   = "bump_decision"

	SessionParameter       = "session"
	QuestionParameter      = "question"
	DefaultActionParameter = "default_action"
	OptionParameter        = "options"
	FrameParameter         = "frames"
	IdentifierParameter    = "identifier"
	ContentParameter       = "content"
	UnderstandingParameter = "understanding"
	HeardParameter         = "heard"
	ResolutionParameter    = "resolution"

	IdentifierColumn = "identifier"

	AnswerField = "answer"

	SessionsTitle = "Sessions"
	SessionsPath  = "/sessions"
	SessionPath   = "/session"
	DecisionPath  = "/decision"
	AnswerPath    = "/answer"
	DismissPath   = "/dismiss"
	ReplyPath     = "/reply"

	CountersTitle = "Counters"
	CountersPath  = "/counters"

	SessionsEvent = "sessions"
	DecisionEvent = "decisions"
	CounterEvent  = "counters"

	NameParameter  = "name"
	OpenParameter  = "open"
	ValueParameter = "value"
	StateParameter = "state"
	KindParameter  = "kind"

	StateColumn = "state"

	DecisionKind       = "decision"
	DecisionMarkPrefix = "decision-"
	ReorderPath        = "/reorder"
	ItemParameter      = "item"

	SeedTableClass   = "seed-table"
	SeedPathClass    = "seed-path"
	SeedOrderMark    = "seed-order"
	DragHandleClass  = "drag-handle"
	CounterNoteClass = "counter-note"

	StatusPart  = "status"
	ChoicePart  = "choices"
	DismissPart = "dismiss"
	ThreadPart  = "thread"
)
