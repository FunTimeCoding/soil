package connection

import (
	"context"
	"database/sql/driver"
)

func (c *Connection) Prepare(query string) (driver.Stmt, error) {
	return c.PrepareContext(context.Background(), query)
}
