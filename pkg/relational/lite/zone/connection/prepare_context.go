package connection

import (
	"context"
	"database/sql/driver"
	"github.com/funtimecoding/soil/pkg/relational/lite/zone/statement"
)

func (c *Connection) PrepareContext(
	x context.Context,
	query string,
) (driver.Stmt, error) {
	s, e := c.inner.PrepareContext(x, query)

	if e != nil {
		return nil, e
	}

	return statement.New(s), nil
}
