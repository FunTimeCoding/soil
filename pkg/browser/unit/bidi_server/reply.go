package bidi_server

import (
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/browser/ladybird/types/message"
)

func (s *Server) reply(command *message.Request) *message.Reply {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.received = append(s.received, command)
	result := message.NewReply()

	if answer, okay := s.answers[command.Method]; okay {
		*result = *answer
	} else {
		result.Type = constant.LadybirdErrorType
		result.Error = constant.LadybirdUnknownCommand
		result.Message = "Command is not supported"
	}

	result.Identifier = command.Identifier

	return result
}
