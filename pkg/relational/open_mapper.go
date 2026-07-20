package relational

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func openMapper(locator string) *gorm.DB {
	c, e := pgx.ParseConfig(locator)
	errors.PanicOnError(e)
	c.DefaultQueryExecMode = pgx.QueryExecModeExec
	d := stdlib.OpenDB(*c)
	d.SetMaxOpenConns(MaxOpenConnections)
	d.SetMaxIdleConns(MaxIdleConnections)
	m, f := gorm.Open(
		postgres.New(postgres.Config{Conn: d}),
		&gorm.Config{},
	)
	errors.PanicOnError(f)

	return m
}
