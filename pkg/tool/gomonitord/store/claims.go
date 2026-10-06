package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/store/claim"
)

func (s *Store) Claims() ([]*claim.Claim, error) {
	var result []*claim.Claim

	if e := s.database.Order("item, owner").Find(&result).Error; e != nil {
		return nil, fmt.Errorf("list claims: %w", e)
	}

	return result, nil
}
