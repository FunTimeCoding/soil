package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store/subscription"
	"time"
)

func (s *Store) MustPrune(cutoff time.Time) []subscription.Subscription {
	result, e := s.Prune(cutoff)
	errors.PanicOnError(e)

	return result
}
