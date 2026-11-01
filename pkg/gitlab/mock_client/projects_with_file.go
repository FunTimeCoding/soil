package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/project"

func (c *Client) ProjectsWithFile(_ string, _ bool) ([]*project.Project, error) {
	return nil, c.failure
}
