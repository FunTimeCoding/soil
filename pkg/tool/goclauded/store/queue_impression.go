package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/impression_outbox"
	"time"
)

func (s *Store) QueueImpression(
	source string,
	content string,
	capturedAt time.Time,
) error {
	return s.database.Create(
		impression_outbox.New(source, content, capturedAt),
	).Error
}
