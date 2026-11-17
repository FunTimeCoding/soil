package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) SetInterfaceEthernetPower(
	d *device.Device,
	name string,
	mode string,
	kind string,
) (*network.Interface, error) {
	q := netbox.NewPatchedWritableInterfaceRequest()
	q.SetPoeModeNil()
	q.SetPoeTypeNil()

	if mode != "" {
		m, e := netbox.NewInterfaceTemplateRequestPoeModeFromValue(mode)

		if e != nil {
			return nil, validation.New("ethernet power mode %s: %v", mode, e)
		}

		q.SetPoeMode(*m)
	}

	if kind != "" {
		k, e := netbox.NewInterfaceTemplateRequestPoeTypeFromValue(kind)

		if e != nil {
			return nil, validation.New("ethernet power type %s: %v", kind, e)
		}

		q.SetPoeType(*k)
	}

	i, e := c.DeviceInterface(d, name)

	if e != nil {
		return nil, e
	}

	result, _, f := c.client.DcimAPI.DcimInterfacesPartialUpdate(
		c.context,
		i.Identifier,
	).PatchedWritableInterfaceRequest(*q).Execute()

	if f != nil {
		return nil, refusal.Classify(f)
	}

	return network.New(result), nil
}
