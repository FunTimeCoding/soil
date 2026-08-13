package netbox

func (c *Client) DeleteDevice(identifier int32) error {
	_, e := c.client.DcimAPI.DcimDevicesDestroy(c.context, identifier).Execute()

	return e
}
