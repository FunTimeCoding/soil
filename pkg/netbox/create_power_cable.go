package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/cable"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/power_outlet"
	"github.com/funtimecoding/soil/pkg/netbox/power_port"
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreatePowerCable(
	o *power_outlet.Outlet,
	p *power_port.Port,
) (*cable.Cable, error) {
	q := netbox.NewWritableCableRequest()
	q.SetATerminations(
		[]netbox.GenericObjectRequest{
			*netbox.NewGenericObjectRequest(
				constant.PowerOutletAddress,
				o.Identifier,
			),
		},
	)
	q.SetBTerminations(
		[]netbox.GenericObjectRequest{
			*netbox.NewGenericObjectRequest(
				constant.PowerPortAddress,
				p.Identifier,
			),
		},
	)
	q.SetStatus(netbox.CABLESTATUSVALUE_CONNECTED)
	result, _, e := c.client.DcimAPI.DcimCablesCreate(c.context).WritableCableRequest(
		*q,
	).Execute()

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return cable.New(result), nil
}
