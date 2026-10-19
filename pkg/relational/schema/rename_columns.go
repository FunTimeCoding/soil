package schema

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"gorm.io/gorm"
)

func RenameColumns(
	d *gorm.DB,
	model any,
	renames map[string]string,
) {
	m := d.Migrator()

	for from, to := range renames {
		if m.HasColumn(model, from) {
			errors.PanicOnError(m.RenameColumn(model, from, to))
		}
	}
}
