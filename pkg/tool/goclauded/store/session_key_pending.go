package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"gorm.io/gorm"
)

func sessionKeyPending(
	d *gorm.DB,
	table string,
) bool {
	if !columnExists(d, table, constant.Callsign) {
		return false
	}

	if !columnExists(d, table, "session_identifier") {
		return true
	}

	return unkeyedRows(d, table) > 0
}
