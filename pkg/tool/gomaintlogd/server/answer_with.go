package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/generated/server"
)

func (s *Server) answerWith(
	x context.Context,
	e error,
	message string,
) *server.ErrorResponse {
	text, event := classify.Capture(x, s.reporter, e, message)

	return &server.ErrorResponse{Error: text, EventIdentifier: event}
}
