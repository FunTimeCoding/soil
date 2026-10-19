package schema

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"gorm.io/gorm"
)

func RenameTable(
	d *gorm.DB,
	from string,
	to string,
) {
	if m := d.Migrator(); m.HasTable(from) {
		errors.PanicOnError(m.RenameTable(from, to))
	}
}
