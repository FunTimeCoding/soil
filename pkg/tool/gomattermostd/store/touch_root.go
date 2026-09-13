package store

import (
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store/subscription"
	"time"
)

func (s *Store) TouchRoot(
	root string,
	at time.Time,
) error {
	return s.database.
		Model(subscription.Stub()).
		Where("root_identifier = ?", root).
		Update("last_event_at", at).Error
}
