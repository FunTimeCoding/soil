package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/foreign_key_check"
	"gorm.io/gorm"
	"log"
)

func foreignKeysPrecondition(d *gorm.DB) bool {
	checks := []*foreign_key_check.Check{
		foreign_key_check.New(
			"event",
			"session_identifier",
			constant.SessionTable,
		),
		foreign_key_check.New(
			"completion",
			"session_identifier",
			constant.SessionTable,
		),
		foreign_key_check.New(
			constant.SummaryTable,
			"session_identifier",
			constant.SessionTable,
		),
		foreign_key_check.New(
			constant.LabelTable,
			"session_identifier",
			constant.SessionTable,
		),
		foreign_key_check.New(
			constant.PulseTable,
			"session_identifier",
			constant.SessionTable,
		),
		foreign_key_check.New(
			"event_metadata",
			constant.EventIdentifierColumn,
			"event",
		),
	}

	for _, c := range checks {
		var count int64
		d.Raw(
			fmt.Sprintf(
				`SELECT COUNT(*) FROM %s
				WHERE %s IS NULL
				OR %s NOT IN (SELECT identifier FROM %s)`,
				c.Child,
				c.Column,
				c.Column,
				c.Parent,
			),
		).Scan(&count)

		if count > 0 {
			log.Printf(
				"foreign key precondition failed: %s has %d rows with invalid %s",
				c.Child,
				count,
				c.Column,
			)

			return false
		}
	}

	return true
}
