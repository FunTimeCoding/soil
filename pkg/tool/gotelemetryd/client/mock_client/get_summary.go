package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/generated/client"
)

func (c *Client) GetSummary(
	_ context.Context,
	_ *client.GetSummaryParams,
) ([]client.SummaryEntry, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	return nil, nil
}
