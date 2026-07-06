package relational

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"log"
)

func (d *Database) IsOwner(
	user string,
	dbname string,
) bool {
	if !d.UserExists(user) {
		log.Panicf("user '%s' does not exist", user)
	}

	if !d.DatabaseExists(dbname) {
		log.Panicf("database '%s' does not exist", dbname)
	}

	var result bool
	errors.PanicOnError(
		d.QueryRow(
			`SELECT EXISTS(
                SELECT 1 FROM pg_database d 
                JOIN pg_roles r ON d.datdba = r.oid 
                WHERE d.datname = $1 AND r.rolname = $2
            )`,
			dbname,
			user,
		).Scan(&result),
	)

	return result
}
