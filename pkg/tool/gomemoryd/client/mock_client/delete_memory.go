package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
)

func (c *Client) DeleteMemory(
	context.Context,
	int64,
	*client.DeleteMemoryParams,
) (*client.MemoryResponse, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return &client.MemoryResponse{}, nil
}
