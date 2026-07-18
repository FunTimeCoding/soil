package store

import "github.com/funtimecoding/soil/pkg/chat/telegram/database/user"

func (s *Store) SaveUser(
	identifier int64,
	name string,
) error {
	u := user.New()
	u.Name = name
	u.Identifier = identifier

	return s.database.
		Where("identifier = ?", identifier).
		FirstOrCreate(u).Error
}
