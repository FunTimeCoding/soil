package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/project"

func (c *Client) FileExists(_ *project.Project, _ string, _ string) (bool, error) {
	return false, c.failure
}
