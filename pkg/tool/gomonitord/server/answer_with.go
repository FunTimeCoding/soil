package server

import (
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/server"
)

func (s *Server) answerWith(
	e error,
	message string,
) *server.ErrorResponse {
	text, event := classify.Capture(s.reporter, e, message)

	return &server.ErrorResponse{Error: text, EventIdentifier: event}
}
