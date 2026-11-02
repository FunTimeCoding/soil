package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) ListContainers(
	_ context.Context,
	p *client.ListContainersParams,
) ([]client.Container, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	var result []client.Container

	for _, v := range c.ContainerList {
		if p == nil || p.Node == nil || (v.Node != nil && *v.Node == *p.Node) {
			result = append(result, v)
		}
	}

	return result, nil
}
