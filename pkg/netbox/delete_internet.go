package netbox

import "github.com/funtimecoding/soil/pkg/netbox/refusal"

func (c *Client) DeleteInternet(identifier int32) error {
	_, e := c.client.IpamAPI.IpamIpAddressesDestroy(
		c.context,
		identifier,
	).Execute()

	return refusal.Classify(e)
}
