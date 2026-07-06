package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) Projects() ([]*project.Project, error) {
	if len(c.groups) > 0 {
		return c.GroupProjects(c.groups...)
	}

	var result []*gitlab.Project
	var number int64

	for {
		page, _, e := c.client.Projects.ListProjects(
			&gitlab.ListProjectsOptions{
				ListOptions: gitlab.ListOptions{
					PerPage: constant.PerPage100,
					Page:    number,
				},
			},
		)

		if e != nil {
			return nil, e
		}

		result = append(result, page...)

		if int64(len(page)) < constant.PerPage100 {
			break
		}

		number++
	}

	return project.NewSlice(result), nil
}
