package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/generated/client"
)

func (c *Client) Processes(context.Context) ([]client.Process, error) {
	return c.ProcessList, c.Failure
}
