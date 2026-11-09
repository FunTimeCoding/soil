package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
)

func (s *Server) answer(x context.Context, e error) *server.ErrorResponse {
	return s.answerMessage(x, e, "")
}
