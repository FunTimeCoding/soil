package relational

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func openMapper(locator string) *gorm.DB {
	configuration, e := pgx.ParseConfig(locator)
	errors.PanicOnError(e)
	configuration.DefaultQueryExecMode = pgx.QueryExecModeExec
	mapper, f := gorm.Open(
		postgres.New(
			postgres.Config{Conn: stdlib.OpenDB(*configuration)},
		),
		&gorm.Config{},
	)
	errors.PanicOnError(f)

	return mapper
}
