package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) ListStorages(
	_ context.Context,
	node string,
	_ *client.ListStoragesParams,
) ([]client.Storage, error) {
	if e := c.node(node); e != nil {
		return nil, e
	}

	return c.StorageList, nil
}
