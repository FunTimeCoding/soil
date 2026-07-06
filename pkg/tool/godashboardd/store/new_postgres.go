package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newPostgres(dsn string) *Store {
	database, e := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	errors.PanicOnError(e)
	errors.PanicOnError(database.AutoMigrate(NewClick()))

	return &Store{mapper: database}
}
