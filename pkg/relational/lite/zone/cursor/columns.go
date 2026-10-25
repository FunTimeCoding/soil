package cursor

func (c *Cursor) Columns() []string {
	return c.inner.Columns()
}
