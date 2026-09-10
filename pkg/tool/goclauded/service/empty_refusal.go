package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/refusal"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
)

func (s *Service) EmptyRefusal(r *session.Session) (*refusal.Refusal, error) {
	if _, found := s.cache.Get(r.Identifier); found {
		return refusal.New("session has a transcript"), nil
	}

	if r.TurnCount > 0 {
		return refusal.New("session has turns"), nil
	}

	if r.Lines > 0 {
		return refusal.New("session has transcript lines"), nil
	}

	if r.Alias != nil && *r.Alias != "" {
		return refusal.New("session has an alias"), nil
	}

	if r.Description != "" {
		return refusal.New("session has a description"), nil
	}

	for _, c := range s.emptyChecks() {
		count, e := c.count(r.Identifier)

		if e != nil {
			return nil, e
		}

		if count > 0 {
			return refusal.New(c.message), nil
		}
	}

	return nil, nil
}
