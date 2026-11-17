package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/power_port"
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreatePowerPort(
	d *device.Device,
	name string,
) (*power_port.Port, error) {
	v := netbox.NewBriefDeviceRequest()
	v.SetName(d.Name)
	q := netbox.NewWritablePowerPortRequest(
		netbox.BriefDeviceRequestAsBriefInterfaceRequestDevice(v),
		name,
	)
	result, _, e := c.client.DcimAPI.DcimPowerPortsCreate(c.context).WritablePowerPortRequest(
		*q,
	).Execute()

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return power_port.New(result), nil
}
