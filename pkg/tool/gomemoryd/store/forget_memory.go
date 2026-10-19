package store

import (
	"database/sql"
	"errors"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"time"
)

func (s *Store) ForgetMemory(
	identifier int64,
	source string,
) error {
	now := time.Now().UTC().Format(time.RFC3339)
	t, e := s.database.Begin()

	if e != nil {
		return e
	}

	defer rollback(t)
	_, e = t.Exec(
		`UPDATE memory SET is_active = 0, updated_at = ? WHERE identifier = ?`,
		now,
		identifier,
	)

	if e != nil {
		return e
	}

	var name, content, description string
	e = t.QueryRow(
		`SELECT name, content, description FROM memory WHERE identifier = ?`,
		identifier,
	).Scan(&name, &content, &description)

	if errors.Is(e, sql.ErrNoRows) {
		return not_found.New(constant.MemoryKind, identifier)
	}

	if e != nil {
		return e
	}

	_, e = t.Exec(
		`INSERT INTO memory_version (memory_identifier, name, content, description, changed_at, change_type, source)
		VALUES (?, ?, ?, ?, ?, 'forgotten', ?)`,
		identifier,
		name,
		content,
		description,
		now,
		source,
	)

	if e != nil {
		return e
	}

	return t.Commit()
}
