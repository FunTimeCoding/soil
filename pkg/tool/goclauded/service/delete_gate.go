package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/refusal"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
)

func (s *Service) deleteGate(
	identifier string,
	confirm string,
) (*session.Session, error) {
	r, found, e := s.store.FindSession(identifier)

	if e != nil {
		return nil, e
	}

	if !found {
		return nil, refusal.New(
			fmt.Sprintf("session not found: %s", identifier),
		)
	}

	empty, f := s.EmptyRefusal(r)

	if f != nil {
		return nil, f
	}

	if empty == nil {
		return r, nil
	}

	if confirm == "" {
		return nil, refusal.New(
			fmt.Sprintf("%s, confirmation required", empty.Error()),
		)
	}

	if confirm != s.DeleteHash(identifier) {
		return nil, refusal.New("confirmation does not match this session")
	}

	return r, nil
}
