package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/convert"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func (s *Server) SetHost(
	_ context.Context,
	r server.SetHostRequestObject,
) (server.SetHostResponseObject, error) {
	e := s.opnsense.SetHost(r.Identifier, convert.HostRequest(r.Body))

	if validation.Is(e) {
		return server.SetHost400JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.SetHost404JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.SetHost500JSONResponse(*s.answer(e)), nil
	}

	if apply(r.Params.Apply) {
		if f := s.opnsense.ReconfigureDnsmasq(); f != nil {
			return server.SetHost500JSONResponse(*s.answer(f)), nil
		}
	}

	return server.SetHost204Response{}, nil
}
