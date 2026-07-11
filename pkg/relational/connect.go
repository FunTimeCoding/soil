package relational

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (d *Database) connect(locator string) {
	configuration, e := pgxpool.ParseConfig(locator)
	errors.PanicOnError(e)
	configuration.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeExec
	client, f := pgxpool.NewWithConfig(d.context, configuration)
	errors.PanicOnError(f)
	d.client = client
	d.mapper = openMapper(locator)
}
