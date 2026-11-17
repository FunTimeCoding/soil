package service_tester

import proxmoxConstant "github.com/funtimecoding/soil/pkg/proxmox/constant"

func (o *Tester) Populate() {
	c := o.Alfa().Hypervisor
	c.AddMachine("alfa", 100, "first")
	c.SetMachineStatus("alfa", 100, proxmoxConstant.RunningStatus)
	c.AddMachine("alfa", 101, "second")
	c.AddGuestNotInBackup(proxmoxConstant.MachineType, 101, "second")
	c.AddUpdatePending("alfa", "pve-docs", "9.2.3", "9.2.4")
}
