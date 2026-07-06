package relational

import (
	"database/sql"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (d *Database) connect(locator string) {
	client, connectFail := pgxpool.New(d.context, locator)
	errors.PanicOnError(connectFail)
	d.client = client
	connection, connectionFail := sql.Open("pgx", locator)
	errors.PanicOnError(connectionFail)
	mapper, mapperFail := gorm.Open(
		postgres.New(postgres.Config{Conn: connection}),
		&gorm.Config{},
	)
	errors.PanicOnError(mapperFail)
	d.mapper = mapper
}
