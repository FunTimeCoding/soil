package schema

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"gorm.io/gorm"
)

func RenameIndexes(
	d *gorm.DB,
	model any,
	renames map[string]string,
) {
	m := d.Migrator()

	for from, to := range renames {
		if m.HasIndex(model, from) {
			errors.PanicOnError(m.RenameIndex(model, from, to))
		}
	}
}
