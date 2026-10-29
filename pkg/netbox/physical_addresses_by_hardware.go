package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/physical_address"
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
	"net"
)

func (c *Client) PhysicalAddressesByHardware(a net.HardwareAddr) ([]*physical_address.Address, error) {
	result, _, e := c.client.DcimAPI.DcimMacAddressesList(c.context).MacAddress(
		[]string{a.String()},
	).Execute()

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return physical_address.NewSlice(result.Results), nil
}
