package zone

import (
	"database/sql"
	"github.com/funtimecoding/soil/pkg/errors"
)

func names(
	t *sql.Tx,
	query string,
	v ...any,
) []string {
	r, e := t.Query(query, v...)
	errors.PanicOnError(e)
	defer errors.PanicClose(r)
	var result []string

	for r.Next() {
		var name string
		errors.PanicOnError(r.Scan(&name))
		result = append(result, name)
	}

	errors.PanicOnError(r.Err())

	return result
}
