package relational_tester

import (
	"database/sql"
	"github.com/funtimecoding/soil/pkg/errors"
)

func TextColumn(
	d *sql.DB,
	query string,
	v ...any,
) []string {
	r, e := d.Query(query, v...)
	errors.PanicOnError(e)
	defer errors.PanicClose(r)
	var result []string

	for r.Next() {
		var s string
		errors.PanicOnError(r.Scan(&s))
		result = append(result, s)
	}

	errors.PanicOnError(r.Err())

	return result
}
