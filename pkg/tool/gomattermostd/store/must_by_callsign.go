package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store/subscription"
)

func (s *Store) MustByCallsign(callsign string) []subscription.Subscription {
	result, e := s.ByCallsign(callsign)
	errors.PanicOnError(e)

	return result
}
