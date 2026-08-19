package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
)

func (s *Server) GetRoot(
	_ context.Context,
	_ server.GetRootRequestObject,
) (server.GetRootResponseObject, error) {
	result, e := s.store.Authority(constant.RootAuthority)

	if e != nil || result == nil {
		return server.GetRoot404JSONResponse(
			*clientError(constant.RootMissing),
		), nil
	}

	return server.GetRoot200TextResponse(result.Certificate), nil
}
