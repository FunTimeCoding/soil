package server

import "github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"

func (s *Server) captureDetail(e error) *server.ErrorResponse {
	return s.answerWith(e, "")
}
