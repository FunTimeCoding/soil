package mock_client

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

func (c *Client) SearchEvents(
	_ string,
	_ string,
	project string,
	_ int,
	_ string,
) ([]response.EventRow, error) {
	c.eventSearchProject = project

	return nil, c.Failure
}
