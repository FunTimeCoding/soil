package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/project"

func (c *Client) GroupProjects(_ ...int64) ([]*project.Project, error) {
	return nil, c.failure
}
