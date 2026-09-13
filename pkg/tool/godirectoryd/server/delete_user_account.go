package server

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/server"
)

func (s *Server) DeleteUserAccount(
	_ context.Context,
	r server.DeleteUserAccountRequestObject,
) (server.DeleteUserAccountResponseObject, error) {
	e := s.service.DeleteUser(r.Account)

	if not_found.Is(e) {
		return server.DeleteUserAccount404JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.DeleteUserAccount500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.DeleteUserAccount200JSONResponse{
		Message: fmt.Sprintf("deleted user: %s", r.Account),
	}, nil
}
