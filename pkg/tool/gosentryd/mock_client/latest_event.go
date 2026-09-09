package mock_client

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

func (c *Client) LatestEvent(
	_ string,
	_ string,
) (*response.Event, error) {
	return nil, nil
}
