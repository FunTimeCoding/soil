package constant

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"time"
)

const CoverageRecentWindow = 30 * 24 * time.Hour
const (
	PermissionWindowDays     = 7
	PermissionSampleCommands = 3
	PermissionCommandLength  = 400
	PermissionReasonPrefix   = "bashMiss:"
)
const CallsignReleaseWindow = 7 * 24 * time.Hour
const AbandonedQueueWindow = 7 * 24 * time.Hour
const SevenDayWindow = 7 * 24 * time.Hour
const UsagePageUnreadable = "usage page carries no weekly meter"
const (
	CompleteTimeoutWindow   = 30 * time.Minute
	InactivityTimeoutWindow = time.Hour
)
const ChannelHold = 5 * time.Minute
const ReconcileInterval = 10 * time.Minute
const (
	TimeoutSweepInterval    = 5 * time.Minute
	MemoryPollInterval      = 30 * time.Second
	ImpressionInterval      = 30 * time.Second
	FindingInterval         = 5 * time.Minute
	UsagePollInterval       = time.Minute
	UtilizationPollInterval = 5 * time.Minute
)
const RosterStaleAfter = 45 * time.Minute
const DaemonStarted = "started"
const (
	PageSize        = 20
	HistoryPageSize = 10
	Page            = "page"
	Active          = "active"

	EventKind       = "history event"
	ActivityPath    = "/activity"
	ActiveLabel     = "active"
	DetailHeadClass = "detail-head"
	IdleLabel       = "idle"

	EditAction   = "edit"
	PulseAction  = "pulse"
	CruiseAction = "cruise"
	DeleteAction = "delete"
)

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
	MemoryPeer            = "memory"
	QueryPeer             = "query"
	UtilizationPeer       = "utilization"
	MemoryLoadsHidden     = "%v; loads hidden"
	MemoryUnreachableKind = "memory_unreachable"
	ImpressionContext     = "impression"
	ImpressionSource      = "source"
	ImpressionContent     = "content"
	ImpressionCapturedAt  = "captured_at"
)

const (
	SessionExportDirectoryEnvironment = "SESSION_EXPORT_DIRECTORY"
	MonitorUsageEnvironment           = "CLAUDE_MONITOR_USAGE"
	UsageEndpointEnvironment          = "CLAUDE_USAGE_ENDPOINT"
	CoverageRootEnvironment           = "CLAUDE_COVERAGE_ROOT"
	ChannelAlwaysEnvironment          = "CLAUDE_CHANNEL_ALWAYS"
)

const (
	DeliveryDescription              = "immediate (default) wakes each named recipient now; one without a channel gets it with their next prompt. queued arrives with their next prompt or after their idle grace - for a note that can wait."
	DeliveryDescriptionChannelAlways = "immediate (default) wakes each named recipient now. queued arrives with their next prompt or after their idle grace - for a note that can wait."
)

const (
	SessionName = "name"
	SessionKind = "session"
	Callsign    = "callsign"
	Topic       = "topic"
	Files       = "files"
	To          = "to"
	Body        = "body"
	Immediate   = "immediate"
)

const (
	Announce      = "announce"
	Complete      = "complete"
	Update        = "update"
	EditEvent     = "edit_event"
	SessionStatus = "session_status"
	Roster        = "roster"
	ListSessions  = "list_sessions"
	History       = "history"
	HistoryCount  = "history_count"
	EditSession   = "edit_session"
	Send          = "send"
	Register      = "register"
	Status        = "status"
	SessionEnd    = "session_end"
	Release       = "release"
	Summarize     = "summarize"
	Moment        = "moment"
	TokenUsage    = "token_usage"
	Description   = "description"
)

const (
	Usage             = "usage"
	Activity          = "activity"
	InactivityTimeout = "inactivity_timeout"
	CompleteTimeout   = "complete_timeout"
)

