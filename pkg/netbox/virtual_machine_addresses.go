package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/internet_address"
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
)

func (c *Client) VirtualMachineAddresses(machine string) ([]*internet_address.Address, error) {
	result, _, e := c.client.IpamAPI.IpamIpAddressesList(c.context).VirtualMachine(
		[]string{machine},
	).Execute()

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return internet_address.NewSlice(result.Results), nil
}
