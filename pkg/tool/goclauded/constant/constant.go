package constant

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"time"
)

const CoverageRecentWindow = 30 * 24 * time.Hour

var (
	ReservedLabelKeys = []string{
		SessionName,
		Alias,
		Description,
		Slug,
		Topic,
		Files,
	}
	Identity = identity.New(
		"goclauded",
		"Session coordination for parallel Claude Code sessions",
		"goclauded",
	).WithInstructions(
		"Session coordination for parallel Claude Code sessions. Read the goclauded://guide/session-workflow resource on your first turn to understand the session lifecycle, tool rhythm, and coordination patterns.",
	)
)

const (
	SessionExportPathEnvironment = "SESSION_EXPORT_PATH"
	MonitorUsageEnvironment      = "CLAUDE_MONITOR_USAGE"
	CoverageRootEnvironment      = "CLAUDE_COVERAGE_ROOT"

	SessionName = "name"
	SessionKind = "session"
	Callsign    = "callsign"
	Topic       = "topic"
	Files       = "files"
	To          = "to"
	Body        = "body"

	Announce     = "announce"
	Complete     = "complete"
	Update       = "update"
	EditEvent    = "edit_event"
	Status       = "status"
	Roster       = "roster"
	ListSessions = "list_sessions"
	History      = "history"
	HistoryCount = "history_count"
	EditSession  = "edit_session"
	Send         = "send"
	Register     = "register"
	Release      = "release"
	Listen       = "listen"
	Summarize    = "summarize"
	Moment       = "moment"
	TokenUsage   = "token_usage"
	Description  = "description"

	Usage             = "usage"
	Activity          = "activity"
	InactivityTimeout = "inactivity_timeout"
	CompleteTimeout   = "complete_timeout"

	EventSummary = "summary"
	EventSession = "session"

	SummaryCollection    = "summaries"
	CompletionCollection = "completions"
	SummarySourceType    = "session-summary"
	CompletionSourceType = "session-completion"

	SessionTable         = "session"
	SummaryTable         = "summary"
	RateSnapshotTable    = "rate_snapshots"
	FableSnapshotTable   = "fable_snapshots"
	SummaryColumn        = "summary"
	TokenUsageColumn     = "token_usage"
	ModelColumn          = "model"
	ContextPercentColumn = "context_percent"
	ContextWindowColumn  = "context_window"

	Pulse      = "pulse"
	Label      = "label"
	Key        = "key"
	Value      = "value"
	Target     = "target"
	Alias      = "alias"
	Slug       = "slug"
	Limit      = "limit"
	Offset     = "offset"
	ListPage   = 500
	Since      = "since"
	Before     = "before"
	Kind       = "kind"
	Full       = "full"
	Line       = "line"
	Message    = "message"
	Identifier = "identifier"

	QueueSessionAnnounce = "session_announce"
	QueueSessionRelease  = "session_release"
	QueueSessionComplete = "session_complete"
	QueueSessionUpdate   = "session_update"
	QueueMessage         = "message"
	QueueNotification    = "notification"
	QueuePulse           = "pulse"
	QueueMemoryUpdate    = "memory_update"
	QueueMemoryCreate    = "memory_create"
	QueueTimeout         = "timeout"
	QueueReannounce      = "reannounce"

	DashboardTitle     = "Dashboard"
	DashboardPath      = "/"
	SessionsTitle      = "Sessions"
	SessionsPath       = "/sessions"
	MessagesTitle      = "Messages"
	MessagesPath       = "/messages"
	HistoryTitle       = "History"
	HistoryPath        = "/history"
	ConversationsTitle = "Conversations"
	ConversationsPath  = "/conversations"
	CoverageTitle      = "Coverage"
	CoveragePath       = "/coverage"

	ModelContextServersFile = ".mcp.json"

	FixtureTopic  = "topic"
	FixtureTarget = "target"
	FixtureBefore = "before"
)

type Action int

const (
	ActionSkip Action = iota
	ActionCopy
	ActionUpdate
)

var DefaultNames = []string{
	"Ash",
	"Blair",
	"Cedar",
	"Dale",
	"Ellis",
	"Frost",
	"Glen",
	"Harbor",
	"Jade",
	"Kent",
}

var MultiWordPrefixes = map[string]bool{
	"go":       true,
	"git":      true,
	"task":     true,
	"kubectl":  true,
	"logcli":   true,
	"goclaude": true,
	"podman":   true,
	"docker":   true,
}
