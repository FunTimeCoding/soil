package store

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/client"
	"gorm.io/gorm"
)

func (s *Store) ClientRow(identifier string) (*client.Client, error) {
	var row client.Client
	e := s.mapper.Where("identifier = ?", identifier).First(&row).Error

	if e != nil {
		if errors.Is(e, gorm.ErrRecordNotFound) {
			return nil, not_found.New(constant.ClientKind, identifier)
		}

		return nil, e
	}

	return &row, nil
}
