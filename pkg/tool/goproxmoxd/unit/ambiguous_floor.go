package unit

import (
	"github.com/funtimecoding/soil/pkg/proxmox/constant"
	netboxConstant "github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/mock_client"
)

func ambiguousFloor() (*mock_client.Client, []*client.VirtualMachine) {
	c := mock_client.New()
	c.AddNode("test")
	c.AddMachine("test", 101, "alfa")
	c.SetMachineStatus("test", 101, constant.RunningStatus)
	here := netboxMachine("alfa", netboxConstant.StatusOffline)
	here.Cluster = new("test")
	there := netboxMachine("alfa", netboxConstant.StatusOffline)
	there.Cluster = new("other")

	return c, []*client.VirtualMachine{here, there}
}
