package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/generated/server"
)

func (s *Server) GetStatus(
	_ context.Context,
	_ server.GetStatusRequestObject,
) (server.GetStatusResponseObject, error) {
	return server.GetStatus200JSONResponse{
		TotalEntries: s.store.Count(),
	}, nil
}
