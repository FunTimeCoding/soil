package store

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
	"time"
)

func (s *Store) CreateMemory(o *save_option.Option) (int64, error) {
	now := time.Now().UTC().Format(time.RFC3339)
	t, e := s.database.Begin()

	if e != nil {
		return 0, e
	}

	defer rollback(t)
	result, e := t.Exec(
		`INSERT INTO memory (name, content, description, type, created_at, updated_at, parent_identifier)
		VALUES (?, ?, ?, ?, ?, ?, ?)`,
		o.Name,
		o.Content,
		o.Description,
		o.Type,
		now,
		now,
		o.ParentIdentifier,
	)

	if e != nil {
		return 0, e
	}

	identifier, e := result.LastInsertId()

	if e != nil {
		return 0, e
	}

	_, e = t.Exec(
		`INSERT INTO memory_version (memory_identifier, name, content, description, changed_at, change_type, source)
		VALUES (?, ?, ?, ?, ?, 'created', ?)`,
		identifier,
		o.Name,
		o.Content,
		o.Description,
		now,
		o.Source,
	)

	if e != nil {
		return 0, e
	}

	for _, tag := range o.Tags {
		_, e = t.Exec(
			`INSERT INTO memory_tag (memory_identifier, tag) VALUES (?, ?)`,
			identifier,
			tag,
		)

		if e != nil {
			return 0, e
		}
	}

	return identifier, t.Commit()
}
