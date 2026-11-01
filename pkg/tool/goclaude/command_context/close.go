package command_context

func (c *Context) Close() {
	c.client.Close()
	c.longClient.Close()
}
