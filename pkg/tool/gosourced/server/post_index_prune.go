package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/server"
)

func (s *Server) PostIndexPrune(
	_ context.Context,
	_ server.PostIndexPruneRequestObject,
) (server.PostIndexPruneResponseObject, error) {
	removed := s.workspace.Prune()

	if removed == nil {
		removed = []string{}
	}

	return server.PostIndexPrune200JSONResponse{Removed: removed}, nil
}
