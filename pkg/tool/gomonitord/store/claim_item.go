package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/store/claim"
)

func (s *Store) ClaimItem(
	item string,
	owner string,
) (*claim.Claim, error) {
	result := claim.New(item, owner)
	e := s.database.Where(
		map[string]any{"item": item, "owner": owner},
	).FirstOrCreate(result).Error

	if e != nil {
		return nil, fmt.Errorf("claim %s for %s: %w", item, owner, e)
	}

	s.notify()

	return result, nil
}
