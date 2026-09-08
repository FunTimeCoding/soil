package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store/subscription"
)

func (s *Store) MustAll() []subscription.Subscription {
	result, e := s.All()
	errors.PanicOnError(e)

	return result
}
