package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gomaintlogd",
	"Operational maintenance log service",
	"gomaintlogd",
).WithInstructions(
	"Operational maintenance log - record, list, update, and delete maintenance entries. Track what was done, to which system, and why.",
)

const (
	AddEntry    = "add_entry"
	ListEntries = "list_entries"
	UpdateEntry = "update_entry"
	DeleteEntry = "delete_entry"
	Action      = "action"
	Description = "description"
	Identifier  = "id"
	Service     = "service"
	Since       = "since"
	System      = "system"
	Timestamp   = "timestamp"
	Until       = "until"
	User        = "user"

	DashboardTitle = "Dashboard"
	DashboardPath  = "/"
	EntriesTitle   = "Entries"
	EntriesPath    = "/entries"
	AddEntryTitle  = "Add Entry"
	EntryTitle     = "Entry"
	EntryPath      = "/entry"
	DetailPath     = "/detail"
	EditPath       = "/edit"
	DeletePath     = "/delete"
	AddEntryPath   = "/add"

	HostEnvironment     = "GOMAINTLOG_HOST"
	PortEnvironment     = "GOMAINTLOG_PORT"
	InsecureEnvironment = "GOMAINTLOG_INSECURE"
	TokenEnvironment    = "GOMAINTLOG_TOKEN" // #nosec G101

	EventSummary = "summary"
	EventRecent  = "recent"
	RecentMark   = "recent-table"
	EntriesMark  = "entries-table"
	AddFormMark  = "add-form"
	DetailPrefix = "detail-"
	RowPrefix    = "row-"
	FilterFormat = "2006-01-02T15:04"

	ClickableRowClass  = "clickable-row"
	DetailRowClass     = "detail-row"
	DetailContentClass = "detail-content"
	DetailActionsClass = "detail-actions"
)
