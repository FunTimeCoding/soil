package store

import (
	"gorm.io/gorm"
	"log"
)

func migrateConsumedAt(d *gorm.DB) {
	if !columnExists(d, "queue", "consumed_at") {
		return
	}

	result := d.Exec(
		`UPDATE queue SET consumed_at = created_at
		WHERE consumed = 1
		AND (consumed_at IS NULL OR consumed_at = '')`,
	)

	if result.Error != nil || result.RowsAffected == 0 {
		return
	}

	log.Printf(
		"delivery stamp migration: %d consumed rows stamped from their creation time",
		result.RowsAffected,
	)
}
