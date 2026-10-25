package lite

import (
	"database/sql"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/relational/constant"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/relational/lite/zone/connector"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func NewMemory() *gorm.DB {
	d := sql.OpenDB(
		connector.New(
			connection.Driver(),
			join.Empty(constant.LiteMemory, constant.LiteMemoryParameters),
		),
	)
	d.SetMaxOpenConns(1)
	m, e := gorm.Open(&sqlite.Dialector{Conn: d}, &gorm.Config{})
	errors.PanicOnError(e)

	return m
}
