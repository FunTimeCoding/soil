package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/server"
)

func (s *Server) DeleteContext(
	_ context.Context,
	r server.DeleteContextRequestObject,
) (server.DeleteContextResponseObject, error) {
	deleted := s.service.RemoveContext(
		r.Params.Collection,
		r.Params.PathPrefix,
	)

	return server.DeleteContext200JSONResponse{
		Deleted: new(bool(deleted)),
	}, nil
}
