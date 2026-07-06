package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/github"
	"github.com/funtimecoding/soil/pkg/github/constant"
)

func CleanJob() {
	c := github.NewEnvironment()
	f := option.ExtendedColor.Copy()
	owner := constant.LibraryNamespace
	repository := constant.LibraryRepository

	for _, w := range c.MustWorkflows(owner, repository) {
		fmt.Printf("Workflow: %s\n", w.Format(f))
	}

	for _, r := range c.MustProjectRuns(owner, repository) {
		fmt.Printf("Run: %s\n", r.Format(f))

		for _, j := range c.MustJobs(owner, repository, *r.Raw.ID) {
			fmt.Printf("  Job: %s\n", j.Format(f))
		}

		c.MustDeleteRun(owner, repository, *r.Raw.ID)
	}
}
