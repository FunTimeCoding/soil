package constant

import "regexp"

const (
	GitHistoryMessage    = "git %s writes history - staging, committing and publishing belong to the user; leave it to them"
	GitUnreadableMessage = "git %s with a path from an expansion, a glob or stdin, or with nothing after --, cannot be read before it runs - write every path out literally"
	GitAskReason         = "git %s rewrites the shared index or work tree: %s"
	GitWholeTree         = "everything it covers"
	GitHiddenMessage     = "git inside a shell string or eval that cannot be read before it runs - run git directly, with every path written out"
	EvalCommand          = "eval"
	GitTagSubcommand     = "tag"
	GitStashSubcommand   = "stash"
	GitMoveSubcommand    = "mv"

	PreToolUseEvent = "PreToolUse"
	AskDecision     = "ask"
)

var (
	GitHistoryWriters = []string{
		"add",
		"am",
		"cherry-pick",
		"commit",
		"merge",
		"push",
		"rebase",
		"revert",
		"tag",
	}
	GitTreeWriters = []string{
		"checkout",
		"clean",
		"mv",
		"reset",
		"restore",
		"rm",
		"stash",
		"update-index",
	}
	GitStashVerbs = []string{
		"apply",
		"branch",
		"clear",
		"create",
		"drop",
		"list",
		"pop",
		"push",
		"save",
		"show",
		"store",
	}
	GitStashReadVerbs   = []string{"list", "show"}
	GitDryRunOptions    = []string{"-n", "--dry-run"}
	GitOptionsWithValue = []string{"-C", "-c"}
	GitTagListOptions   = []string{"-l", "--list", "-v", "--verify"}
	GitTagReadOptions   = []string{
		"-i",
		"--ignore-case",
		"--column",
		"--no-column",
		"--omit-empty",
	}
	GitTagValueOptions = []string{
		"--sort",
		"--format",
		"--contains",
		"--no-contains",
		"--merged",
		"--no-merged",
		"--points-at",
	}
	ScriptShells        = []string{"bash", "dash", "ksh", "sh", "zsh"}

	GitTagLineCount = regexp.MustCompile(`^-n[0-9]*$`)
	GitInvocation   = regexp.MustCompile(`(^|[|&;(\s])git(\s|$)`)
	GitMention    = regexp.MustCompile(`(^|[^\w./-])git(\s|$)`)
)
