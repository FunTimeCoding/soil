package lite

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func New(
	l *logger.Logger,
	path string,
) *gorm.DB {
	m, e := gorm.Open(
		&sqlite.Dialector{Conn: connection.New(l, path)},
		&gorm.Config{},
	)
	errors.PanicOnError(e)

	return m
}
