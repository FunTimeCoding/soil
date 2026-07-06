package relational

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"log"
)

func (d *Database) HasCreatePrivilege(
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
			`SELECT has_database_privilege($1, $2, 'CREATE')`,
			user,
			dbname,
		).Scan(&result),
	)

	return result
}
