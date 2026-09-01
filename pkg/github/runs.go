package github

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/forge"
	"github.com/funtimecoding/soil/pkg/github/constant"
	"github.com/funtimecoding/soil/pkg/github/run"
)

func (c *Client) Runs(
	loadJobs bool,
	verbose bool,
) []*run.Run {
	var result []*run.Run
	cleanup := forge.AutoCleanup()
	f := constant.Format
	owner := c.MustUser().Name

	for _, a := range c.ActionRepository() {
		if verbose {
			console.Format("Repository: %s/%s\n", owner, a.Name)
		}

		for i, r := range c.MustProjectRuns(owner, a.Name) {
			if i > 0 {
				if cleanup {
					c.MustDeleteRun(owner, a.Name, r.Identifier)
				}

				continue
			}

			if verbose {
				console.Format("Run %d: %s\n", i, r.Format(f))
			}

			if loadJobs {
				r.Jobs = c.MustJobs(owner, a.Name, r.Identifier)
			}

			r.Validate()
			result = append(result, r)
		}
	}

	return result
}
