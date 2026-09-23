package worker_tester

import (
	"github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/mock_client"
)

func PopulatedClient(name string) *mock_client.Client {
	c := mock_client.New()
	c.AddNode(name)
	c.AddMachine(name, 100, "first")
	c.SetMachineStatus(name, 100, constant.RunningStatus)
	c.AddMachine(name, 101, "second")
	c.AddGuestNotInBackup(constant.MachineType, 101, "second")
	c.AddUpdatePending(name, "pve-docs", "9.2.3", "9.2.4")

	return c
}
