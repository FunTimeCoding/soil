package mock_client

import "github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"

type Client struct {
	DeviceList          []*client.Device
	VirtualMachineList  []*client.VirtualMachine
	PhysicalAddressList []client.PhysicalAddressOwner
	Failure             error
	Closed              int
	Updated             map[string]client.UpdateDeviceRequest
	UpdatedMachines     map[string]client.UpdateVirtualMachineRequest
}
