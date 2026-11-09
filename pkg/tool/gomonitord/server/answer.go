package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/server"
)

func (s *Server) answer(x context.Context, e error) *server.ErrorResponse {
	return s.answerMessage(x, e, "")
}
