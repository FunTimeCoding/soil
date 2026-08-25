package netbox

func (c *Client) DeleteInternet(identifier int32) error {
	_, e := c.client.IpamAPI.IpamIpAddressesDestroy(
		c.context,
		identifier,
	).Execute()

	return e
}
