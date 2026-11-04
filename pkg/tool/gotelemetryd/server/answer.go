package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/generated/server"
)

func (s *Server) answer(x context.Context, e error) *server.ErrorResponse {
	return s.answerWith(x, e, "")
}
