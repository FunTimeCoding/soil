package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/mirror"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/score"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (s *Store) PutScores(
	scores []*score.Score,
	m *mirror.Mirror,
) error {
	return s.database.Transaction(
		func(t *gorm.DB) error {
			if len(scores) > 0 {
				if e := t.Clauses(clause.OnConflict{UpdateAll: true}).
					CreateInBatches(scores, 500).Error; e != nil {
					return fmt.Errorf("save scores: %w", e)
				}
			}

			if e := t.Save(m).Error; e != nil {
				return fmt.Errorf("save score mirror: %w", e)
			}

			return nil
		},
	)
}
