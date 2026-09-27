package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"gorm.io/gorm"
	"log"
)

func timestampsNormalized(d *gorm.DB) bool {
	for _, c := range constant.TimestampColumns {
		var count int64
		d.Raw(
			fmt.Sprintf(
				`SELECT COUNT(*) FROM %s
				WHERE %s IS NOT NULL AND %s != '' AND %s NOT LIKE ?`,
				c.Table,
				c.Name,
				c.Name,
				c.Name,
			),
			fmt.Sprintf("%%%s", constant.UniversalSuffix),
		).Scan(&count)

		if count > 0 {
			log.Printf(
				"timestamp precondition failed: %s.%s has %d values not stored in UTC",
				c.Table,
				c.Name,
				count,
			)

			return false
		}
	}

	return true
}
