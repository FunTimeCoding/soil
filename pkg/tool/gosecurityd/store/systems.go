package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/system"
)

func (s *Store) Systems() ([]system.System, error) {
	var result []system.System

	if e := s.database.Order("name").Find(&result).Error; e != nil {
		return nil, fmt.Errorf("list systems: %w", e)
	}

	return result, nil
}
