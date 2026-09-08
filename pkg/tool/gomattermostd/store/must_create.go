package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store/subscription"
)

func (s *Store) MustCreate(v *subscription.Subscription) {
	errors.PanicOnError(s.Create(v))
}
