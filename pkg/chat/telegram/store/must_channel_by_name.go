package store

import (
	"github.com/funtimecoding/soil/pkg/chat/telegram/database/channel"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (s *Store) MustChannelByName(name string) *channel.Channel {
	result, e := s.ChannelByName(name)
	errors.PanicOnError(e)

	return result
}
