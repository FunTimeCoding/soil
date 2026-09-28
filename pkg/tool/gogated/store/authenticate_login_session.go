package store

import "github.com/funtimecoding/soil/pkg/tool/gogated/model/login_session"

func (s *Store) AuthenticateLoginSession(
	identifier string,
	userIdentifier string,
) error {
	return s.mapper.Model(login_session.Stub()).
		Where("identifier = ?", identifier).
		Updates(
			map[string]any{
				"user_identifier": userIdentifier,
				"authenticated":   true,
			},
		).Error
}
