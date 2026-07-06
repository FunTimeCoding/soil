package constant

import "github.com/funtimecoding/go-library/pkg/identity"

var Identity = identity.New(
	"gomemoryd",
	"Persistent memory across Claude Code sessions",
	"gomemoryd",
).WithInstructions(
	"Persistent memory across Claude Code sessions. Call profile on your first turn to load memories. Read the gomemoryd://guide/memory-workflow resource for the memory lifecycle, tags, search, and profile tiers.",
)

const (
	HostEnvironment = "MEMORY_HOST"
	PortEnvironment = "MEMORY_PORT"

	SaveMemory     = "save_memory"
	UpdateMemory   = "update_memory"
	Profile        = "profile"
	ListMemories   = "list_memories"
	GetMemory      = "get_memory"
	ForgetMemory   = "forget_memory"
	SearchMemories = "search_memories"
	RelateMemories = "relate_memories"
	SetParent      = "set_parent"
	TagMemory      = "tag_memory"
	ListTags       = "list_tags"

	MemoryName       = "name"
	Source           = "source"
	Content          = "content"
	Description      = "description"
	Tag              = "tag"
	Type             = "type"
	MemoryIdentifier = "memory_id"
	SourceIdentifier = "source_id"
	TargetIdentifier = "target_id"
	ParentIdentifier = "parent_id"
	IncludeHistory   = "include_history"
	Topic            = "topic"
	AlwaysTag        = "always"
	NoIndexTag       = "no-index"
	Add              = "add"
	Remove           = "remove"
	ReplaceAll       = "replace_all"

	DashboardTitle   = "Dashboard"
	DashboardPath    = "/"
	MemoriesTitle    = "Memories"
	MemoriesPath     = "/memories"
	ImpressionsTitle = "Impressions"
	ImpressionsPath  = "/impressions"
	SearchTitle      = "Search"
	SearchPath       = "/search"
	Identifier       = "identifier"
	Query            = "query"

	ProfileBudget = 15000

	FixtureName    = "name"
	FixtureContent = "content"
)
