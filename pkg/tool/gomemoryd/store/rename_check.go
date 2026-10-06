package store

import (
	"database/sql"
	"github.com/funtimecoding/soil/pkg/errors/validation"
)

func renameCheck(
	t *sql.Tx,
	identifier int64,
	scope string,
	name string,
) error {
	var count int
	e := t.QueryRow(
		`SELECT COUNT(*) FROM memory
		WHERE scope = ? AND name = ? AND is_active = 1 AND identifier != ?`,
		scope,
		name,
		identifier,
	).Scan(&count)

	if e != nil {
		return e
	}

	if count > 0 {
		return validation.New(
			"memory name already exists in its scope: %s",
			name,
		)
	}

	return nil
}
