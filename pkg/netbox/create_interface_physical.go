package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/netbox-community/go-netbox/v4"
	"net"
)

func (c *Client) CreateInterfacePhysical(
	d *device.Device,
	name string,
	t netbox.InterfaceTypeValue,
	h net.HardwareAddr,
) (*network.Interface, error) {
	p, e := c.EnsurePhysical(h)

	if e != nil {
		return nil, e
	}

	v := netbox.NewBriefDeviceRequest()
	v.SetName(d.Name)
	i, f := c.createInterfaceWriteable(
		netbox.NewWritableInterfaceRequest(
			netbox.BriefDeviceRequestAsBriefInterfaceRequestDevice(v),
			name,
			t,
		),
	)

	if f != nil {
		return nil, f
	}

	_, g := c.AssignPhysicalToInterface(p, i)

	if g != nil {
		return nil, g
	}

	return c.UpdateInterface(d, name, t, h)
}
