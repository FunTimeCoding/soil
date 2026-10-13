package command_context

func (c *Context) Close() {
	c.pool.Close()
	c.longPool.Close()
}
