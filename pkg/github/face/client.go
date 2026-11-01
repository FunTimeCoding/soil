package face

import (
	"github.com/funtimecoding/soil/pkg/github/code"
	"github.com/funtimecoding/soil/pkg/github/container"
	"github.com/funtimecoding/soil/pkg/github/image"
	"github.com/funtimecoding/soil/pkg/github/issue"
	"github.com/funtimecoding/soil/pkg/github/job"
	"github.com/funtimecoding/soil/pkg/github/pull_request"
	"github.com/funtimecoding/soil/pkg/github/release"
	"github.com/funtimecoding/soil/pkg/github/repository"
	"github.com/funtimecoding/soil/pkg/github/run"
	"github.com/funtimecoding/soil/pkg/github/user"
	"github.com/funtimecoding/soil/pkg/github/workflow"
	"github.com/google/go-github/v93/github"
)

type Client interface {
	ActionRepository() ([]*repository.Repository, error)
	BranchRequest(
		owner string,
		repository string,
		branch string,
	) (*pull_request.Request, error)
	Close()
	DeletePackageVersion(
		packageName string,
		versionIdentifier int64,
	) error
	DeleteRun(
		owner string,
		name string,
		run int64,
	) error
	DeleteTag(
		owner string,
		repository string,
		name string,
	) error
	FailedRuns(verbose bool) ([]*run.Run, error)
	FindBranchRequest(
		owner string,
		repository string,
		branch string,
	) (*pull_request.Request, bool, error)
	Issues() ([]*issue.Issue, error)
	Jobs(
		owner string,
		name string,
		run int64,
	) ([]*job.Job, error)
	LatestRuns(owner, repo string) ([]*run.Run, error)
	MustActionRepository() []*repository.Repository
	MustBranchRequest(
		owner string,
		repository string,
		branch string,
	) *pull_request.Request
	MustDeletePackageVersion(
		packageName string,
		versionIdentifier int64,
	)
	MustDeleteRun(
		owner string,
		name string,
		run int64,
	)
	MustDeleteTag(
		owner string,
		repository string,
		name string,
	)
	MustJobs(
		owner string,
		name string,
		run int64,
	) []*job.Job
	MustPackageVersions(packageName string) []*image.Image
	MustPackages(user string) []*container.Container
	MustProjectIssues(
		owner string,
		name string,
	) []*github.Issue
	MustProjectRuns(
		owner string,
		name string,
	) []*run.Run
	MustPullRequests(
		owner string,
		name string,
	) []*github.PullRequest
	MustReleases(
		namespace string,
		repository string,
	) []*release.Release
	MustRepositories(owner string) []*github.Repository
	MustRuns(
		loadJobs bool,
		verbose bool,
	) []*run.Run
	MustSearchRepository(
		query string,
		a ...any,
	) []*repository.Repository
	MustTags(
		owner string,
		repository string,
	) []*github.RepositoryTag
	MustUser() *user.User
	MustWorkflows(
		owner string,
		name string,
	) []*workflow.Workflow
	Notifications() ([]*github.Notification, error)
	OpenUserIssues() ([]*issue.Issue, error)
	Organizations(user string) ([]*github.Organization, error)
	PackageVersions(packageName string) ([]*image.Image, error)
	Packages(user string) ([]*container.Container, error)
	ProjectIssues(
		owner string,
		name string,
	) ([]*github.Issue, error)
	ProjectRuns(
		owner string,
		name string,
	) ([]*run.Run, error)
	PullRequests(
		owner string,
		name string,
	) ([]*github.PullRequest, error)
	Releases(
		namespace string,
		repository string,
	) ([]*release.Release, error)
	Repositories(owner string) ([]*github.Repository, error)
	Runs(
		loadJobs bool,
		verbose bool,
	) ([]*run.Run, error)
	SearchCode(
		query string,
		a ...any,
	) ([]*code.Code, error)
	SearchIssue(
		query string,
		a ...any,
	) ([]*issue.Issue, error)
	SearchRepository(
		query string,
		a ...any,
	) ([]*repository.Repository, error)
	Tags(
		owner string,
		repository string,
	) ([]*github.RepositoryTag, error)
	User() (*user.User, error)
	Workflows(
		owner string,
		name string,
	) ([]*workflow.Workflow, error)
}
