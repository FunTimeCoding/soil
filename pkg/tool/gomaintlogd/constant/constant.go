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
	AddEntryPath   = "/add"
)
