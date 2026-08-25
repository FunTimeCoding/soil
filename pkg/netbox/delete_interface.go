package netbox

func (c *Client) DeleteInterface(identifier int32) error {
	_, e := c.client.DcimAPI.DcimInterfacesDestroy(
		c.context,
		identifier,
	).Execute()

	return e
}
