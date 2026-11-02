package mock_client

import (
	"context"
	proxmoxConstant "github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) StartMachine(
	_ context.Context,
	identifier int64,
	_ *client.StartMachineParams,
) (*client.TaskResult, error) {
	return c.setMachineStatus(identifier, proxmoxConstant.RunningStatus)
}
