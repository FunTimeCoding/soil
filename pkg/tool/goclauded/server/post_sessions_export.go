package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) PostSessionsExport(
	_ context.Context,
	_ server.PostSessionsExportRequestObject,
) (server.PostSessionsExportResponseObject, error) {
	var paths []string

	for _, e := range s.transcript.Sessions() {
		paths = append(paths, s.transcript.ExportSession(e))
	}

	return server.PostSessionsExport200JSONResponse{Paths: paths}, nil
}
