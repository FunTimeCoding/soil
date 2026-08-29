package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func (s *Server) DeleteHost(
	_ context.Context,
	r server.DeleteHostRequestObject,
) (server.DeleteHostResponseObject, error) {
	e := s.opnsense.DeleteHost(r.Identifier)

	if not_found.Is(e) {
		return server.DeleteHost404JSONResponse(*clientError(e.Error())), nil
	}

	if e != nil {
		return server.DeleteHost500JSONResponse(*s.captureDetail(e)), nil
	}

	if apply(r.Params.Apply) {
		if f := s.opnsense.ReconfigureDnsmasq(); f != nil {
			return server.DeleteHost500JSONResponse(*s.captureDetail(f)), nil
		}
	}

	return server.DeleteHost204Response{}, nil
}
