package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"gorm.io/gorm"
)

func migrateColumns(d *gorm.DB) {
	renameIfExists(d, constant.SessionTable, "cwd", "work_directory")
	renameIfExists(d, constant.SessionTable, "created_at", "started_at")
	renameIfExists(d, constant.SessionTable, "updated_at", "last_active_at")
	renameIfExists(d, "event", constant.Target, "scope")
	renameIfExists(d, "event", constant.SessionName, "actor")
	dropIfExists(d, constant.SessionTable, "needs_roster")
	dropIfExists(d, constant.SessionTable, "needs_reannounce")
}
