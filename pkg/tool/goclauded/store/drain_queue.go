package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"

func (s *Store) DrainQueue(callsign string) ([]queue.Entry, error) {
	var result []queue.Entry

	if e := s.database.Where(
		"callsign = ? AND consumed = ?",
		callsign,
		false,
	).Order("created_at").Find(&result).Error; e != nil {
		return nil, e
	}

	if len(result) > 0 {
		if e := s.database.Model(queue.Stub()).Where(
			"callsign = ? AND consumed = ?",
			callsign,
			false,
		).Update("consumed", true).Error; e != nil {
			return nil, e
		}
	}

	return result, nil
}
