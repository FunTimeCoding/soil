package statement

import (
	"context"
	"database/sql/driver"
	"github.com/funtimecoding/soil/pkg/relational/lite/zone/cursor"
)

func (s *Statement) QueryContext(
	x context.Context,
	v []driver.NamedValue,
) (driver.Rows, error) {
	r, e := s.inner.QueryContext(x, v)

	if e != nil {
		return nil, e
	}

	return cursor.New(r), nil
}
