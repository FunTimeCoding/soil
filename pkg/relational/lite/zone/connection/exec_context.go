package connection

import (
	"context"
	"database/sql/driver"
)

func (c *Connection) ExecContext(
	x context.Context,
	query string,
	v []driver.NamedValue,
) (driver.Result, error) {
	return c.inner.ExecContext(x, query, v)
}
