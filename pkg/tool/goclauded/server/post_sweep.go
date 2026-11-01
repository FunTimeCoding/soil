package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) PostSweep(
	_ context.Context,
	_ server.PostSweepRequestObject,
) (server.PostSweepResponseObject, error) {
	result := s.service.Sweep()

	return server.PostSweep200JSONResponse{
		Copied:  result.Copied,
		Updated: result.Updated,
		Skipped: result.Skipped,
	}, nil
}
