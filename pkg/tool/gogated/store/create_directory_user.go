package store

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/user"
	"github.com/google/uuid"
)

func (s *Store) CreateDirectoryUser(
	unique string,
	account string,
	mail string,
	name string,
) (*user.User, error) {
	row := user.NewDirectory(uuid.New().String(), unique, account, mail, name)

	if e := s.mapper.Create(row).Error; e != nil {
		return nil, e
	}

	return row, nil
}
