package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
	"github.com/funtimecoding/soil/pkg/gitlab/response"
	"github.com/funtimecoding/soil/pkg/gitlab/runner"
)

type RunnerForge interface {
	FindRunnerByDescription(
		x context.Context,
		match string,
	) (*runner.Runner, bool, error)
	GraphRunner(
		x context.Context,
		identifier int64,
	) (*response.Runner, error)
	MustGraphRunner(
		x context.Context,
		identifier int64,
	) *response.Runner
	MustRunner(
		x context.Context,
		identifier int64,
	) *runner.Runner
	MustRunnerJobs(
		x context.Context,
		runner int64,
		stopAfter int,
	) []*job.Job
	MustRunners(
		x context.Context,
		all bool,
	) []*runner.Runner
	Runner(
		x context.Context,
		identifier int64,
	) (*runner.Runner, error)
	RunnerJobs(
		x context.Context,
		runner int64,
		stopAfter int,
	) ([]*job.Job, error)
	Runners(
		x context.Context,
		all bool,
	) ([]*runner.Runner, error)
	TagRunners(
		x context.Context,
		tag string,
	) ([]*runner.Runner, error)
}
