package netbox

import "github.com/funtimecoding/soil/pkg/netbox/refusal"

func (c *Client) DeleteCable(identifier int32) error {
	_, e := c.client.DcimAPI.DcimCablesDestroy(c.context, identifier).Execute()

	return refusal.Classify(e)
}
