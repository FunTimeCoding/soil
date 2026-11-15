package message

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func (s *Service) requireSession(identifier string) error {
	_, found, e := s.store.FindSession(identifier)

	if e != nil {
		return e
	}

	if !found {
		return not_found.New(constant.SessionKind, identifier)
	}

	return nil
}
