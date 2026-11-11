package migrate

import "gorm.io/gorm"

func migrateEventMetadataConstraint(d *gorm.DB) {
	recreateWithFK(
		d,
		"event_metadata",
		`CREATE TABLE event_metadata_new (
			event_identifier INTEGER NOT NULL
				REFERENCES event(identifier) ON DELETE CASCADE,
			key   TEXT NOT NULL,
			value TEXT NOT NULL,
			PRIMARY KEY (event_identifier, key)
		)`,
		"event_identifier, key, value",
	)
}
