package store

import "github.com/funtimecoding/soil/pkg/chat/telegram/database/channel"

func (s *Store) ChannelByName(name string) (*channel.Channel, error) {
	var result []*channel.Channel
	e := s.database.
		Where("name = ?", name).
		Limit(1).
		Find(&result).Error

	if e != nil {
		return nil, e
	}

	if len(result) == 0 {
		return nil, nil
	}

	return result[0], nil
}
