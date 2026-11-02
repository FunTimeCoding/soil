package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) GetContainer(
	_ context.Context,
	identifier int64,
	_ *client.GetContainerParams,
) (*client.ContainerDetail, error) {
	i, e := c.container(identifier)

	if e != nil {
		return nil, e
	}

	v := c.ContainerList[i]

	return &client.ContainerDetail{
		Identifier: v.Identifier,
		Name:       v.Name,
		Node:       v.Node,
		Status:     v.Status,
		Tags:       v.Tags,
	}, nil
}
