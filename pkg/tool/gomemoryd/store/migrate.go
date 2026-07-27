package store

import (
	"database/sql"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
)

func migrate(d *sql.DB) {
	addColumnIfMissing(
		d,
		constant.MemoryTable,
		constant.MemoryName,
		"TEXT NOT NULL DEFAULT ''",
	)
	addColumnIfMissing(
		d,
		constant.VersionTable,
		constant.MemoryName,
		"TEXT NOT NULL DEFAULT ''",
	)
	addColumnIfMissing(
		d,
		constant.VersionTable,
		constant.Source,
		"TEXT NOT NULL DEFAULT ''",
	)
	addColumnIfMissing(
		d,
		constant.MemoryTable,
		"parent_identifier",
		"INTEGER REFERENCES memory(identifier)",
	)
	addColumnIfMissing(
		d,
		constant.MemoryTable,
		constant.Scope,
		"TEXT NOT NULL DEFAULT ''",
	)
	addColumnIfMissing(
		d,
		constant.MemoryTable,
		"provenance_file",
		"TEXT NOT NULL DEFAULT ''",
	)
	addColumnIfMissing(
		d,
		constant.MemoryTable,
		"provenance_anchor",
		"TEXT NOT NULL DEFAULT ''",
	)
	addColumnIfMissing(
		d,
		constant.MemoryTable,
		"provenance_hash",
		"TEXT NOT NULL DEFAULT ''",
	)
	addColumnIfMissing(
		d,
		constant.MemoryTable,
		"ordinal",
		"INTEGER NOT NULL DEFAULT 0",
	)
}
