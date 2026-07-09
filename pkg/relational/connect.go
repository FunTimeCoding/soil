package relational

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (d *Database) connect(locator string) {
	client, e := pgxpool.New(d.context, locator)
	errors.PanicOnError(e)
	d.client = client
	d.mapper = NewMapper(locator)
}
