package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/server"
)

func (s *Server) DeleteMemory(
	_ context.Context,
	r server.DeleteMemoryRequestObject,
) (server.DeleteMemoryResponseObject, error) {
	source := ""

	if r.Params.Source != nil {
		source = *r.Params.Source
	}

	if e := s.service.ForgetMemory(r.Identifier, source); e != nil {
		return server.DeleteMemory500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.DeleteMemory200JSONResponse{Identifier: r.Identifier}, nil
}
