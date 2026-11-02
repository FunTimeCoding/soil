package server

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/server"
)

func (s *Server) PostRechunk(
	_ context.Context,
	_ server.PostRechunkRequestObject,
) (server.PostRechunkResponseObject, error) {
	documents, e := s.service.Rechunk()

	if e != nil {
		return server.PostRechunk500JSONResponse(
			*s.captureFail(e, soil.UnexpectedError),
		), nil
	}

	if documents == nil {
		documents = []string{}
	}

	return server.PostRechunk200JSONResponse{Documents: documents}, nil
}
