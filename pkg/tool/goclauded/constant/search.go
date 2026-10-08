package constant

const SearchIndexFile = "search.sqlite"
const SearchIndexVersion = 2
const (
	BlockMessage = "message"
	BlockEdit    = "edit"
	BlockCall    = "call"
)

const (
	SearchConversationLimit   = 20
	SearchConversationMaximum = 100
	SearchSnippetLimit        = 3
	SnippetRadius             = 60
	BlockTextLimit            = 2000
	ReadWindowDefault         = 5
	ReadWindowMaximum         = 25
)

const (
	SearchConversations = "search_conversations"
	ReadConversation    = "read_conversation"
)

const (
	Session = "session"
	Around  = "around"
	Count   = "count"
)

const (
	QueryField = "query"
	KindField  = "kind"
)

const UnnamedSession = "unnamed"

var EditTools = []string{"Edit", "Write", "MultiEdit", "NotebookEdit"}
