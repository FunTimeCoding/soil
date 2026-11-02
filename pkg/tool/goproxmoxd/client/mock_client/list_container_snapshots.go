package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) ListContainerSnapshots(
	_ context.Context,
	identifier int64,
	_ *client.ListContainerSnapshotsParams,
) ([]client.Snapshot, error) {
	if _, e := c.container(identifier); e != nil {
		return nil, e
	}

	return c.ContainerSnapshots[identifier], nil
}
