package store

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/user"
	"gorm.io/gorm"
)

func (s *Store) UserMail(identifier string) (string, error) {
	var row user.User
	e := s.mapper.Where("identifier = ?", identifier).First(&row).Error

	if errors.Is(e, gorm.ErrRecordNotFound) {
		return "", nil
	}

	if e != nil {
		return "", e
	}

	return row.Mail, nil
}
