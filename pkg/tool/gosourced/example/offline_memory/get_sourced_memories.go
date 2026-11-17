package offline_memory

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
)

func (c *Client) GetSourcedMemories(
	context.Context,
	*client.GetSourcedMemoriesParams,
) ([]client.SourcedMemory, error) {
	return nil, nil
}
