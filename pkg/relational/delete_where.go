package relational

import "github.com/funtimecoding/soil/pkg/errors"

func (d *Database) DeleteWhere(
	t any,
	query string,
	a ...any,
) {
	errors.PanicOnError(d.Mapper().Where(query, a...).Delete(t).Error)
}
