package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/installed"
)

func (s *Store) Packages(system uint) ([]installed.Installed, error) {
	var result []installed.Installed

	if e := s.database.Where(
		map[string]any{"system_identifier": system},
	).Order("name").Find(&result).Error; e != nil {
		return nil, fmt.Errorf("list packages: %w", e)
	}

	return result, nil
}
