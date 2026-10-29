package netbox

import "github.com/funtimecoding/soil/pkg/netbox/refusal"

func (c *Client) DeleteInterface(identifier int32) error {
	_, e := c.client.DcimAPI.DcimInterfacesDestroy(
		c.context,
		identifier,
	).Execute()

	return refusal.Classify(e)
}
