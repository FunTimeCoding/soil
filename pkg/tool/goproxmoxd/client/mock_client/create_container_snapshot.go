package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) CreateContainerSnapshot(
	_ context.Context,
	identifier int64,
	_ *client.CreateContainerSnapshotParams,
	b client.CreateContainerSnapshotJSONRequestBody,
) (*client.TaskResult, error) {
	if _, e := c.container(identifier); e != nil {
		return nil, e
	}

	c.ContainerSnapshots[identifier] = append(
		c.ContainerSnapshots[identifier],
		client.Snapshot{Name: b.Name},
	)

	return c.task(), nil
}
