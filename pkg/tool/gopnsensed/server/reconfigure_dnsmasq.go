package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func (s *Server) ReconfigureDnsmasq(
	_ context.Context,
	_ server.ReconfigureDnsmasqRequestObject,
) (server.ReconfigureDnsmasqResponseObject, error) {
	if e := s.opnsense.ReconfigureDnsmasq(); e != nil {
		return server.ReconfigureDnsmasq500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.ReconfigureDnsmasq204Response{}, nil
}
