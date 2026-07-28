package constant

import "github.com/funtimecoding/soil/pkg/console/status/option"

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

	RevParse     = "rev-parse"
	GitDirectory = "--git-dir"

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

	Format = option.ExtendedColor.Copy()
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
