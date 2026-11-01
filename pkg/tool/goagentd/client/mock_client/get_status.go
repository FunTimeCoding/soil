package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/generated/client"
)

func (c *Client) GetStatus(context.Context) (client.StatusResponse, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.Asked++

	if c.Failure != nil {
		return client.StatusResponse{}, c.Failure
	}

	return c.Status, nil
}
