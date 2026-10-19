package server

import "github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"

func (s *Server) answer(e error) *server.ErrorResponse {
	return s.answerWith(e, "")
}
