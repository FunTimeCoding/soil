package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/store/claim"
)

func (s *Store) Release(
	item string,
	owner string,
) error {
	e := s.database.Where(
		map[string]any{"item": item, "owner": owner},
	).Delete(claim.New("", "")).Error

	if e != nil {
		return fmt.Errorf("release %s for %s: %w", item, owner, e)
	}

	s.notify()

	return nil
}
