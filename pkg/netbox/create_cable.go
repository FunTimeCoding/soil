package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/cable"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateCable(
	a *network.Interface,
	b *network.Interface,
) (*cable.Cable, error) {
	q := netbox.NewWritableCableRequest()
	q.SetATerminations(
		[]netbox.GenericObjectRequest{
			*netbox.NewGenericObjectRequest(
				constant.InterfaceAddress,
				a.Identifier,
			),
		},
	)
	q.SetBTerminations(
		[]netbox.GenericObjectRequest{
			*netbox.NewGenericObjectRequest(
				constant.InterfaceAddress,
				b.Identifier,
			),
		},
	)
	q.SetStatus(netbox.CABLESTATUSVALUE_CONNECTED)
	result, _, e := c.client.DcimAPI.DcimCablesCreate(c.context).WritableCableRequest(
		*q,
	).Execute()

	if e != nil {
		return nil, e
	}

	return cable.New(result), nil
}
