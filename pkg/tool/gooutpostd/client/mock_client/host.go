package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/generated/client"
)

func (c *Client) Host(context.Context) (*client.Host, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.HostValue, nil
}
