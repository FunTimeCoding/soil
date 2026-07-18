package store

import "github.com/funtimecoding/soil/pkg/chat/telegram/database/user"

func (s *Store) Users() ([]*user.User, error) {
	var result []*user.User

	return result, s.database.
		Order("name").
		Find(&result).Error
}
