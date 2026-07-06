package relational

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/jackc/pgx/v5"
)

func (d *Database) Transact(
	t pgx.Tx,
	sql string,
	a ...any,
) {
	_, e := t.Exec(d.context, sql, a...)
	errors.PanicOnError(e)
}
