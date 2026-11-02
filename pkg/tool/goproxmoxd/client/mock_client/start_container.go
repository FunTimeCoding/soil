package mock_client

import (
	"context"
	proxmoxConstant "github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) StartContainer(
	_ context.Context,
	identifier int64,
	_ *client.StartContainerParams,
) (*client.TaskResult, error) {
	return c.setContainerStatus(identifier, proxmoxConstant.RunningStatus)
}
