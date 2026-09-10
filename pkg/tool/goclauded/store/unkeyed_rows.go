package store

import (
	"fmt"
	"gorm.io/gorm"
)

func unkeyedRows(
	d *gorm.DB,
	table string,
) int64 {
	var result int64
	d.Raw(
		fmt.Sprintf(
			`SELECT COUNT(*) FROM %s
			WHERE session_identifier IS NULL OR session_identifier = ''`,
			table,
		),
	).Scan(&result)

	return result
}
