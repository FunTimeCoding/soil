package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/message"

func (s *Store) SendMessage(
	fromName string,
	toName string,
	body string,
) (*message.Message, error) {
	result := message.New(fromName, toName, body)
	result.CreatedAt = s.clock().UTC()

	return result, s.database.Create(result).Error
}
