package cursor

func (c *Cursor) Close() error {
	return c.inner.Close()
}
