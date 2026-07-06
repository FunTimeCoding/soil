package store

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
	"gorm.io/gorm"
)

func (s *Store) Get(identifier uint) (*entry.Entry, error) {
	var result entry.Entry

	if e := s.database.First(&result, identifier).Error; e != nil {
		if errors.Is(e, gorm.ErrRecordNotFound) {
			return nil, not_found.New("entry", identifier)
		}

		return nil, e
	}

	return &result, nil
}
