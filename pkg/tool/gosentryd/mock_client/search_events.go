package mock_client

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

func (c *Client) SearchEvents(
	_ string,
	_ string,
	_ string,
	_ int,
	_ string,
) ([]response.EventRow, error) {
	return nil, nil
}
