package constant

import "github.com/funtimecoding/soil/pkg/console/constant"

const (
	VersionPrefix = "v"

	OriginRemote = "origin"

	Directory = ".git"

	MainBranch   = "main"
	MasterBranch = "master"

	GitHubHost = "github.com"
	GitLabHost = "gitlab.com"

	Command = "git"
	Tag     = "tag"
	Clone   = "clone"
	Status  = "status"
	Log     = "log"
	Diff    = "diff"

	NameOnly         = "--name-only"
	Relative         = "--relative"
	CommitTimeFormat = "--format=%ct"
	Pathspec         = "--"

	RevParse             = "rev-parse"
	GitDirectory         = "--git-dir"
	AbbreviatedReference = "--abbrev-ref"
	SymbolicFullName     = "--symbolic-full-name"
	Upstream             = "@{upstream}"
	Cached               = "--cached"
	ShowToplevel         = "--show-toplevel"
	HooksDirectory       = "hooks"
	Configuration        = "config"
	Get                  = "--get"
	HooksPathKey         = "core.hooksPath"
	RemoteHead           = "refs/remotes/origin/HEAD"

	HookPreCommit     = "pre-commit"
	HookPrePush       = "pre-push"
	HookCommitMessage = "commit-msg"
	HookPostCheckout  = "post-checkout"
	HookPostMerge     = "post-merge"

	RangeAll    = "all files"
	RangeNone   = "nothing pushed"
	ZeroHash    = "0000000000000000000000000000000000000000"
	RangeStaged = "staged files"

	Porcelain = "--porcelain"

	Fetch     = "fetch"
	Prune     = "--prune"
	PruneTags = "--prune-tags"

	Push = "push"
	Tags = "--tags"

	HeadReference = "HEAD"

	HashLength = 7
)

var (
	MainBranches = []string{MainBranch, MasterBranch}

	Hooks = []string{
		HookPreCommit,
		HookPrePush,
		HookCommitMessage,
		HookPostCheckout,
		HookPostMerge,
	}

	Format = constant.ExtendedColorFormat.Copy()
)

const (
	RepositoryRootEnvironment    = "REPOSITORY_ROOT"
	RepositoryExcludeEnvironment = "REPOSITORY_EXCLUDE"
)

const (
	UnknownProvider = "unknown"
	GitLabProvider  = "gitlab"
	GitHubProvider  = "github"
)

const NotClean = "not_clean"
