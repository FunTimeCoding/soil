package connection

func (c *Connection) Close() error {
	return c.inner.Close()
}
