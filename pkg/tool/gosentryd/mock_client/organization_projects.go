package mock_client

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

func (c *Client) OrganizationProjects(_ string) ([]response.Project, error) {
	return nil, nil
}
