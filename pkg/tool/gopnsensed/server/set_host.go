package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/convert"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func (s *Server) SetHost(
	_ context.Context,
	r server.SetHostRequestObject,
) (server.SetHostResponseObject, error) {
	e := s.opnsense.SetHost(r.Identifier, convert.HostRequest(r.Body))

	if e != nil {
		return server.SetHost500JSONResponse(*s.captureDetail(e)), nil
	}

	if apply(r.Params.Apply) {
		if f := s.opnsense.ReconfigureDnsmasq(); f != nil {
			return server.SetHost500JSONResponse(*s.captureDetail(f)), nil
		}
	}

	return server.SetHost204Response{}, nil
}
