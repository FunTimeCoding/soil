package relational

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/jackc/pgx/v5"
)

func (d *Database) CreateUserOwner(
	user string,
	password string,
	dbname string,
) {
	d.validateInputs(user, password, dbname)
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
			"CREATE DATABASE %s OWNER %s",
			pgx.Identifier{dbname}.Sanitize(),
			pgx.Identifier{user}.Sanitize(),
		),
	)
	d.Commit(t)
}
