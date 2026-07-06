package relational

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"log"
)

func (d *Database) HasAllPrivileges(
	user string,
	dbname string,
) bool {
	if !d.UserExists(user) {
		log.Panicf("user '%s' does not exist", user)
	}

	if !d.DatabaseExists(dbname) {
		log.Panicf("database '%s' does not exist", dbname)
	}

	if d.IsOwner(user, dbname) {
		return true
	}

	var result bool
	errors.PanicOnError(
		d.QueryRow(
			`SELECT 
                CASE 
                    WHEN has_database_privilege($1, $2, 'CONNECT') AND
                         has_database_privilege($1, $2, 'CREATE') AND 
                         has_database_privilege($1, $2, 'TEMPORARY')
                    THEN true 
                    ELSE false 
                END`,
			user,
			dbname,
		).Scan(&result),
	)

	return result
}
