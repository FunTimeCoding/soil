package mock_client

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

func (c *Client) Teams(_ string) ([]response.Team, error) {
	return nil, c.Failure
}
