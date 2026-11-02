package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/forge"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
)

func (c *Client) Jobs(x context.Context) ([]*job.Job, error) {
	var result []*job.Job
	cleanup := forge.AutoCleanup()
	f := constant.Format
	projects, e := c.PipelineProjects(x)

	if e != nil {
		return nil, e
	}

	for _, p := range projects {
		if c.verbose {
			console.Format("Project: %s\n", p.Raw.NameWithNamespace)
		}

		jobs, g := c.ProjectJobs(x, p)

		if g != nil {
			return nil, g
		}

		for i, j := range jobs {
			if i > 0 {
				if cleanup {
					if h := c.DeletePipeline(
						x,
						p.Identifier,
						j.Raw.Pipeline.ID,
					); h != nil {
						return nil, h
					}
				}

				continue
			}

			if c.verbose {
				console.Format("  Job: %s\n", j.Format(f))
			}

			result = append(result, j)
		}
	}

	return result, nil
}
