package server

import "github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/server"

func (s *Server) captureFail(
	e error,
	message string,
) *server.ErrorResponse {
	return &server.ErrorResponse{
		Error:           message,
		EventIdentifier: s.reporter.CaptureException(e),
	}
}
