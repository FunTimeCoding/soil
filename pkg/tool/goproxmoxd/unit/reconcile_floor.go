package unit

import (
	proxmoxConstant "github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/funtimecoding/soil/pkg/proxmox/mock_client"
	gonetboxdConstant "github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func reconcileFloor() (*mock_client.Client, []*client.VirtualMachine) {
	c := mock_client.New()
	c.AddNode("test")
	c.AddMachine("test", 101, "alfa")
	c.SetMachineStatus("test", 101, proxmoxConstant.RunningStatus)
	c.AddMachine("test", 102, "bravo")
	c.SetMachineStatus("test", 102, proxmoxConstant.StoppedStatus)
	c.AddMachine("test", 103, "charlie")
	c.SetMachineStatus("test", 103, proxmoxConstant.RunningStatus)
	c.AddMachine("test", 104, "delta")
	c.SetMachineStatus("test", 104, proxmoxConstant.RunningStatus)

	return c, []*client.VirtualMachine{
		netboxMachine("alfa", gonetboxdConstant.StatusOffline),
		netboxMachine("bravo", gonetboxdConstant.StatusActive),
		netboxMachine("charlie", gonetboxdConstant.StatusActive),
	}
}
