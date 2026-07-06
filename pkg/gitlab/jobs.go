package gitlab

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/forge"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
)

func (c *Client) Jobs() []*job.Job {
	var result []*job.Job
	cleanup := forge.AutoCleanup()
	f := constant.Format

	for _, p := range c.PipelineProjects() {
		if c.verbose {
			fmt.Printf("Project: %s\n", p.Raw.NameWithNamespace)
		}

		for i, j := range c.MustProjectJobs(p) {
			if i > 0 {
				if cleanup {
					c.MustDeletePipeline(p.Identifier, j.Raw.Pipeline.ID)
				}

				continue
			}

			if c.verbose {
				fmt.Printf("  Job: %s\n", j.Format(f))
			}

			result = append(result, j)
		}
	}

	return result
}
