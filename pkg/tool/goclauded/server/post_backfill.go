package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) PostBackfill(
	_ context.Context,
	_ server.PostBackfillRequestObject,
) (server.PostBackfillResponseObject, error) {
	result := s.service.BackfillAllSessions()

	return server.PostBackfill200JSONResponse{
		Enriched: result.Enriched,
		Skipped:  result.Skipped,
	}, nil
}
