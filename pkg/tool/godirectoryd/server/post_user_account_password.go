package server

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/server"
)

func (s *Server) PostUserAccountPassword(
	_ context.Context,
	r server.PostUserAccountPasswordRequestObject,
) (server.PostUserAccountPasswordResponseObject, error) {
	e := s.service.SetPassword(r.Account, r.Body.Password)

	if not_found.Is(e) {
		return server.PostUserAccountPassword404JSONResponse{
			Error: e.Error(),
		}, nil
	}

	if e != nil {
		return server.PostUserAccountPassword500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.PostUserAccountPassword200JSONResponse{
		Message: fmt.Sprintf("password set: %s", r.Account),
	}, nil
}
