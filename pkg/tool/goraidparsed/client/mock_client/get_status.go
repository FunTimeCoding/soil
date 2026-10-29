package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/generated/client"
)

func (c *Client) GetStatus(context.Context) (client.StatusResponse, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return client.StatusResponse{}, c.Failure
	}

	return client.StatusResponse{Status: "ok"}, nil
}
