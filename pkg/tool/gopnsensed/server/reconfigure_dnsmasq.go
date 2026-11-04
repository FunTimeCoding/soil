package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func (s *Server) ReconfigureDnsmasq(
	x context.Context,
	_ server.ReconfigureDnsmasqRequestObject,
) (server.ReconfigureDnsmasqResponseObject, error) {
	if e := s.service.ReconfigureDnsmasq(); e != nil {
		return server.ReconfigureDnsmasq500JSONResponse(*s.answer(x, e)), nil
	}

	return server.ReconfigureDnsmasq204Response{}, nil
}
