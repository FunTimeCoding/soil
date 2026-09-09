package mock_client

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

func (c *Client) Event(
	_ string,
	_ string,
	_ string,
) (*response.Event, error) {
	return nil, nil
}
