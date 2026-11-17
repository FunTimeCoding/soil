package offline_memory

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
)

func (c *Client) PutMemory(
	context.Context,
	int64,
	client.PutMemoryJSONRequestBody,
) (*client.MemoryResponse, error) {
	return nil, nil
}
