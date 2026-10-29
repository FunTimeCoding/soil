package mock_client

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

func (c *Client) OrganizationProjects(_ string) ([]response.Project, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.projects, nil
}
