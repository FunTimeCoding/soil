package server

import "github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"

func (s *Server) captureFail(
	e error,
	message string,
) *server.ErrorResponse {
	return &server.ErrorResponse{
		Error:           message,
		EventIdentifier: s.reporter.CaptureException(e),
	}
}
