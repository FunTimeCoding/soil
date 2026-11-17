package offline_memory

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
)

func (c *Client) PostMemory(
	context.Context,
	client.PostMemoryJSONRequestBody,
) (*client.MemoryResponse, error) {
	return nil, nil
}
