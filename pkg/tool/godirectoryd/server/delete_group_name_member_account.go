package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/convert"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/server"
)

func (s *Server) DeleteGroupNameMemberAccount(
	_ context.Context,
	r server.DeleteGroupNameMemberAccountRequestObject,
) (server.DeleteGroupNameMemberAccountResponseObject, error) {
	result, e := s.service.RemoveMember(r.Name, r.Account)

	if not_found.Is(e) {
		return server.DeleteGroupNameMemberAccount404JSONResponse{
			Error: e.Error(),
		}, nil
	}

	if e != nil {
		return server.DeleteGroupNameMemberAccount500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.DeleteGroupNameMemberAccount200JSONResponse(
		*convert.Group(result),
	), nil
}
