package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/message"

func (s *Store) PendingMessages(name string) ([]message.Message, error) {
	var result []message.Message

	if e := s.database.Where(
		"(to_name = ? OR to_name = '') AND read = ?",
		name,
		false,
	).Order("created_at").Find(&result).Error; e != nil {
		return nil, e
	}

	if len(result) > 0 {
		if e := s.database.Model(message.Stub()).Where(
			"(to_name = ? OR to_name = '') AND read = ?",
			name,
			false,
		).Update("read", true).Error; e != nil {
			return nil, e
		}
	}

	return result, nil
}
