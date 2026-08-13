package server

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/constant"
	gomemoryd "github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/server"
)

func (s *Server) PostMemory(
	_ context.Context,
	r server.PostMemoryRequestObject,
) (server.PostMemoryResponseObject, error) {
	m, e := s.service.CreateMemory(saveOption(*r.Body))

	if e != nil {
		if errors.Is(e, gomemoryd.ErrorReservedScope) {
			return server.PostMemory400JSONResponse(*clientError(e)), nil
		}

		return server.PostMemory500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.PostMemory200JSONResponse{Identifier: m.Identifier}, nil
}
