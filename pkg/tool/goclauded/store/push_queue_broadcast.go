package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
)

func (s *Store) PushQueueBroadcast(
	sessions []session.Session,
	template *queue.Entry,
) error {
	var entries []queue.Entry

	for _, v := range sessions {
		e := *template
		e.SessionIdentifier = v.Identifier
		e.Callsign = v.CallsignValue()
		entries = append(entries, e)
	}

	if len(entries) == 0 {
		return nil
	}

	return s.database.Create(&entries).Error
}
