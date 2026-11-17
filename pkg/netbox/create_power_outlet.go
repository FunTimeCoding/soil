package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/power_outlet"
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreatePowerOutlet(
	d *device.Device,
	name string,
) (*power_outlet.Outlet, error) {
	v := netbox.NewBriefDeviceRequest()
	v.SetName(d.Name)
	q := netbox.NewWritablePowerOutletRequest(
		netbox.BriefDeviceRequestAsBriefInterfaceRequestDevice(v),
		name,
	)
	result, _, e := c.client.DcimAPI.DcimPowerOutletsCreate(c.context).WritablePowerOutletRequest(
		*q,
	).Execute()

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return power_outlet.New(result), nil
}
