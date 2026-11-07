package constant

import "regexp"

const (
	RipgrepCommand               = "rg"
	RipgrepConfigurationVariable = "RIPGREP_CONFIG_PATH"
	RipgrepProgramReason         = "rg %s runs a program"
	RipgrepUnreadableReason      = "rg with an argument from an expansion, a glob or xargs input, or with a config file from RIPGREP_CONFIG_PATH, could run a program through --pre"
	TimeoutCommand               = "timeout"
	XargsCommand                 = "xargs"

	UnquotedExpansionCharacters = "*?[{\\"
)

var (
	RipgrepProgramOptions = []string{"--pre", "--hostname-bin"}
	CommandWrappers       = []string{
		"builtin",
		"command",
		"nice",
		"nohup",
		"noglob",
		"stdbuf",
		TimeoutCommand,
		XargsCommand,
	}
	WrapperValueOptions = map[string][]string{
		"nice":         {"-n"},
		"stdbuf":       {"-i", "-o", "-e"},
		TimeoutCommand: {"-s", "-k", "--signal", "--kill-after"},
		XargsCommand:   {"-E", "-I", "-J", "-L", "-n", "-P", "-R", "-S", "-s"},
	}

	RipgrepInvocation = regexp.MustCompile(`(^|[|&;(\s])rg(\s|$)`)
)
