package cursor

import "reflect"

func (c *Cursor) ColumnTypeScanType(index int) reflect.Type {
	return c.inner.ColumnTypeScanType(index)
}
