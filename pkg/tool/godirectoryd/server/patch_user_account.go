package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/convert"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/server"
)

func (s *Server) PatchUserAccount(
	_ context.Context,
	r server.PatchUserAccountRequestObject,
) (server.PatchUserAccountResponseObject, error) {
	result, e := s.service.ModifyUser(
		r.Account,
		text(r.Body.Name),
		text(r.Body.Surname),
		text(r.Body.Mail),
	)

	if not_found.Is(e) {
		return server.PatchUserAccount404JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.PatchUserAccount500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.PatchUserAccount200JSONResponse(*convert.User(result)), nil
}
