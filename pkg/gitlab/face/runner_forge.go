package face

import (
	"github.com/funtimecoding/soil/pkg/gitlab/job"
	"github.com/funtimecoding/soil/pkg/gitlab/response"
	"github.com/funtimecoding/soil/pkg/gitlab/runner"
)

type RunnerForge interface {
	FindRunnerByDescription(
		match string,
	) (*runner.Runner, bool, error)
	GraphRunner(identifier int64) (*response.Runner, error)
	MustGraphRunner(identifier int64) *response.Runner
	MustRunner(identifier int64) *runner.Runner
	MustRunnerJobs(
		runner int64,
		stopAfter int,
	) []*job.Job
	MustRunners(all bool) []*runner.Runner
	Runner(identifier int64) (*runner.Runner, error)
	RunnerJobs(
		runner int64,
		stopAfter int,
	) ([]*job.Job, error)
	Runners(all bool) ([]*runner.Runner, error)
	RunnersByTag(tag string) ([]*runner.Runner, error)
}
