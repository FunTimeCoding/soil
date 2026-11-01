package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/project"

func (c *Client) ProjectByName(_ string, _ string) (*project.Project, error) {
	return nil, c.failure
}
