package mock_client

import "github.com/funtimecoding/soil/pkg/github/code"

func (c *Client) SearchCode(_ string, _ ...any) ([]*code.Code, error) {
	return nil, c.Failure
}
