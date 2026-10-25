package statement

import "database/sql/driver"

func New(s driver.Stmt) *Statement {
	return &Statement{inner: s.(Inner)}
}
