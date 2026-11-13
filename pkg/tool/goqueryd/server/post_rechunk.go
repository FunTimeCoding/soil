package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/server"
)

func (s *Server) PostRechunk(
	x context.Context,
	_ server.PostRechunkRequestObject,
) (server.PostRechunkResponseObject, error) {
	documents, e := s.service.Rechunk()

	if e != nil {
		return server.PostRechunk500JSONResponse(*s.answer(x, e)), nil
	}

	if documents == nil {
		documents = []string{}
	}

	return server.PostRechunk200JSONResponse{Documents: documents}, nil
}
