package store

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/user"
	"gorm.io/gorm"
)

func (s *Store) UserByMail(mail string) (*user.User, error) {
	var row user.User
	e := s.mapper.Where("mail = ?", mail).First(&row).Error

	if e != nil {
		if errors.Is(e, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, e
	}

	return &row, nil
}
