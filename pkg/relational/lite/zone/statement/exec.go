package statement

import (
	"context"
	"database/sql/driver"
)

func (s *Statement) Exec(v []driver.Value) (driver.Result, error) {
	return s.ExecContext(context.Background(), named(v))
}
