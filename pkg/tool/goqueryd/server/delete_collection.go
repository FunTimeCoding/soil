package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/server"
)

func (s *Server) DeleteCollection(
	x context.Context,
	r server.DeleteCollectionRequestObject,
) (server.DeleteCollectionResponseObject, error) {
	deleted := s.service.DeleteCollection(r.Params.Name)

	return server.DeleteCollection200JSONResponse{Deleted: new(bool(deleted))}, nil
}
