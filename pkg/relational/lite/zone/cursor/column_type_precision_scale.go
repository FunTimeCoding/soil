package cursor

func (c *Cursor) ColumnTypePrecisionScale(index int) (int64, int64, bool) {
	return c.inner.ColumnTypePrecisionScale(index)
}
