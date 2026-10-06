package nextcloud

func (c *Client) Status() error {
	return c.basic.Propfind()
}
