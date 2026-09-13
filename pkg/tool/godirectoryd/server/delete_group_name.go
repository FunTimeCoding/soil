package server

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/server"
)

func (s *Server) DeleteGroupName(
	_ context.Context,
	r server.DeleteGroupNameRequestObject,
) (server.DeleteGroupNameResponseObject, error) {
	e := s.service.DeleteGroup(r.Name)

	if not_found.Is(e) {
		return server.DeleteGroupName404JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.DeleteGroupName500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.DeleteGroupName200JSONResponse{
		Message: fmt.Sprintf("deleted group: %s", r.Name),
	}, nil
}
