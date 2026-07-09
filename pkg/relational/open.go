package relational

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"gorm.io/gorm"
)

func Open(
	l *logger.Logger,
	locator string,
	litePath string,
) *gorm.DB {
	if locator != "" {
		return NewMapper(l, locator)
	}

	if litePath != "" {
		return lite.New(l, litePath)
	}

	panic("no database configured")
}
