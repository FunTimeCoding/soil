package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/project"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) SearchProject(query string) ([]*project.Project, error) {
	result, _, e := c.client.Search.Projects(query, &gitlab.SearchOptions{})

	if e != nil {
		return nil, wrapError(e)
	}

	return project.NewSlice(result), nil
}
