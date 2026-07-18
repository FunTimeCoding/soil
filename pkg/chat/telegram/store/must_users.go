package store

import (
	"github.com/funtimecoding/soil/pkg/chat/telegram/database/user"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (s *Store) MustUsers() []*user.User {
	result, e := s.Users()
	errors.PanicOnError(e)

	return result
}
