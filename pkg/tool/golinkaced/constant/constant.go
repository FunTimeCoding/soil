package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"golinkaced",
	"LinkAce bookmark archive with REST and MCP",
	"golinkaced",
).WithInstructions(
	"LinkAce bookmark archive - list links, lists, tags, search, and manage bookmarks.",
)

const (
	ListLinks = "list_links"
	Search    = "search"
	AddLink   = "add_link"
	EditLink  = "edit_link"
	Delete    = "delete"
	ListLists = "list_lists"
	ListTags  = "list_tags"
	ListNotes = "list_notes"
	AddNote   = "add_note"

	ListParameter = "list"
)