const (
	EventSummary = "summary"
	EventSession = "session"

	SummaryCollection    = "summaries"
	CompletionCollection = "completions"
	SummarySourceType    = "session-summary"
	CompletionSourceType = "session-completion"

	CompletedAtMetadata = "completed_at"

	Collection    = "collection"
	Outcome       = "outcome"
	OutcomePushed = "pushed"
	OutcomeFailed = "failed"
	OutcomeSwept  = "swept"
)

const (
	SessionTable         = "session"
	SummaryTable         = "summary"
	LabelTable           = "label"
	PulseTable           = "pulse"
	MessageTable         = "message"
	RateSnapshotTable    = "rate_snapshot"
	FableSnapshotTable   = "fable_snapshot"
	SummaryColumn        = "summary"
	TokenUsageColumn     = "token_usage"
	ModelColumn          = "model"
	ContextPercentColumn = "context_percent"
	ContextWindowColumn  = "context_window"
	CompletionTable      = "completion"
	CruiseSettingTable   = "cruise_setting"
	CruiseModeColumn     = "mode"
	CompletedAtColumn    = "completed_at"
	GraceColumn          = "grace"
	ReleasedColumn       = "released"
)

const (
	Pulse          = "pulse"
	Label          = "label"
	Key            = "key"
	Value          = "value"
	Past           = "past"
	Now            = "now"
	UnsetMarker    = "(unset)"
	LabelNotSet    = "%s (not set)"
	LabelUnchanged = "%s %s (unchanged)"
	LabelAdd       = "add"
	LabelRemove    = "remove"
	LabelHeld      = "label %s already holds a value, nothing written"
	LabelNotHeld   = "label %s does not hold %q, nothing written"
	ChangeArrow    = "→"
	LegacyChange   = "change"
)

const (
	Subscription          = "subscription"
	EventIdentifierColumn = "event_identifier"
	LastSeenColumn        = "last_seen"

	Subscriber      = "subscriber"
	Kinds           = "kinds"
	After           = "after"
	EventStreamName = "entry"
	EventStreamPath = "/event-stream"
	StreamBatch     = 200

	StreamBuffer           = 64
	StreamLineBuffer       = 1048576
	StreamIdentifierPrefix = "id: "
	StreamPayloadPrefix    = "data: "
	StreamOperation        = "stream"
	StreamEnded            = "stream ended"
	Target                 = "target"
	Alias                  = "alias"
	Slug                   = "slug"
	Limit                  = "limit"
	Offset                 = "offset"
	ListPage               = 500
	Since                  = "since"
	Before                 = "before"
	Kind                   = "kind"
	Full                   = "full"
	Line                   = "line"
	Message                = "message"
	Identifier             = "identifier"
)

const (
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

	ImmediateLimit  = 6
	ImmediateWindow = 5 * time.Minute
)

const (
	BroadcastLabel = "broadcast"
	BroadcastOff   = "off"
)

const (
	Compaction           = "compaction"
	Phase                = "phase"
	Trigger              = "trigger"
	CompactionNotice     = "%s %s compacting"
	CompactionTrigger    = " (%s)"
	CompactionTopic      = " - %s"
	CompactionTopicLimit = 300
	CompactionUnknown    = "unknown compaction phase %q"
	CompactionWatchLabel = "role"
	CompactionWatchValue = "lead"
)

const (
	AutocompleteOff = "off"

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
	PermissionsTitle   = "Permissions"
	PermissionsPath    = "/permissions"
	UsageTitle         = "Usage"
	UsagePath          = "/usage"
	UsageChart         = "usage_chart"
	StatusTitle        = "Status"
	StatusPath         = "/status"
)

const (
	ModelContextServersFile = ".mcp.json"

	FixtureTopic  = "topic"
	FixtureTarget = "target"
	FixtureBefore = "before"

	FixtureLabelKey   = "role"
	FixtureCompatible = "lead"

	FixtureUnknownDelivery = "eventually"
)

const FixtureChannelHold = 50 * time.Millisecond

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
