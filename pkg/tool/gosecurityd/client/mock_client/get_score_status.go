package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
)

func (c *Client) GetScoreStatus(context.Context) (*client.ScoreStatus, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.failure != nil {
		return nil, c.failure
	}

	result := c.scoreStatus

	return &result, nil
}
