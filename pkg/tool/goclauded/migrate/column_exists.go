package migrate

import "gorm.io/gorm"

func columnExists(
	d *gorm.DB,
	table string,
	column string,
) bool {
	var count int64
	d.Raw(
		"SELECT COUNT(*) FROM pragma_table_info(?) WHERE name = ?",
		table,
		column,
	).Scan(&count)

	return count > 0
}
