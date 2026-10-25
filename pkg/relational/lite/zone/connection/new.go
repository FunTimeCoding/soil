package connection

import "database/sql/driver"

func New(c driver.Conn) *Connection {
	return &Connection{inner: c.(Inner)}
}
