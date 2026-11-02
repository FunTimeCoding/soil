package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"slices"
)

func (c *Client) DeleteContainerSnapshot(
	_ context.Context,
	identifier int64,
	name string,
	_ *client.DeleteContainerSnapshotParams,
) (*client.TaskResult, error) {
	if _, e := c.container(identifier); e != nil {
		return nil, e
	}

	i, e := snapshot(c.ContainerSnapshots[identifier], name)

	if e != nil {
		return nil, e
	}

	c.ContainerSnapshots[identifier] = slices.Delete(
		c.ContainerSnapshots[identifier],
		i,
		i+1,
	)

	return c.task(), nil
}
