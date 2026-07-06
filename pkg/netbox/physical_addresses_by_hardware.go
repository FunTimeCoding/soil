package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/physical_address"
	"net"
)

func (c *Client) PhysicalAddressesByHardware(a net.HardwareAddr) ([]*physical_address.Address, error) {
	result, _, e := c.client.DcimAPI.DcimMacAddressesList(
		c.context,
	).MacAddress([]string{a.String()}).Execute()

	if e != nil {
		return nil, e
	}

	return physical_address.NewSlice(result.Results), nil
}
