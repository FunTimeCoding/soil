package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
)

func (c *Client) GetStatus(_ context.Context) (*client.StatusResponse, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	result := &client.StatusResponse{TotalRecords: len(c.alerts)}

	for _, a := range c.alerts {
		result.LatestIdentifier = max(result.LatestIdentifier, a.Identifier)
	}

	return result, nil
}
