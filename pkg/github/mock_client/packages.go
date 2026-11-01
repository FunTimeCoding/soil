package mock_client

import "github.com/funtimecoding/soil/pkg/github/container"

func (c *Client) Packages(_ string) ([]*container.Container, error) {
	return nil, c.Failure
}
