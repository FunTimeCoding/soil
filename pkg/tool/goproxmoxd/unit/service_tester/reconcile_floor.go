package service_tester

import (
	proxmoxConstant "github.com/funtimecoding/soil/pkg/proxmox/constant"
	gonetboxdConstant "github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
)

func (o *Tester) ReconcileFloor() {
	c := o.Alfa().Hypervisor
	c.AddMachine("alfa", 101, "alfa")
	c.SetMachineStatus("alfa", 101, proxmoxConstant.RunningStatus)
	c.AddMachine("alfa", 102, "bravo")
	c.SetMachineStatus("alfa", 102, proxmoxConstant.StoppedStatus)
	c.AddMachine("alfa", 103, "charlie")
	c.SetMachineStatus("alfa", 103, proxmoxConstant.RunningStatus)
	c.AddMachine("alfa", 104, "delta")
	c.SetMachineStatus("alfa", 104, proxmoxConstant.RunningStatus)

	if o.Netbox == nil {
		return
	}

	o.Netbox.VirtualMachineList = append(
		o.Netbox.VirtualMachineList,
		netboxMachine("alfa", gonetboxdConstant.StatusOffline),
		netboxMachine("bravo", gonetboxdConstant.StatusActive),
		netboxMachine("charlie", gonetboxdConstant.StatusActive),
	)
}
