package relational

import "github.com/funtimecoding/soil/pkg/errors"

func (d *Database) UserDatabases(u string) []string {
	r := d.Query(
		`
        SELECT d.datname 
        FROM pg_database d 
        JOIN pg_roles r ON d.datdba = r.oid 
        WHERE r.rolname = $1 
        AND d.datistemplate = false
    `,
		u,
	)
	defer r.Close()
	var result []string

	for r.Next() {
		var name string
		errors.PanicOnError(r.Scan(&name))
		result = append(result, name)
	}

	errors.PanicOnError(r.Err())

	return result
}
