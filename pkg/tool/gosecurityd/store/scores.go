package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/score"
)

func (s *Store) Scores() ([]score.Score, error) {
	var result []score.Score

	if e := s.database.Find(&result).Error; e != nil {
		return nil, fmt.Errorf("list scores: %w", e)
	}

	return result, nil
}
