package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/generated/client"
)

func (c *Client) GetLogs(
	context.Context,
	*client.GetLogsParams,
) (client.LogsResponse, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return client.LogsResponse{}, c.Failure
	}

	return c.Logs, nil
}
