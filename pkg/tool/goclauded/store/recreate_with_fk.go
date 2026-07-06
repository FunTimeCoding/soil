package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"gorm.io/gorm"
)

func recreateWithFK(
	d *gorm.DB,
	table string,
	createNew string,
	columns string,
) {
	var hasFK int64
	d.Raw(
		fmt.Sprintf(
			"SELECT COUNT(*) FROM pragma_foreign_key_list('%s')",
			table,
		),
	).Scan(&hasFK)

	if hasFK > 0 {
		return
	}

	d.Exec("PRAGMA foreign_keys = OFF")
	d.Exec(createNew)

	if e := d.Exec(
		fmt.Sprintf(
			"INSERT INTO %s_new (%s) SELECT %s FROM %s",
			table,
			columns,
			columns,
			table,
		),
	).Error; e != nil {
		d.Exec(fmt.Sprintf("DROP TABLE %s_new", table))
		d.Exec("PRAGMA foreign_keys = ON")
		errors.PanicOnError(e)
	}

	d.Exec(fmt.Sprintf("DROP TABLE %s", table))
	d.Exec(
		fmt.Sprintf(
			"ALTER TABLE %s_new RENAME TO %s",
			table,
			table,
		),
	)
	d.Exec("PRAGMA foreign_keys = ON")
}
