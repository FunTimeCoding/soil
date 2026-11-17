package offline_memory

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
)

func (c *Client) DeleteMemory(
	context.Context,
	int64,
	*client.DeleteMemoryParams,
) (*client.MemoryResponse, error) {
	return nil, nil
}
