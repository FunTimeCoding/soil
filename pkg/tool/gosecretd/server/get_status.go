package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/server"
)

func (s *Server) GetStatus(
	_ context.Context,
	_ server.GetStatusRequestObject,
) (server.GetStatusResponseObject, error) {
	result, e := s.service.Status()

	if e != nil {
		return server.GetStatus500JSONResponse(*s.answer(e)), nil
	}

	return server.GetStatus200JSONResponse{
		Initialized: result.Initialized,
		Sealed:      result.Sealed,
		Version:     result.Version,
	}, nil
}
