package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/provision/model/run"
	"gorm.io/gorm"
)

func New(
	m *gorm.DB,
	tableName string,
) *Store {
	errors.PanicOnError(m.Table(tableName).AutoMigrate(run.New()))

	return &Store{mapper: m, tableName: tableName}
}
