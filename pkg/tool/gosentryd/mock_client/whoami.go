package mock_client

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

func (c *Client) Whoami() (*response.User, error) {
	return nil, nil
}
