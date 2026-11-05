package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func (s *Server) ReconfigureDnsmasq(
	x context.Context,
	_ server.ReconfigureDnsmasqRequestObject,
) (server.ReconfigureDnsmasqResponseObject, error) {
	e := s.service.ReconfigureDnsmasq()

	if unavailable.Is(e) {
		return server.ReconfigureDnsmasq503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ReconfigureDnsmasq500JSONResponse(*s.answer(x, e)), nil
	}

	return server.ReconfigureDnsmasq204Response{}, nil
}
