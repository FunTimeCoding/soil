package relational

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/jackc/pgx/v5"
	"log"
)

func (d *Database) CreateUserGrant(
	user string,
	password string,
	dbname string,
) {
	d.validateInputs(user, password, dbname)

	if !d.DatabaseExists(dbname) {
		log.Panicf("database '%s' does not exist", dbname)
	}

	t := d.Begin()
	defer errors.PanicOnError(t.Rollback(d.context))
	d.Transact(
		t,
		fmt.Sprintf(
			"CREATE USER %s WITH PASSWORD '%s'",
			pgx.Identifier{user}.Sanitize(),
			password,
		),
	)
	d.Transact(
		t,
		fmt.Sprintf(
			"GRANT ALL PRIVILEGES ON DATABASE %s TO %s",
			pgx.Identifier{dbname}.Sanitize(),
			pgx.Identifier{user}.Sanitize(),
		),
	)
	d.Commit(t)
}
