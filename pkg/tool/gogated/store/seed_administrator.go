package store

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/user"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s *Store) SeedAdministrator(
	mail string,
	password string,
) error {
	var count int64
	s.mapper.Model(user.Stub()).Count(&count)

	if count > 0 {
		return nil
	}

	hash, e := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if e != nil {
		return e
	}

	return s.mapper.Create(
		user.New(uuid.New().String(), mail, string(hash)),
	).Error
}
