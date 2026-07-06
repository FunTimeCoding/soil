package relational

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/jackc/pgx/v5"
)

func (d *Database) Begin() pgx.Tx {
	result, e := d.client.Begin(d.context)
	errors.PanicOnError(e)

	return result
}
