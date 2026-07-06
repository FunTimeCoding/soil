package gitlab

import "github.com/funtimecoding/soil/pkg/gitlab/job"

func (c *Client) enrichJob(j *job.Job) *job.Job {
	if j.Raw.Project.ID == 0 {
		panic("project cannot be 0")
	}

	j.Project = c.MustProject(j.Raw.Project.ID)
	c.enrichJobCommon(j)

	return j
}
