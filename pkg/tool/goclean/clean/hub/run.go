package hub

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/github"
	"github.com/funtimecoding/soil/pkg/github/run"
)

func Run(
	c *github.Client,
	namespace string,
	repository string,
) {
	runs := c.MustProjectRuns(namespace, repository)

	if false {
		latestRun := run.Latest(runs)
		fmt.Printf("Latest run: %s\n", latestRun.Name)
	}

	for _, r := range runs {
		if r.Status != run.Completed {
			continue
		}

		fmt.Printf("Delete run: %d\n", r.Identifier)
		c.MustDeleteRun(namespace, repository, r.Identifier)
	}
}
