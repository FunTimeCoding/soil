package lite

import (
	"database/sql"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/relational/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func NewMemory() *gorm.DB {
	m, e := gorm.Open(
		sqlite.Open(
			join.Empty(constant.LiteMemory, constant.LiteMemoryParameters),
		),
		&gorm.Config{},
	)
	errors.PanicOnError(e)
	m.ConnPool.(*sql.DB).SetMaxOpenConns(1)

	return m
}
