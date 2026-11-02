package worker_tester

import (
	proxmoxConstant "github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/funtimecoding/soil/pkg/proxmox/mock_client"
)

func PopulatedClient(name string) *mock_client.Client {
	c := mock_client.New()
	c.AddNode(name)
	c.AddMachine(name, 100, "first")
	c.SetMachineStatus(name, 100, proxmoxConstant.RunningStatus)
	c.AddMachine(name, 101, "second")
	c.AddGuestNotInBackup(proxmoxConstant.MachineType, 101, "second")
	c.AddUpdatePending(name, "pve-docs", "9.2.3", "9.2.4")

	return c
}
