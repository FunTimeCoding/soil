package cursor

func (c *Cursor) ColumnTypeLength(index int) (int64, bool) {
	return c.inner.ColumnTypeLength(index)
}
