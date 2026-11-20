package constant

import (
	"regexp"
	"time"
)

const GrantLookupTimeout = 2 * time.Second
const (
	GitHistoryReason    = "git %s writes history - staging, committing and publishing belong to the user; leave it to them"
	GitUnreadableReason = "git %s with a path from an expansion, a glob or stdin, or with nothing after --, cannot be read before it runs - write every path out literally"
	GitAskReason        = "git %s rewrites the shared index or work tree: %s"
	GitWholeTree        = "everything it covers"
	GitHiddenReason     = "git inside a shell string or eval that cannot be read before it runs - run git directly, with every path written out"
	EvalCommand         = "eval"
	GitTagSubcommand    = "tag"
	GitStashSubcommand  = "stash"
	GitMoveSubcommand   = "mv"

	GitWriteLabel        = "git-write"
	GitPushSubcommand    = "push"
	GitDirectoryOption   = "-C"
	GitMetadataDirectory = ".git"
	GitGrantedNote       = "git %s allowed by the %s label for %s"
	GitGrantUnreadReason = "git %s writes history, and the %s label could not be read from goclauded (%v) - leave it to the user"
	GitRefusedPushShort  = "df"

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
	GitOptionsWithValue = []string{
		"-C",
		"-c",
		"--git-dir",
		"--work-tree",
		"--namespace",
		"--config-env",
	}
	GitRedirectOptions = []string{
		"-c",
		"--git-dir",
		"--work-tree",
		"--namespace",
		"--config-env",
		"--exec-path",
		"--super-prefix",
	}
	GitGrantedWriters     = []string{"add", "commit", "push"}
	GitRefusedPushOptions = []string{
		"--force",
		"--force-with-lease",
		"--force-if-includes",
		"--tags",
		"--follow-tags",
		"--mirror",
		"--delete",
		"--all",
		"--prune",
		"--repo",
		"--receive-pack",
		"--exec",
	}
	DirectoryCommands = []string{"cd", "pushd", "popd"}
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
	ScriptShells = []string{"bash", "dash", "ksh", "sh", "zsh"}

	GitTagLineCount = regexp.MustCompile(`^-n[0-9]*$`)
	GitInvocation   = regexp.MustCompile(`(^|[|&;(\s])git(\s|$)`)
	GitMention      = regexp.MustCompile(`(^|[^\w./-])git(\s|$)`)
	GitEnvironment  = regexp.MustCompile(`\bGIT_[A-Z_]+=`)
	GitRemoteName   = regexp.MustCompile(`^[A-Za-z0-9._-]+$`)
)
