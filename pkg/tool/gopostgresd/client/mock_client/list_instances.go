package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
)

func (c *Client) ListInstances(_ context.Context) ([]client.Instance, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.Instances, nil
}
