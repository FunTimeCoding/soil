package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) ListStorageContent(
	_ context.Context,
	node string,
	_ string,
	_ *client.ListStorageContentParams,
) ([]client.StorageContentItem, error) {
	if e := c.node(node); e != nil {
		return nil, e
	}

	return c.StorageContentList, nil
}
