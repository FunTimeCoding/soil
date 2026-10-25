package connector

import (
	"context"
	"database/sql/driver"
	"github.com/funtimecoding/soil/pkg/relational/lite/zone/connection"
)

func (c *Connector) Connect(context.Context) (driver.Conn, error) {
	o, e := c.inner.Open(c.locator)

	if e != nil {
		return nil, e
	}

	return connection.New(o), nil
}
