package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/seed"
)

func (s *Store) UpsertSeed(
	name string,
	path string,
	contentHash string,
	content string,
) {
	var existing seed.Seed
	result := s.mapper.Where("path = ?", path).Limit(1).Find(&existing)
	errors.PanicOnError(result.Error)

	if result.RowsAffected == 0 {
		var maxPosition int
		s.mapper.Model(seed.Stub()).Select("COALESCE(MAX(position), 0)").Scan(&maxPosition)
		errors.PanicOnError(
			s.mapper.Create(
				seed.New(name, path, contentHash, content, maxPosition+1),
			).Error,
		)

		return
	}

	if existing.ContentHash != contentHash {
		errors.PanicOnError(
			s.mapper.Model(&existing).Updates(
				map[string]any{
					"name":         name,
					"content_hash": contentHash,
					"content":      content,
				},
			).Error,
		)
	}
}
