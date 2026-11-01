package assistant

func (c *Client) Start() {
	c.hold.Start()
}
