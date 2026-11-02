package mock_client

import "github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"

func (c *Client) answer() (*client.QueryResult, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.Result, nil
}
