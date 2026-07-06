package relational

import "github.com/funtimecoding/soil/pkg/errors"

func (d *Database) DatabaseExists(dbname string) bool {
	var result bool
	errors.PanicOnError(
		d.QueryRow(
			"SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)",
			dbname,
		).Scan(&result),
	)

	return result
}
