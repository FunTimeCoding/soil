package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"gorm.io/gorm"
)

func migrateEventMetadata(d *gorm.DB) {
	if columnExists(d, "event", constant.Body) {
		type legacyEvent struct {
			Identifier uint   `gorm:"column:identifier"`
			Kind       string `gorm:"column:kind"`
			Scope      string `gorm:"column:scope"`
			Body       string `gorm:"column:body"`
		}
		var events []legacyEvent
		d.Raw(
			"SELECT identifier, kind, scope, body FROM event",
		).Scan(&events)

		for _, e := range events {
			metadata := legacyMetadata(e.Kind, e.Scope, e.Body)

			for key, value := range metadata {
				d.Exec(
					"INSERT OR IGNORE INTO event_metadata (event_identifier, key, value) VALUES (?, ?, ?)",
					e.Identifier,
					key,
					value,
				)
			}
		}

		dropIfExists(d, "event", constant.Body)
	}

	dropIfExists(d, "event", "scope")
}
