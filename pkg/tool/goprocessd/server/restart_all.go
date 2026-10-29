package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/generated/server"
)

func (s *Server) RestartAll(
	_ context.Context,
	_ server.RestartAllRequestObject,
) (server.RestartAllResponseObject, error) {
	count, e := s.service.RestartAll()

	if e != nil {
		return server.RestartAll409JSONResponse{Error: e.Error()}, nil
	}

	return server.RestartAll200JSONResponse{Count: count}, nil
}
