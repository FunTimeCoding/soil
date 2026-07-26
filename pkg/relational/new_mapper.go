package relational

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational/constant"
	"gorm.io/gorm"
)

func NewMapper(
	l *logger.Logger,
	locator string,
) *gorm.DB {
	l.Structured(constant.PostgresMessage)

	return openMapper(locator)
}
