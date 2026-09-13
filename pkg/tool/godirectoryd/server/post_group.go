package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/convert"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/server"
)

func (s *Server) PostGroup(
	_ context.Context,
	r server.PostGroupRequestObject,
) (server.PostGroupResponseObject, error) {
	result, e := s.service.CreateGroup(r.Body.Name)

	if conflict.Is(e) {
		return server.PostGroup409JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.PostGroup500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.PostGroup200JSONResponse(*convert.Group(result)), nil
}
