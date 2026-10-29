package mock_client

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

func (c *Client) Whoami() (*response.User, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return response.NewUser(), nil
}
