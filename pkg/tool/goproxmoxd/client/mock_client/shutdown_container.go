package mock_client

import (
	"context"
	proxmoxConstant "github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) ShutdownContainer(
	_ context.Context,
	identifier int64,
	_ *client.ShutdownContainerParams,
) (*client.TaskResult, error) {
	return c.setContainerStatus(identifier, proxmoxConstant.StoppedStatus)
}
