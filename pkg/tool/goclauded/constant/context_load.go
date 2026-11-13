package constant

import (
	"regexp"
	"time"
)

const (
	MemoryToolPrefix = "mcp__memory__"
	SkillTool        = "Skill"
	ModeSkill        = "mode"
	ReadTool         = "Read"

	ProfileTool      = "mcp__memory__profile"
	ListMemoriesTool = "mcp__memory__list_memories"
	GetMemoryTool    = "mcp__memory__get_memory"
	SearchMemoryTool = "mcp__memory__search_memories"

	Skill    = "skill"
	Tag      = "tag"
	FilePath = "file_path"

	LoadKindMemory  = "memory"
	LoadKindProfile = "profile"
	LoadKindMode    = "mode"
	LoadKindSearch  = "search"

	TierAlways   = "always"
	TierRelevant = "relevant"

	ContextLoadTable = "context_load"
	RedactedName     = "hidden"

	PersistedOutput     = "<persisted-output>"
	ToolResultDirectory = "/tool-results/"
)

const PersistedProfileWindow = 24 * time.Hour

var (
	FollowedTools    = []string{MemoryToolPrefix, SkillTool, ReadTool}
	NumberedLine     = regexp.MustCompile(`^\s*\d+(\t|→)`)
	FixtureTimestamp = regexp.MustCompile(`"timestamp": ?"[^"]*"`)
)
