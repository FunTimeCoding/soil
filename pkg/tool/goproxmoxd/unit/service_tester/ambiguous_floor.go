package service_tester

import (
	proxmoxConstant "github.com/funtimecoding/soil/pkg/proxmox/constant"
	gonetboxdConstant "github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
)

func (o *Tester) AmbiguousFloor() {
	c := o.Alfa().Hypervisor
	c.AddMachine("alfa", 101, "alfa")
	c.SetMachineStatus("alfa", 101, proxmoxConstant.RunningStatus)

	if o.Netbox == nil {
		return
	}

	here := netboxMachine("alfa", gonetboxdConstant.StatusOffline)
	here.Cluster = new("alfa")
	there := netboxMachine("alfa", gonetboxdConstant.StatusOffline)
	there.Cluster = new("other")
	o.Netbox.VirtualMachineList = append(
		o.Netbox.VirtualMachineList,
		here,
		there,
	)
}
