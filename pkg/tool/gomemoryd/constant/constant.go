package constant

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/identity"
)

var Identity = identity.New(
	"gomemoryd",
	"Persistent memory across Claude Code sessions",
	"gomemoryd",
).WithInstructions(
	"Persistent memory across Claude Code sessions. Call profile on your first turn to load memories. Read the gomemoryd://guide/memory-workflow resource for the memory lifecycle, tags, search, and profile tiers.",
)

const (
	HostEnvironment      = "MEMORY_HOST"
	PortEnvironment      = "MEMORY_PORT"
	HiddenTagEnvironment = "MEMORY_HIDDEN_TAG"

	SaveMemory       = "save_memory"
	UpdateMemory     = "update_memory"
	Profile          = "profile"
	ListMemories     = "list_memories"
	GetMemory        = "get_memory"
	GetMemoryGroup   = "get_memory_group"
	ForgetMemory     = "forget_memory"
	SearchMemories   = "search_memories"
	RelateMemories   = "relate_memories"
	UnrelateMemories = "unrelate_memories"
	SetParent        = "set_parent"
	TagMemory        = "tag_memory"
	ListTags         = "list_tags"

	MemoryName        = "name"
	Source            = "source"
	Content           = "content"
	Description       = "description"
	Tag               = "tag"
	Type              = "type"
	MemoryIdentifier  = "memory_id"
	MemoryIdentifiers = "memory_ids"
	SourceIdentifier  = "source_id"
	TargetIdentifier  = "target_id"
	ParentIdentifier  = "parent_id"
	IncludeHistory    = "include_history"
	Detail            = "detail"
	Topic             = "topic"
	Scope             = "scope"
	AllScope          = "all"
	DefaultScope      = "default"
	AlwaysTag         = "always"
	NoIndexTag        = "no-index"
	Add               = "add"
	Remove            = "remove"
	ReplaceAll        = "replace_all"

	DashboardTitle   = "Dashboard"
	DashboardPath    = "/"
	MemoriesTitle    = "Memories"
	MemoriesPath     = "/memories"
	ImpressionsTitle = "Impressions"
	ImpressionsPath  = "/impressions"
	RelationsTitle   = "Relations"
	RelationsPath    = "/relations"
	UntypedFilter    = "untyped"
	SearchTitle      = "Search"
	SearchPath       = "/search"
	Identifier       = "identifier"
	Query            = "query"

	ProfileBudget = 15000

	DefaultCollection = "memories"
	MemoryTable       = "memory"
	VersionTable      = "memory_version"
	RelationTable     = "memory_relation"
	MemorySourceType  = "memory"

	FixtureName    = "name"
	FixtureContent = "content"
)

var RelationTypes = []string{
	"affinity",
	"informs",
	"grounds",
	"mechanism",
	"sequence",
	"deep-dive",
}

var (
	ErrorAlwaysLoad     = errors.New("failed to load always memories")
	ErrorRelevantSearch = errors.New("failed to search relevant memories")
	ErrorMemoryList     = errors.New("failed to list memories")
	ErrorReservedScope  = errors.New("scope name is reserved")
	ErrorRelationType   = errors.New("unknown relation type")
)
