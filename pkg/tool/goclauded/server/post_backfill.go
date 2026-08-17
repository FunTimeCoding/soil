package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) PostBackfill(
	_ context.Context,
	r server.PostBackfillRequestObject,
) (server.PostBackfillResponseObject, error) {
	result := s.service.BackfillAllSessions()

	if r.Params.Cold != nil && *r.Params.Cold {
		result = s.service.ColdBackfillAllSessions()
	}

	return server.PostBackfill200JSONResponse{
		Enriched: result.Enriched,
		Skipped:  result.Skipped,
	}, nil
}
