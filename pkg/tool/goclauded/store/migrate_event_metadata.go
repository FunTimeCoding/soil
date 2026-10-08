package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/record"
	"gorm.io/gorm"
)

func migrateEventMetadata(d *gorm.DB) {
	if columnExists(d, "event", constant.Body) {
		var events []*record.LegacyEvent
		d.Raw("SELECT identifier, kind, scope, body FROM event").Scan(&events)

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
