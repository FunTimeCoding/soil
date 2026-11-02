package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) RollbackContainerSnapshot(
	_ context.Context,
	identifier int64,
	name string,
	_ *client.RollbackContainerSnapshotParams,
) (*client.TaskResult, error) {
	if _, e := c.container(identifier); e != nil {
		return nil, e
	}

	if _, e := snapshot(c.ContainerSnapshots[identifier], name); e != nil {
		return nil, e
	}

	return c.task(), nil
}
