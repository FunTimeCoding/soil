package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/generated/client"
)

func (c *Client) GetReports(context.Context) ([]client.ReportResponse, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.Reports, nil
}
