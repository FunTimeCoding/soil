package schema

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"gorm.io/gorm"
)

func DropTable(
	d *gorm.DB,
	name string,
) {
	errors.PanicOnError(d.Migrator().DropTable(name))
}
