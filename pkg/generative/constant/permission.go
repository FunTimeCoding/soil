package constant

const (
	BashTool           = "Bash"
	QuestionTool       = "AskUserQuestion"
	UserDecisionPrefix = "user_"
	RejectDecision     = "reject"
)

const (
	AllowlistSource = "allowlist"
	BuiltInSource   = "built-in"
	ScratchSource   = "scratch"
)

const (
	SettingsDirectory       = ".claude"
	SettingsFile            = "settings.local.json"
	SharedSettingsFile      = "settings.json"
	ProjectScratchDirectory = "tmp"
	NullDevice              = "/dev/null"
)

const (
	RulePrefix          = "Bash("
	RuleSuffix          = ")"
	RuleColonWildcard   = ":*"
	RuleSpaceWildcard   = " *"
	OptionTerminator    = "--"
	OptionPrefix        = "-"
	ChangeDirectory     = "cd"
	RemoveCommand       = "rm"
	MakeDirectory       = "mkdir"
	CopyCommand         = "cp"
	MoveCommand         = "mv"
	TargetOption        = "-t"
	TargetOptionLong    = "--target-directory"
	GlobCharacters      = "*?["
	UnsafePathCharacter = "{~"
	HiddenSegmentPrefix = "."
)

const (
	BlockerArray                   = "(array assignment)"
	BlockerAssignment              = "(assignment before a command)"
	BlockerBackground              = "(background job)"
	BlockerConstruct               = "(unsupported construct)"
	BlockerEmpty                   = "(no command)"
	BlockerExpansion               = "(command from an expansion)"
	BlockerGitAfterChangeDirectory = "(git after cd)"
	BlockerRedirect                = "(redirect outside scratch)"
	BlockerUnparsed                = "(unparsed)"
	BlockerVariable                = "(assignment to %s)"
)

const (
	FindCommand    = "find"
	GitCommand     = "git"
	TimeoutCommand = "timeout"
	XargsCommand   = "xargs"
	CommandBuiltin = "command"
	ReadCommand    = "read"
	PrintfCommand  = "printf"
	VariableOption = "-v"
	FieldSeparator = "IFS"
)

var (
	FindEscapes           = []string{"-exec", "-ok", "-delete"}
	GlobSensitiveCommands = []string{"find", "sort", "sed", "gsed", "git"}
)

var ImplicitCommands = []string{
	"[",
	"cat",
	"diff",
	"du",
	"echo",
	"false",
	"grep",
	"head",
	"ls",
	"printf",
	"pwd",
	"sleep",
	"stat",
	"tail",
	"test",
	"true",
	"wc",
	"which",
}

var Wrappers = []string{
	"builtin",
	"command",
	"nice",
	"nohup",
	"noglob",
	"stdbuf",
	"timeout",
	"xargs",
}

var WrapperValueOptions = map[string][]string{
	"nice":    {"-n"},
	"stdbuf":  {"-i", "-o", "-e"},
	"timeout": {"-s", "-k", "--signal", "--kill-after"},
}

var LocaleVariables = []string{
	"LANG",
	"LC_ALL",
	"LC_COLLATE",
	"LC_CTYPE",
	"LC_MESSAGES",
}

var DangerousVariables = []string{
	"PATH",
	"IFS",
	"BASH_ENV",
	"ENV",
	"LD_PRELOAD",
	"LD_LIBRARY_PATH",
	"DYLD_INSERT_LIBRARIES",
	"DYLD_LIBRARY_PATH",
	"PROMPT_COMMAND",
}
