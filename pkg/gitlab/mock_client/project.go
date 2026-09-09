package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/project"

func (c *Client) Project(_ int64) (*project.Project, error) {
	return nil, nil
}
