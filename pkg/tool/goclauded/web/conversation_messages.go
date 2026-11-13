package web

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/message"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func (s *Server) conversationMessages(
	identifier string,
) ([]message.Message, error) {
	session := s.transcript.Resolve(identifier)

	if session.Identifier == "" {
		return nil, not_found.New(constant.SessionKind, identifier)
	}

	return s.transcript.Messages(session.Identifier), nil
}
