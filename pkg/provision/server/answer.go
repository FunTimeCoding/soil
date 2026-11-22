package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/provision/generated/server"
)

func (s *Server) answer(
	x context.Context,
	e error,
) *server.ErrorResponse {
	text, event := classify.Capture(x, s.reporter, e, "")

	return &server.ErrorResponse{Error: text, EventIdentifier: event}
}
