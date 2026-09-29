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
) ([]*run.Run, error) {
	var result []*run.Run
	cleanup := forge.AutoCleanup()
	f := constant.Format
	u, e := c.User()

	if e != nil {
		return nil, e
	}

	owner := u.Name
	repositories, g := c.ActionRepository()

	if g != nil {
		return nil, g
	}

	for _, a := range repositories {
		if verbose {
			console.Format("Repository: %s/%s\n", owner, a.Name)
		}

		runs, h := c.ProjectRuns(owner, a.Name)

		if h != nil {
			return nil, h
		}

		for i, r := range runs {
			if i > 0 {
				if cleanup {
					if deleteFail := c.DeleteRun(
						owner,
						a.Name,
						r.Identifier,
					); deleteFail != nil {
						return nil, deleteFail
					}
				}

				continue
			}

			if verbose {
				console.Format("Run %d: %s\n", i, r.Format(f))
			}

			if loadJobs {
				jobs, jobFail := c.Jobs(owner, a.Name, r.Identifier)

				if jobFail != nil {
					return nil, jobFail
				}

				r.Jobs = jobs
			}

			r.Validate()
			result = append(result, r)
		}
	}

	return result, nil
}
