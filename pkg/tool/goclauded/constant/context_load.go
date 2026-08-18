package constant

const (
	MemoryToolPrefix = "mcp__memory__"
	SkillTool        = "Skill"
	ModeSkill        = "mode"

	ProfileTool      = "mcp__memory__profile"
	ListMemoriesTool = "mcp__memory__list_memories"
	GetMemoryTool    = "mcp__memory__get_memory"
	SearchMemoryTool = "mcp__memory__search_memories"

	Skill = "skill"
	Tag   = "tag"

	LoadKindMemory  = "memory"
	LoadKindProfile = "profile"
	LoadKindMode   = "mode"
	LoadKindSearch = "search"

	TierAlways   = "always"
	TierRelevant = "relevant"

	ContextLoadTable = "context_load"
	RedactedName     = "hidden"
)

var FollowedTools = []string{MemoryToolPrefix, SkillTool}
