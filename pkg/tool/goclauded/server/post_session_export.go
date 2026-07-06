package server

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) PostSessionExport(
	_ context.Context,
	r server.PostSessionExportRequestObject,
) (server.PostSessionExportResponseObject, error) {
	e := s.service.Resolve(r.Identifier)

	if e.Identifier == "" {
		return server.PostSessionExport404JSONResponse{
			Error: fmt.Sprintf("session not found: %s", r.Identifier),
		}, nil
	}

	return server.PostSessionExport200JSONResponse{
		Paths: []string{s.exportSession(e, s.sessionExportPath)},
	}, nil
}
