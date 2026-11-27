package service

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/browser/ladybird/types/request"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
)

func (s *Service) act(
	tab string,
	uid string,
	declaration string,
	arguments ...any,
) error {
	shared, e := s.resolveUID(tab, uid)

	if e != nil {
		return e
	}

	_, f := guarded(
		s,
		func() (json.RawMessage, error) {
			return s.client.CallFunction(
				tab,
				declaration,
				append([]any{request.NewReference(shared)}, arguments...)...,
			)
		},
	)

	if not_found.Is(f) {
		return not_found.Format(
			"uid %s is gone from the page - take a new snapshot",
			uid,
		)
	}

	return f
}
