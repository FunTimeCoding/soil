package connector

import "database/sql/driver"

type Connector struct {
	inner   driver.Driver
	locator string
}
