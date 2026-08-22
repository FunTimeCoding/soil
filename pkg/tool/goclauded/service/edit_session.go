package service

import (
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/argument/edit_session"
)

func (s *Service) EditSession(
	identifier string,
	a *edit_session.Session,
) error {
	if a.Alias != nil && *a.Alias != "" {
		owner, e := s.store.AliasOwner(*a.Alias)

		if e != nil {
			return e
		}

		if owner != "" && owner != identifier {
			return conflict.Format(
				"alias already in use: %q is used by session %s",
				*a.Alias,
				owner,
			)
		}
	}

	if e := s.store.EditSession(identifier, a); e != nil {
		return e
	}

	s.notify()

	return nil
}
