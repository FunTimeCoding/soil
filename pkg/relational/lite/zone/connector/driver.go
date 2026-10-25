package connector

import "database/sql/driver"

func (c *Connector) Driver() driver.Driver {
	return c.inner
}
