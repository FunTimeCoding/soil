package statement

import (
	"context"
	"database/sql/driver"
)

func (s *Statement) ExecContext(
	x context.Context,
	v []driver.NamedValue,
) (driver.Result, error) {
	return s.inner.ExecContext(x, v)
}
