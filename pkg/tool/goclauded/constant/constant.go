package constant

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/identity"
)

var (
	ErrorAliasCollision   = errors.New("alias already in use")
	ErrorCallsignNotFound = errors.New("callsign does not exist")
	ErrorEventNotFound    = errors.New("event not found")

	ReservedLabelKeys = []string{
		SessionName,
		Alias,
		Description,
		Slug,
		Topic,
		Files,
	}
	Identity              = identity.New(
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

	SessionName = "name"
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

	SessionTable  = "session"
	SummaryTable  = "summary"
	SummaryColumn = "summary"

	Pulse      = "pulse"
	Label      = "label"
	Key        = "key"
	Value      = "value"
	Target     = "target"
	Alias      = "alias"
	Slug       = "slug"
	Limit      = "limit"
	Offset     = "offset"
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

	FixtureTopic  = "topic"
	FixtureTarget = "target"
	FixtureBefore = "before"
)
