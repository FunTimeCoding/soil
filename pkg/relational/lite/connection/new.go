package connection

import (
	"database/sql"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational/constant"
	"github.com/funtimecoding/soil/pkg/relational/lite/zone"
	"github.com/funtimecoding/soil/pkg/relational/lite/zone/connector"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
)

func New(
	l *logger.Logger,
	path string,
) *sql.DB {
	l.Structured(constant.LiteMessage)
	system.MakeDirectory(filepath.Dir(path))
	result := sql.OpenDB(
		connector.New(Driver(), join.Empty(path, constant.LiteFileParameters)),
	)
	zone.Normalize(result)

	return result
}
