package store

import "gorm.io/gorm"

func insertEventMetadata(
	d *gorm.DB,
	identifier uint,
	key string,
	value string,
) {
	d.Exec(
		"INSERT OR IGNORE INTO event_metadata (event_identifier, key, value) VALUES (?, ?, ?)",
		identifier,
		key,
		value,
	)
}
