package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
)

func (s *Store) PushQueueBroadcast(
	sessions []session.Session,
	kind string,
	body string,
) error {
	var entries []queue.Entry

	for _, e := range sessions {
		entries = append(
			entries,
			*queue.New(e.Identifier, e.CallsignValue(), kind, body),
		)
	}

	if len(entries) == 0 {
		return nil
	}

	return s.database.Create(&entries).Error
}
