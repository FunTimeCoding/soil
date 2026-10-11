package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"gorm.io/gorm"
	"testing"
)

func columnCount(
	t *testing.T,
	d *gorm.DB,
	table string,
	column string,
) int64 {
	t.Helper()
	var result int64
	assert.FatalOnError(
		t,
		d.Raw(
			"SELECT COUNT(*) FROM pragma_table_info(?) WHERE name = ?",
			table,
			column,
		).Scan(&result).Error,
	)

	return result
}
