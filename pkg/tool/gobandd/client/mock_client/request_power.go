package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"
)

func (c *Client) RequestPower(
	_ context.Context,
	p client.PowerRequest,
) (*client.PowerResult, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	c.Requested = append(c.Requested, p)
	result := &client.PowerResult{Requested: string(p.Operation)}

	if p.Instance != nil {
		result.Instance = *p.Instance
	}

	return result, nil
}
