package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/project"

func (c *Client) Projects() ([]*project.Project, error) {
	if c.projectFailure != nil {
		return nil, c.projectFailure
	}

	return c.projects, nil
}
