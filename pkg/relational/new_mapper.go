package relational

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"gorm.io/gorm"
)

func NewMapper(
	l *logger.Logger,
	locator string,
) *gorm.DB {
	l.Structured(PostgresMessage)

	return openMapper(locator)
}
