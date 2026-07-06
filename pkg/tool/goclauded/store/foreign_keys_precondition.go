package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"gorm.io/gorm"
	"log"
)

func foreignKeysPrecondition(d *gorm.DB) bool {
	checks := []foreignKeyCheck{
		{"event", "session_identifier", constant.SessionTable},
		{"completion", "session_identifier", constant.SessionTable},
		{constant.SummaryTable, "session_identifier", constant.SessionTable},
		{"label", "session_identifier", constant.SessionTable},
		{"pulse", "session_identifier", constant.SessionTable},
		{"event_metadata", "event_identifier", "event"},
	}

	for _, c := range checks {
		var count int64
		d.Raw(
			fmt.Sprintf(
				`SELECT COUNT(*) FROM %s
				WHERE %s IS NULL
				OR %s NOT IN (SELECT identifier FROM %s)`,
				c.child,
				c.column,
				c.column,
				c.parent,
			),
		).Scan(&count)

		if count > 0 {
			log.Printf(
				"foreign key precondition failed: %s has %d rows with invalid %s",
				c.child,
				count,
				c.column,
			)

			return false
		}
	}

	return true
}
