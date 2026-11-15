package lite

import (
	"github.com/funtimecoding/soil/pkg/relational/constant"
	"gorm.io/gorm"
)

func ColumnExists(
	d *gorm.DB,
	table string,
	column string,
) bool {
	var count int64
	d.Raw(constant.LiteColumnExistsQuery, table, column).Scan(&count)

	return count > 0
}
