package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"
)

func (c *Client) EnableRedirection(
	_ context.Context,
	p client.RedirectionRequest,
) (*client.RedirectionResult, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	result := &client.RedirectionResult{Enabled: []string{}}

	if p.Instance != nil {
		result.Instance = *p.Instance
	}

	return result, nil
}
