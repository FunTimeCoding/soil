package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/client"
)

func (c *Client) ScoreTask(
	_ context.Context,
	_ string,
	_ string,
) (*client.ScoreResult, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.ScoreValue, nil
}
