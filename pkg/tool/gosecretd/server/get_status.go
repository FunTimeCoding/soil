package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/server"
)

func (s *Server) GetStatus(
	x context.Context,
	_ server.GetStatusRequestObject,
) (server.GetStatusResponseObject, error) {
	result, e := s.service.Status()

	if unavailable.Is(e) {
		return server.GetStatus503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetStatus500JSONResponse(*s.answer(x, e)), nil
	}

	return server.GetStatus200JSONResponse{
		Initialized: result.Initialized,
		Sealed:      result.Sealed,
		Version:     result.Version,
	}, nil
}
