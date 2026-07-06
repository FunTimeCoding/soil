package netbox

import (
	"github.com/funtimecoding/soil/pkg/integers64"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/netbox-community/go-netbox/v4"
	"net"
)

// UpdateInterface Update existing interface by name and assign MAC address to it
// If MAC address does not exist, it will be created
func (c *Client) UpdateInterface(
	d *device.Device,
	name string,
	t netbox.InterfaceTypeValue,
	h net.HardwareAddr,
) (*network.Interface, error) {
	p, e := c.EnsurePhysical(h)

	if e != nil {
		return nil, e
	}

	i, f := c.DeviceInterfaceByName(d, name)

	if f != nil {
		return nil, f
	}

	var assigned bool

	if p.ObjectType == constant.InterfaceAddress {
		if integers64.To32(p.ObjectIdentifier) == i.Identifier {
			assigned = true
		}
	}

	if !assigned {
		_, g := c.AssignPhysicalToInterface(p, i)

		if g != nil {
			return nil, g
		}
	}

	v := netbox.NewBriefDeviceRequest()
	v.SetName(d.Name)
	q := netbox.NewWritableInterfaceRequest(
		netbox.BriefDeviceRequestAsBriefInterfaceRequestDevice(v),
		name,
		t,
	)
	q.SetPrimaryMacAddress(
		netbox.BriefMACAddressRequestAsInterfaceRequestPrimaryMacAddress(
			netbox.NewBriefMACAddressRequest(h.String()),
		),
	)
	j, h2 := c.DeviceInterfaceByName(d, name)

	if h2 != nil {
		return nil, h2
	}

	result, _, k := c.client.DcimAPI.DcimInterfacesUpdate(
		c.context,
		j.Identifier,
	).WritableInterfaceRequest(*q).Execute()

	if k != nil {
		return nil, k
	}

	return network.New(result), nil
}
