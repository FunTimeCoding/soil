package hub

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/github"
	"github.com/funtimecoding/soil/pkg/github/constant"
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
		console.Format("Latest run: %s\n", latestRun.Name)
	}

	for _, r := range runs {
		if r.Status != constant.CompletedStatus {
			continue
		}

		console.Format("Delete run: %d\n", r.Identifier)
		c.MustDeleteRun(namespace, repository, r.Identifier)
	}
}
