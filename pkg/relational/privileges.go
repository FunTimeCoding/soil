package relational

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"log"
)

func (d *Database) Privileges(
	user string,
	dbname string,
) map[string]bool {
	if !d.UserExists(user) {
		log.Panicf("user '%s' does not exist", user)
	}

	if !d.DatabaseExists(dbname) {
		log.Panicf("database '%s' does not exist", dbname)
	}

	var connect, create, temporary bool
	errors.PanicOnError(
		d.QueryRow(
			`SELECT 
                has_database_privilege($1, $2, 'CONNECT'),
                has_database_privilege($1, $2, 'CREATE'),
                has_database_privilege($1, $2, 'TEMPORARY')`,
			user,
			dbname,
		).Scan(&connect, &create, &temporary),
	)

	return map[string]bool{
		"CONNECT":   connect,
		"CREATE":    create,
		"TEMPORARY": temporary,
		"IS_OWNER":  d.IsOwner(user, dbname),
	}
}
