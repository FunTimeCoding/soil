package store

import (
	"github.com/funtimecoding/soil/pkg/relational/schema"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"gorm.io/gorm"
)

func migrateColumns(d *gorm.DB) {
	schema.RenameColumns(
		d,
		constant.SessionTable,
		map[string]string{
			"cwd":        "work_directory",
			"created_at": "started_at",
			"updated_at": "last_active_at",
		},
	)
	schema.RenameColumns(
		d,
		"event",
		map[string]string{
			constant.Target:      "scope",
			constant.SessionName: "actor",
		},
	)
	schema.DropColumns(
		d,
		constant.SessionTable,
		"needs_roster",
		"needs_reannounce",
		"listening",
	)
	schema.DropColumns(d, constant.MessageTable, "read")
	schema.DropTable(d, "snapshots")
	schema.RenameTable(d, "rate_snapshots", constant.RateSnapshotTable)
	schema.RenameTable(d, "fable_snapshots", constant.FableSnapshotTable)
}
