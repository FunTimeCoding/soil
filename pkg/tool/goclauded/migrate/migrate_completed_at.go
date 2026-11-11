package migrate

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"gorm.io/gorm"
	"log"
)

func migrateCompletedAt(d *gorm.DB) {
	if !columnExists(d, constant.CompletionTable, constant.CompletedAtColumn) {
		return
	}

	result := d.Exec(
		`UPDATE completion SET completed_at = created_at
		WHERE completed_at IS NULL`,
	)

	if result.Error != nil || result.RowsAffected == 0 {
		return
	}

	log.Printf(
		"completion stamp migration: %d completions stamped from their creation time",
		result.RowsAffected,
	)
}
