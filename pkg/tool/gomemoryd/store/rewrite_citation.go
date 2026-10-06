package store

import (
	"database/sql"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
)

func rewriteCitation(
	t *sql.Tx,
	v *record.Version,
	source string,
	now string,
) error {
	_, e := t.Exec(
		`UPDATE memory SET content = ?, updated_at = ? WHERE identifier = ?`,
		v.Content,
		now,
		v.MemoryIdentifier,
	)

	if e != nil {
		return e
	}

	_, e = t.Exec(
		`INSERT INTO memory_version (memory_identifier, name, content, description, changed_at, change_type, source)
		VALUES (?, ?, ?, ?, ?, 'rewritten', ?)`,
		v.MemoryIdentifier,
		v.Name,
		v.Content,
		v.Description,
		now,
		source,
	)

	return e
}
