package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/convert"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/server"
)

func (s *Server) GetUser(
	_ context.Context,
	_ server.GetUserRequestObject,
) (server.GetUserResponseObject, error) {
	result, e := s.service.Users()

	if e != nil {
		return server.GetUser500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.GetUser200JSONResponse(convert.Users(result)), nil
}
