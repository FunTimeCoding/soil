package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) Variables(project int64) ([]*gitlab.ProjectVariable, error) {
	var result []*gitlab.ProjectVariable
	number := int64(1)

	for {
		page, r, e := c.client.ProjectVariables.ListVariables(
			project,
			&gitlab.ListProjectVariablesOptions{
				ListOptions: gitlab.ListOptions{
					PerPage: constant.PerPage100,
					Page:    number,
				},
			},
		)

		if r != nil && r.StatusCode == 403 {
			// Do not panic
			return result, nil
		}

		if e != nil {
			return nil, e
		}

		result = append(result, page...)

		if int64(len(page)) < constant.PerPage100 {
			break
		}

		number++
	}

	return result, nil
}
