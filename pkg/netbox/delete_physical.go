package netbox

import "github.com/funtimecoding/soil/pkg/netbox/refusal"

func (c *Client) DeletePhysical(identifier int32) error {
	_, e := c.client.DcimAPI.DcimMacAddressesDestroy(
		c.context,
		identifier,
	).Execute()

	return refusal.Classify(e)
}
