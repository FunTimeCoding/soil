package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
)

func (c *Client) GetSourcedMemories(
	context.Context,
	*client.GetSourcedMemoriesParams,
) ([]client.SourcedMemory, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return []client.SourcedMemory{}, nil
}
