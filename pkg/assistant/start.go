package assistant

func (c *Client) Start() {
	c.connection.Connect()
	c.connection.Subscribe("", c.dispatch)
	go c.connection.Read()
	go c.poll()
}
