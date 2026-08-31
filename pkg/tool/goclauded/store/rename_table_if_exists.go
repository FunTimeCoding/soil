package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"gorm.io/gorm"
)

func renameTableIfExists(
	d *gorm.DB,
	from string,
	to string,
) {
	if !d.Migrator().HasTable(from) || d.Migrator().HasTable(to) {
		return
	}

	errors.PanicOnError(
		d.Exec(fmt.Sprintf("ALTER TABLE %s RENAME TO %s", from, to)).Error,
	)
}
