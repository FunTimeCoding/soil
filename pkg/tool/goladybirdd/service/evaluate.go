package service

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors/validation"
)

func (s *Service) Evaluate(
	tab string,
	expression string,
) (json.RawMessage, error) {
	if expression == "" {
		return nil, validation.New("expression is required")
	}

	t, e := s.resolveTab(tab)

	if e != nil {
		return nil, e
	}

	s.navigation.Settle(t)

	return guarded(
		s,
		func() (json.RawMessage, error) {
			return s.client.Evaluate(t, expression)
		},
	)
}
