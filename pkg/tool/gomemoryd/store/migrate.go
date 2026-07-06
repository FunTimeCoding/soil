package store

import (
	"database/sql"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
)

func migrate(d *sql.DB) {
	addColumnIfMissing(
		d,
		"memory",
		constant.MemoryName,
		"TEXT NOT NULL DEFAULT ''",
	)
	addColumnIfMissing(
		d,
		"memory_version",
		constant.MemoryName,
		"TEXT NOT NULL DEFAULT ''",
	)
	addColumnIfMissing(
		d,
		"memory_version",
		constant.Source,
		"TEXT NOT NULL DEFAULT ''",
	)
	addColumnIfMissing(
		d,
		"memory",
		"parent_identifier",
		"INTEGER REFERENCES memory(identifier)",
	)
}
