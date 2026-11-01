package unit

import (
	"github.com/funtimecoding/soil/pkg/proxmox/constant"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/client/mock_client"
	netboxConstant "github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/mock_client"
)

func reconcileFloor() (*mock_client.Client, *netbox.Client) {
	c := mock_client.New()
	c.AddNode("test")
	c.AddMachine("test", 101, "alfa")
	c.SetMachineStatus("test", 101, constant.RunningStatus)
	c.AddMachine("test", 102, "bravo")
	c.SetMachineStatus("test", 102, constant.StoppedStatus)
	c.AddMachine("test", 103, "charlie")
	c.SetMachineStatus("test", 103, constant.RunningStatus)
	c.AddMachine("test", 104, "delta")
	c.SetMachineStatus("test", 104, constant.RunningStatus)
	n := netbox.New()
	n.VirtualMachineList = []*client.VirtualMachine{
		netboxMachine("alfa", netboxConstant.StatusOffline),
		netboxMachine("bravo", netboxConstant.StatusActive),
		netboxMachine("charlie", netboxConstant.StatusActive),
	}

	return c, n
}
