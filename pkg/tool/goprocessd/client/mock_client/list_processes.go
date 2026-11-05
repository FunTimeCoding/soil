package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/generated/client"
)

func (c *Client) ListProcesses(context.Context) ([]client.Process, error) {
	return c.ProcessList, c.Failure
}
