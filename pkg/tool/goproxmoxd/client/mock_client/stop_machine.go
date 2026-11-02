package mock_client

import (
	"context"
	proxmoxConstant "github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) StopMachine(
	_ context.Context,
	identifier int64,
	_ *client.StopMachineParams,
) (*client.TaskResult, error) {
	return c.setMachineStatus(identifier, proxmoxConstant.StoppedStatus)
}
