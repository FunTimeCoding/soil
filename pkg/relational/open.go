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
		l.Structured(PostgresMessage)

		return NewMapper(locator)
	}

	if litePath != "" {
		l.Structured(LiteMessage)

		return lite.New(litePath)
	}

	panic("no database configured")
}
