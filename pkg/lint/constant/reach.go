package constant

import "regexp"

const (
	TranscriptToolUse    = "tool_use"
	TranscriptToolResult = "tool_result"
	SourceToolPrefix     = "mcp__source__"
	SourceUseModule      = "mcp__source__use_module"
	BashTool             = "Bash"
	FilePathInput        = "file_path"
	ModuleInput          = "module"
	CommandInput         = "command"
	DryRunInput          = "dry_run"

	NullSeparated        = "-z"
	AllUntracked         = "--untracked-files=all"
	PorcelainStatusWidth = 3

	ProjectsDirectory   = ".claude/projects"
	TranscriptSuffix    = ".jsonl"
	TranscriptLineLimit = 64 << 20

	ChangeDirectoryCommand = "cd"
	CopyCommand            = "cp"
	GitCommand             = "git"
	NullDevice             = "/dev/null"
	InPlaceFlag            = "-i"
	InPlaceLongFlag        = "--in-place"
	FixFlag                = "--fix"
	DiffFlag               = "--diff"
	WriteFlag              = "-w"
)

var (
	EditTools     = []string{"Edit", "Write", "NotebookEdit", "MultiEdit"}
	SourceLocated = regexp.MustCompile(`^(\S+\.(go|md)):\d+: `)
	SourceRenamed = regexp.MustCompile(`^(\S+\.go): (\S+\.go) → (\S+\.go)$`)
	SourceEmptied = regexp.MustCompile(`^(\S+\.go): (empty file|.* → .*)$`)

	PathCommands    = []string{"mv", "rm", "rmdir", "touch", "tee"}
	GitPathCommands = []string{"mv", "rm"}
	SedCommands     = []string{"sed", "gsed"}
	ScriptCommands  = []string{"python", "python3", "perl"}
)
