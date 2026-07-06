package relational

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/jackc/pgx/v5"
)

func (d *Database) RowValues(r pgx.Rows) []any {
	result, e := r.Values()
	errors.PanicOnError(e)

	return result
}
