package netbox

import "github.com/funtimecoding/soil/pkg/netbox/refusal"

func (c *Client) DeleteDevice(identifier int32) error {
	_, e := c.client.DcimAPI.DcimDevicesDestroy(c.context, identifier).Execute()

	return refusal.Classify(e)
}
