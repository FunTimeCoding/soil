package connection

import (
	"context"
	"database/sql/driver"
)

func (c *Connection) Begin() (driver.Tx, error) {
	return c.inner.BeginTx(context.Background(), driver.TxOptions{})
}
