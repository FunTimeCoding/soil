package constant

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"time"
)

var Identity = identity.New(
	"gomemoryd",
	"Persistent memory across Claude Code sessions",
	"gomemoryd",
).WithInstructions(
	"Persistent memory across Claude Code sessions. Call profile on your first turn to load memories. Read the gomemoryd://guide/memory-workflow resource for the memory lifecycle, tags, search, and profile tiers.",
)

const ReconcileInterval = 10 * time.Minute
const (
	HostEnvironment      = "GOMEMORY_HOST"
	TokenEnvironment     = "GOMEMORY_TOKEN"
	PortEnvironment      = "GOMEMORY_PORT"
	InsecureEnvironment  = "GOMEMORY_INSECURE"
	HiddenTagEnvironment = "GOMEMORY_HIDDEN_TAG"
	TagCutset            = "[]\" \t"
	TagStripNotice       = "Stripped [ ] \" from tag input, saved as:"

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
	Tags              = "tags"
	Base              = "base"
	ClearBase         = "clear_base"

	BaseKey = "base"

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
	Identifier       = "identifier"
	Query            = "query"
	Page             = "page"
	DetailHeadClass  = "detail-head"
	Days             = "days"

	ProfileBudget = 15000

	DefaultCollection = "memories"
	MemoryTable       = "memory"
	VersionTable      = "memory_version"
	RelationTable     = "memory_relation"
	MemorySourceType  = "memory"
	MemoryKind        = "memory"

	FixtureName    = "name"
	FixtureContent = "content"
)

var MemoryTypes = []string{"feedback", "user", "project", "reference"}
var RelationTypes = []string{
	"affinity",
	"informs",
	"grounds",
	"mechanism",
	"sequence",
	"deep-dive",
}

var ImpressionWindows = []int{2, 7, 30}

const (
	DefaultImpressionDays     = 7
	MemoryPageLimit           = 20
	DashboardVersionLimit     = 10
	DashboardVersionWindow    = 7 * 24 * time.Hour
	DashboardImpressionWindow = 48 * time.Hour
)

const (
	SearchResultsMark = "search-results"
	SearchPlaceholder = "search memories..."
	PrefixMark        = "*"
	SearchLimit       = 20
)
