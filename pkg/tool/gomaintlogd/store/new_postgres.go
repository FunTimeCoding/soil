package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newPostgres(dsn string) *Store {
	database, e := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	errors.PanicOnError(e)
	errors.PanicOnError(database.AutoMigrate(entry.New()))

	return &Store{database: database}
}
