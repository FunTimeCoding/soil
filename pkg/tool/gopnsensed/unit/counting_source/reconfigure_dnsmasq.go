package counting_source

func (c *Client) ReconfigureDnsmasq() error {
	c.reconfigured++

	return c.OpnsenseSource.ReconfigureDnsmasq()
}
