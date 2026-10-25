package connection

import (
	"context"
	"database/sql/driver"
	"github.com/funtimecoding/soil/pkg/relational/lite/zone/cursor"
)

func (c *Connection) QueryContext(
	x context.Context,
	query string,
	v []driver.NamedValue,
) (driver.Rows, error) {
	r, e := c.inner.QueryContext(x, query, v)

	if e != nil {
		return nil, e
	}

	return cursor.New(r), nil
}
