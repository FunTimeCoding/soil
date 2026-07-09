package connection

import (
	"database/sql"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational/lite/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	_ "github.com/glebarez/go-sqlite"
	"path/filepath"
)

func New(
	l *logger.Logger,
	path string,
) *sql.DB {
	l.Structured(constant.LiteMessage)
	system.MakeDirectory(filepath.Dir(path))
	database, e := sql.Open(
		constant.DriverName,
		join.Empty(path, constant.FileParameters),
	)
	errors.PanicOnError(e)

	return database
}
