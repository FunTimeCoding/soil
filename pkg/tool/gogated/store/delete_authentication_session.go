package store

import "github.com/funtimecoding/soil/pkg/tool/gogated/model/authentication_session"

func (s *Store) DeleteAuthenticationSession(identifier string) error {
	return s.mapper.Where("identifier = ?", identifier).
		Delete(authentication_session.Stub()).Error
}
