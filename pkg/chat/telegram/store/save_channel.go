package store

import "github.com/funtimecoding/soil/pkg/chat/telegram/database/channel"

func (s *Store) SaveChannel(
	identifier int64,
	name string,
) error {
	h := channel.New()
	h.Name = name
	h.Identifier = identifier

	return s.database.
		Where("identifier = ?", identifier).
		FirstOrCreate(h).Error
}
