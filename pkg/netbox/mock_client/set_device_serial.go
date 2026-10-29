package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/device"

func (c *Client) SetDeviceSerial(
	name string,
	serial string,
) (*device.Device, error) {
	return c.changeDevice("SetDeviceSerial", name, serial)
}
