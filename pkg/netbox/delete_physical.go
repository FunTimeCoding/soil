package netbox

func (c *Client) DeletePhysical(identifier int32) error {
	_, e := c.client.DcimAPI.DcimMacAddressesDestroy(
		c.context,
		identifier,
	).Execute()

	return e
}
