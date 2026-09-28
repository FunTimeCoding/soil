package store

import "github.com/funtimecoding/soil/pkg/tool/gogated/model/login_session"

func (s *Store) DeleteLoginSession(identifier string) error {
	return s.mapper.Where("identifier = ?", identifier).
		Delete(login_session.Stub()).Error
}
