package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
)

func (s *Server) GetRoot(
	_ context.Context,
	_ server.GetRootRequestObject,
) (server.GetRootResponseObject, error) {
	result, e := s.service.Authority(constant.RootAuthority)

	if not_found.Is(e) {
		return server.GetRoot404JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetRoot500JSONResponse(*s.answer(e)), nil
	}

	return server.GetRoot200TextResponse(result.Certificate), nil
}
