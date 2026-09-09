package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/project"

func (c *Client) Projects() ([]*project.Project, error) {
	return nil, nil
}
