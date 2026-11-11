package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/backfill_result"
)

func (s *Server) PostBackfill(
	x context.Context,
	r server.PostBackfillRequestObject,
) (server.PostBackfillResponseObject, error) {
	var result *backfill_result.Result

	if r.Params.Cold != nil && *r.Params.Cold {
		result = s.transcript.ColdBackfillAllSessions()
	} else {
		result = s.transcript.BackfillAllSessions()
	}

	return server.PostBackfill200JSONResponse{
		Enriched: result.Enriched,
		Skipped:  result.Skipped,
	}, nil
}
