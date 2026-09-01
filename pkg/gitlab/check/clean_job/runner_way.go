package clean_job

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/runner"
)

func RunnerWay(
	g *gitlab.Client,
	r *runner.Runner,
	f *option.Format,
) {
	console.Format("Runner: %s\n", r.Format(f))
	jobs := g.MustRunnerJobs(r.Identifier, 1000)
	console.Format("Job count: %d\n", len(jobs))
	f2 := f.Copy().Extended()

	for _, j := range jobs {
		if j.Fail() {
			console.Line(j.Format(f2))
		} else {
			console.Line(j.Format(f))
		}

		if j.HasConcern(constant.JobTimeout) {
			console.Format("  Start timeout: %s\n", j.Stage)

			if console.AskConfirmation("Retry job?") {
				// TODO: Untested
				g.MustRetryJob(j)
			}
		}
	}
}
