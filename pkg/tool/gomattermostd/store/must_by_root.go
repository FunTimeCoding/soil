package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store/subscription"
)

func (s *Store) MustByRoot(root string) []subscription.Subscription {
	result, e := s.ByRoot(root)
	errors.PanicOnError(e)

	return result
}
