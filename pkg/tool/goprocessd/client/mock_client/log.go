package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/generated/client"
)

func (c *Client) Log(
	_ context.Context,
	name string,
	_ bool,
) (*client.Log, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if _, e := c.index(name); e != nil {
		return nil, e
	}

	return &client.Log{Lines: c.Logs[name]}, nil
}
