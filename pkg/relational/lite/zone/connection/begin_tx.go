package connection

import (
	"context"
	"database/sql/driver"
)

func (c *Connection) BeginTx(
	x context.Context,
	o driver.TxOptions,
) (driver.Tx, error) {
	return c.inner.BeginTx(x, o)
}
