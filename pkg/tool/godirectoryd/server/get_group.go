package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/convert"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/server"
)

func (s *Server) GetGroup(
	_ context.Context,
	_ server.GetGroupRequestObject,
) (server.GetGroupResponseObject, error) {
	result, e := s.service.Groups()

	if e != nil {
		return server.GetGroup500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.GetGroup200JSONResponse(convert.Groups(result)), nil
}
