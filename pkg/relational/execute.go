package relational

import "github.com/funtimecoding/soil/pkg/errors"

func (d *Database) Execute(
	sql string,
	a ...any,
) {
	_, e := d.client.Exec(d.context, sql, a...)
	errors.PanicOnError(e)
}
