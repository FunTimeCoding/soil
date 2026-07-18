package store

import (
	"github.com/funtimecoding/soil/pkg/chat/telegram/database/channel"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (s *Store) MustChannels() []*channel.Channel {
	result, e := s.Channels()
	errors.PanicOnError(e)

	return result
}
