package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gocredentiald",
	"Credential store daemon over a KeePass database",
	"gocredentiald",
).WithInstructions(
	"Credential store over a KeePass database. Use list_entries and search_entries to browse, get_entry for details (passwords masked), reveal_password only when a value is genuinely needed, audit for staleness findings, and the create/update/move/delete verbs for maintenance. load_group returns an environment variable group by name.",
)

const (
	RevealedFieldEnvironment = "CREDENTIAL_REVEALED_FIELD"

	ListEntries    = "list_entries"
	GetEntry       = "get_entry"
	RevealPassword = "reveal_password"
	SearchEntries  = "search_entries"
	AuditEntries   = "audit_entries"
	CreateEntry    = "create_entry"
	UpdateEntry    = "update_entry"
	MoveEntry      = "move_entry"
	DeleteEntry    = "delete_entry"
	LoadGroup      = "load_group"

	EnvironmentGroup = "Environment"
	MaskedValue      = "•••"

	TitleKey    = "Title"
	PasswordKey = "Password"
	LocatorKey  = "URL"
	NotesKey    = "Notes"

	DefaultStaleYears = 3
	AuditSampleLimit  = 10
	AuditPageLimit    = 100

	BucketStale         = "stale"
	BucketEmptyUser     = "empty_user"
	BucketEmptyPassword = "empty_password"
	BucketDuplicates    = "duplicates"
)
