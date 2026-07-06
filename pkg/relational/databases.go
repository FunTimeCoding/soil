package relational

import "github.com/funtimecoding/soil/pkg/errors"

func (d *Database) Databases() []string {
	r := d.Query(
		`SELECT datname 
         FROM pg_database 
         WHERE datistemplate = false 
         ORDER BY datname`,
	)
	defer r.Close()
	var result []string

	for r.Next() {
		var dbname string
		errors.PanicOnError(r.Scan(&dbname))
		result = append(result, dbname)
	}

	errors.PanicOnError(r.Err())

	return result
}
