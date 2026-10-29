package relational

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"gorm.io/gorm"
)

func CloseMapper(m *gorm.DB) {
	d, e := m.DB()
	errors.PanicOnError(e)
	errors.PanicOnError(d.Close())
}
