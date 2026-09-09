package mock_client

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

func (c *Client) Releases(
	_ string,
	_ string,
	_ int,
) ([]response.Release, error) {
	return nil, nil
}
