package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/github"
	"github.com/funtimecoding/soil/pkg/github/constant"
)

func CleanJob() {
	c := github.NewEnvironment()
	f := consoleConstant.ExtendedColorFormat.Copy()
	owner := constant.Namespace
	repository := constant.Repository

	for _, w := range c.MustWorkflows(owner, repository) {
		console.Format("Workflow: %s\n", w.Format(f))
	}

	for _, r := range c.MustProjectRuns(owner, repository) {
		console.Format("Run: %s\n", r.Format(f))

		for _, j := range c.MustJobs(owner, repository, *r.Raw.ID) {
			console.Format("  Job: %s\n", j.Format(f))
		}

		c.MustDeleteRun(owner, repository, *r.Raw.ID)
	}
}
