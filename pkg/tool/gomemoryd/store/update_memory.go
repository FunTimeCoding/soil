package store

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
	"time"
)

func (s *Store) UpdateMemory(
	identifier int64,
	o *save_option.Option,
) error {
	now := time.Now().UTC().Format(time.RFC3339)
	t, e := s.database.Begin()

	if e != nil {
		return e
	}

	defer rollback(t)
	_, e = t.Exec(
		`UPDATE memory SET name = ?, content = ?, description = ?, updated_at = ? WHERE identifier = ?`,
		o.Name,
		o.Content,
		o.Description,
		now,
		identifier,
	)

	if e != nil {
		return e
	}

	_, e = t.Exec(
		`INSERT INTO memory_version (memory_identifier, name, content, description, changed_at, change_type, source)
		VALUES (?, ?, ?, ?, ?, 'updated', ?)`,
		identifier,
		o.Name,
		o.Content,
		o.Description,
		now,
		o.Source,
	)

	if e != nil {
		return e
	}

	_, e = t.Exec(
		`DELETE FROM memory_tag WHERE memory_identifier = ?`,
		identifier,
	)

	if e != nil {
		return e
	}

	for _, tag := range o.Tags {
		_, e = t.Exec(
			`INSERT INTO memory_tag (memory_identifier, tag) VALUES (?, ?)`,
			identifier,
			tag,
		)

		if e != nil {
			return e
		}
	}

	return t.Commit()
}
