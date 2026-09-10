package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"gorm.io/gorm"
	"log"
	"time"
)

func migrateTimestamps(d *gorm.DB) {
	converted := 0
	skipped := 0

	for _, c := range timestampColumns() {
		var rows []timestampValue
		d.Raw(
			fmt.Sprintf(
				`SELECT identifier, CAST(%s AS TEXT) AS value FROM %s
				WHERE %s IS NOT NULL AND %s != '' AND %s NOT LIKE ?`,
				c.column,
				c.table,
				c.column,
				c.column,
				c.column,
			),
			fmt.Sprintf("%%%s", constant.UniversalSuffix),
		).Scan(&rows)

		for _, r := range rows {
			t, e := time.Parse(constant.TimestampLayout, r.Value)

			if e != nil {
				skipped++

				continue
			}

			d.Exec(
				fmt.Sprintf(
					"UPDATE %s SET %s = ? WHERE identifier = ?",
					c.table,
					c.column,
				),
				t.UTC().Format(constant.TimestampLayout),
				r.Identifier,
			)
			converted++
		}
	}

	if converted == 0 && skipped == 0 {
		return
	}

	log.Printf(
		"timestamp migration: %d converted to UTC, %d unparseable and left alone",
		converted,
		skipped,
	)
}
