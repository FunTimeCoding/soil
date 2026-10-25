package connector

import "database/sql/driver"

func New(
	inner driver.Driver,
	locator string,
) *Connector {
	return &Connector{inner: inner, locator: locator}
}
