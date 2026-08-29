package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/convert"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func (s *Server) AddHost(
	_ context.Context,
	r server.AddHostRequestObject,
) (server.AddHostResponseObject, error) {
	identifier, e := s.opnsense.AddHost(convert.HostRequest(r.Body))

	if e != nil {
		return server.AddHost500JSONResponse(*s.captureDetail(e)), nil
	}

	if apply(r.Params.Apply) {
		if f := s.opnsense.ReconfigureDnsmasq(); f != nil {
			return server.AddHost500JSONResponse(*s.captureDetail(f)), nil
		}
	}

	return server.AddHost200JSONResponse(*convert.Identifier(identifier)), nil
}
