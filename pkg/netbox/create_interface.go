package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateInterface(
	d *device.Device,
	name string,
	t netbox.InterfaceTypeValue,
) (*network.Interface, error) {
	v := netbox.NewBriefDeviceRequest()
	v.SetName(d.Name)
	q := netbox.NewWritableInterfaceRequest(
		netbox.BriefDeviceRequestAsBriefInterfaceRequestDevice(v),
		name,
		t,
	)
	result, _, e := c.client.DcimAPI.DcimInterfacesCreate(
		c.context,
	).WritableInterfaceRequest(*q).Execute()

	if e != nil {
		return nil, e
	}

	return network.New(result), nil
}
