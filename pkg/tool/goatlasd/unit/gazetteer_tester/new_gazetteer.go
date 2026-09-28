package gazetteer_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/gazetteer"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func NewGazetteer() *gazetteer.Gazetteer {
	device := "198.51.100.8/24"
	deviceKind := "dcim.device"
	deviceIdentifier := int32(41)
	deviceName := "delta"
	machineKind := "virtualization.virtualmachine"
	machineIdentifier := int32(2)
	machineName := "echo"
	machineInterface := "ens3"

	return gazetteer.New(
		[]*client.Device{
			{Identifier: 41, Name: "delta", PrimaryAddress: &device},
		},
		[]*client.VirtualMachine{{Identifier: 2, Name: "echo"}},
		[]client.PhysicalAddressOwner{
			{
				Address:          "02:aa:bb:cc:dd:08",
				ObjectKind:       &deviceKind,
				ObjectIdentifier: &deviceIdentifier,
				ObjectName:       &deviceName,
			},
			{
				Address:          "02:aa:bb:cc:dd:04",
				ObjectKind:       &machineKind,
				ObjectIdentifier: &machineIdentifier,
				ObjectName:       &machineName,
				Interface:        &machineInterface,
			},
		},
	)
}
