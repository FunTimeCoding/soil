package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) Projects(x context.Context) ([]*project.Project, error) {
	if len(c.projects) > 0 {
		var selected []*project.Project

		for _, i := range c.projects {
			p, e := c.Project(x, i)

			if e != nil {
				return nil, e
			}

			selected = append(selected, p)
		}

		return selected, nil
	}

	if len(c.groups) > 0 {
		return c.GroupProjects(x, c.groups...)
	}

	var result []*gitlab.Project
	number := int64(1)

	for {
		page, _, e := c.client.Projects.ListProjects(
			&gitlab.ListProjectsOptions{
				ListOptions: gitlab.ListOptions{
					PerPage: constant.PerPage100,
					Page:    number,
				},
			},
			gitlab.WithContext(x),
		)

		if e != nil {
			return nil, wrapError(e)
		}

		result = append(result, page...)

		if int64(len(page)) < constant.PerPage100 {
			break
		}

		number++
	}

	return project.NewSlice(result), nil
}
