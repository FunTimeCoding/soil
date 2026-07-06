package relational

import "github.com/funtimecoding/soil/pkg/errors"

func (d *Database) DatabaseOwner(dbname string) string {
	var result string
	errors.PanicOnError(
		d.QueryRow(
			`
        SELECT r.rolname 
        FROM pg_database d 
        JOIN pg_roles r ON d.datdba = r.oid 
        WHERE d.datname = $1
    `,
			dbname,
		).Scan(&result),
	)

	return result
}
