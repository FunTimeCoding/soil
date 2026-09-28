package store

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/client"
	"github.com/funtimecoding/soil/pkg/tool/gogated/store/fosite_client"
	"github.com/ory/fosite"
	"gorm.io/gorm"
)

func (s *Store) getClientInternal(identifier string) (fosite.Client, error) {
	var row client.Client
	e := s.mapper.Where("identifier = ?", identifier).First(&row).Error

	if e != nil {
		if errors.Is(e, gorm.ErrRecordNotFound) {
			return nil, fosite.ErrNotFound
		}

		return nil, e
	}

	return fosite_client.New(&row), nil
}
