package unit

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"gorm.io/gorm"
)

func closeMapper(m *gorm.DB) {
	inner, e := m.DB()
	errors.PanicOnError(e)
	errors.PanicOnError(inner.Close())
}
