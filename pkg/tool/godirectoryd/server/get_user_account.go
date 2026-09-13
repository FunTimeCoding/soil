package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/convert"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/server"
)

func (s *Server) GetUserAccount(
	_ context.Context,
	r server.GetUserAccountRequestObject,
) (server.GetUserAccountResponseObject, error) {
	result, e := s.service.User(r.Account)

	if not_found.Is(e) {
		return server.GetUserAccount404JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetUserAccount500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.GetUserAccount200JSONResponse(*convert.User(result)), nil
}
