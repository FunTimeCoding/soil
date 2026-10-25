package statement

import (
	"context"
	"database/sql/driver"
)

func (s *Statement) Query(v []driver.Value) (driver.Rows, error) {
	return s.QueryContext(context.Background(), named(v))
}
