package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"

func (s *Store) PushQueueBroadcast(
	callsigns []string,
	kind string,
	body string,
) error {
	var entries []queue.Entry

	for _, c := range callsigns {
		entries = append(entries, *queue.New(c, kind, body))
	}

	if len(entries) == 0 {
		return nil
	}

	return s.database.Create(&entries).Error
}
