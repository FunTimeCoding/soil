package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
)

func (c *Client) PostMemory(
	context.Context,
	client.PostMemoryJSONRequestBody,
) (*client.MemoryResponse, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return &client.MemoryResponse{}, nil
}
