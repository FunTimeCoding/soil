package lite

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational/lite/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"path/filepath"
)

func New(
	l *logger.Logger,
	path string,
) *gorm.DB {
	l.Structured(constant.LiteMessage)
	system.MakeDirectory(filepath.Dir(path))
	m, e := gorm.Open(
		sqlite.Open(join.Empty(path, constant.FileParameters)),
		&gorm.Config{},
	)
	errors.PanicOnError(e)

	return m
}
