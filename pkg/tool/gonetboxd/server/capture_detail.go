package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) captureDetail(x context.Context, e error) *server.ErrorResponse {
	return s.answerMessage(x, e, "")
}
