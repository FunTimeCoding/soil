package face

import (
	"github.com/funtimecoding/soil/pkg/github/job"
	"github.com/funtimecoding/soil/pkg/github/repository"
	"github.com/funtimecoding/soil/pkg/github/run"
	"github.com/funtimecoding/soil/pkg/github/workflow"
)

type ActionClient interface {
	ActionRepository() ([]*repository.Repository, error)
	DeleteRun(
		owner string,
		name string,
		run int64,
	) error
	FailedRuns(verbose bool) ([]*run.Run, error)
	Jobs(
		owner string,
		name string,
		run int64,
	) ([]*job.Job, error)
	LatestRuns(owner, repo string) ([]*run.Run, error)
	MustActionRepository() []*repository.Repository
	MustDeleteRun(
		owner string,
		name string,
		run int64,
	)
	MustJobs(
		owner string,
		name string,
		run int64,
	) []*job.Job
	MustProjectRuns(
		owner string,
		name string,
	) []*run.Run
	MustRuns(
		loadJobs bool,
		verbose bool,
	) []*run.Run
	MustWorkflows(
		owner string,
		name string,
	) []*workflow.Workflow
	ProjectRuns(
		owner string,
		name string,
	) ([]*run.Run, error)
	Runs(
		loadJobs bool,
		verbose bool,
	) ([]*run.Run, error)
	Workflows(
		owner string,
		name string,
	) ([]*workflow.Workflow, error)
}
