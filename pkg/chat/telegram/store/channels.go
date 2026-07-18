package store

import "github.com/funtimecoding/soil/pkg/chat/telegram/database/channel"

func (s *Store) Channels() ([]*channel.Channel, error) {
	var result []*channel.Channel

	return result, s.database.
		Order("name").
		Find(&result).Error
}
