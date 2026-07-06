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

	for _, e := range s.service.Sessions() {
		paths = append(
			paths,
			s.exportSession(e, s.sessionExportPath),
		)
	}

	return server.PostSessionsExport200JSONResponse{Paths: paths}, nil
}
