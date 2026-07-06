package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) createInterfaceWriteable(q *netbox.WritableInterfaceRequest) (*network.Interface, error) {
	result, _, e := c.client.DcimAPI.DcimInterfacesCreate(
		c.context,
	).WritableInterfaceRequest(*q).Execute()

	if e != nil {
		return nil, e
	}

	return network.New(result), nil
}
