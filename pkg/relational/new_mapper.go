package relational

import (
	"database/sql"
	"github.com/funtimecoding/soil/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewMapper(locator string) *gorm.DB {
	connection, e := sql.Open("pgx", locator)
	errors.PanicOnError(e)
	mapper, f := gorm.Open(
		postgres.New(postgres.Config{Conn: connection}),
		&gorm.Config{},
	)
	errors.PanicOnError(f)

	return mapper
}
