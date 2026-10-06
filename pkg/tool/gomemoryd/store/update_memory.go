package store

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
	"time"
)

func (s *Store) UpdateMemory(
	identifier int64,
	o *save_option.Option,
) ([]int64, error) {
	now := time.Now().UTC().Format(time.RFC3339)
	t, e := s.database.Begin()

	if e != nil {
		return nil, e
	}

	defer rollback(t)
	var name, scope string
	e = t.QueryRow(
		`SELECT name, scope FROM memory WHERE identifier = ?`,
		identifier,
	).Scan(&name, &scope)

	if e != nil {
		return nil, e
	}

	renamed := o.Name != name

	if renamed {
		if e = renameCheck(t, identifier, scope, o.Name); e != nil {
			return nil, e
		}
	}

	if e = writeMemory(t, identifier, o, now); e != nil {
		return nil, e
	}

	var rewritten []int64

	if renamed {
		rewritten, e = rewriteCitations(
			t,
			citation(scope, name),
			citation(scope, o.Name),
			o.Source,
			now,
		)

		if e != nil {
			return nil, e
		}
	}

	return rewritten, t.Commit()
}
