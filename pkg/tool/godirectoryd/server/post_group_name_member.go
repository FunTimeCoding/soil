package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/convert"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/server"
)

func (s *Server) PostGroupNameMember(
	_ context.Context,
	r server.PostGroupNameMemberRequestObject,
) (server.PostGroupNameMemberResponseObject, error) {
	result, e := s.service.AddMember(r.Name, r.Body.Account)

	if not_found.Is(e) {
		return server.PostGroupNameMember404JSONResponse{Error: e.Error()}, nil
	}

	if conflict.Is(e) {
		return server.PostGroupNameMember409JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.PostGroupNameMember500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.PostGroupNameMember200JSONResponse(*convert.Group(result)), nil
}
