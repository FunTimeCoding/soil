package service_tester

import "github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"

func (t *Tester) SeedInventory() {
	device := "198.51.100.8/24"
	t.Netbox.DeviceList = []*client.Device{
		{Identifier: 41, Name: "delta", PrimaryAddress: &device},
	}
	t.Netbox.VirtualMachineList = []*client.VirtualMachine{
		{Identifier: 2, Name: "echo"},
	}
}
