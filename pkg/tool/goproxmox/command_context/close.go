package command_context

func (c *Context) Close() {
	if c.client != nil {
		c.client.Close()
	}
}
