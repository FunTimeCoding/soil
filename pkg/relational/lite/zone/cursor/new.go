package cursor

import "database/sql/driver"

func New(r driver.Rows) *Cursor {
	return &Cursor{inner: r.(Inner)}
}
