package cursor

func (c *Cursor) ColumnTypeDatabaseTypeName(index int) string {
	return c.inner.ColumnTypeDatabaseTypeName(index)
}
