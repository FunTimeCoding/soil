package store

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/user"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *Store) AuthenticateUser(
	mail string,
	password string,
) (*user.User, error) {
	var row user.User
	e := s.mapper.Where("mail = ?", mail).First(&row).Error

	if e != nil {
		if errors.Is(e, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, e
	}

	e = bcrypt.CompareHashAndPassword(
		[]byte(row.PasswordHash),
		[]byte(password),
	)

	if e != nil {
		return nil, nil
	}

	return &row, nil
}
