package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
)

func (c *Client) GetOversize(
	_ context.Context,
	_ *client.GetOversizeParams,
) (*client.OversizeReport, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	return &client.OversizeReport{}, nil
}
