package cursor

func (c *Cursor) ColumnTypeNullable(index int) (bool, bool) {
	return c.inner.ColumnTypeNullable(index)
}
