package schema

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/relational/constant"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func DropColumns(
	d *gorm.DB,
	table string,
	columns ...string,
) {
	for _, c := range columns {
		if d.Migrator().HasColumn(table, c) {
			errors.PanicOnError(
				d.Exec(
					constant.DropColumn,
					clause.Table{Name: table},
					clause.Column{Name: c},
				).Error,
			)
		}
	}
}
