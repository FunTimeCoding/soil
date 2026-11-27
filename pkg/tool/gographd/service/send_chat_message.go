package service

import (
	"github.com/funtimecoding/soil/pkg/chat/teams/destination"
	"github.com/funtimecoding/soil/pkg/chat/teams/message"
)

func (s *Service) SendChatMessage(
	chat string,
	text string,
	mention []string,
) (*message.Message, error) {
	if e := checkChat(chat); e != nil {
		return nil, e
	}

	return s.post(destination.Chat(chat), text, mention)
}
