package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/server"
)

func (s *Server) DeleteDocument(
	x context.Context,
	r server.DeleteDocumentRequestObject,
) (server.DeleteDocumentResponseObject, error) {
	deleted, e := s.service.DeleteDocument(r.Params.Collection, r.Params.Path)

	if e != nil {
		return server.DeleteDocument500JSONResponse(*s.answer(x, e)), nil
	}

	return server.DeleteDocument200JSONResponse{Deleted: new(bool(deleted))}, nil
}
