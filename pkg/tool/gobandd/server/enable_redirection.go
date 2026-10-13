package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/server"
)

func (s *Server) EnableRedirection(
	_ context.Context,
	r server.EnableRedirectionRequestObject,
) (server.EnableRedirectionResponseObject, error) {
	instance, e := s.resolveInstance(r.Body.Instance)

	if e != nil {
		return server.EnableRedirection400JSONResponse(*clientError(e)), nil
	}

	enabled, f := s.service.EnableRedirection(instance)

	if f != nil {
		return server.EnableRedirection500JSONResponse(*s.captureDetail(f)), nil
	}

	if enabled == nil {
		enabled = []string{}
	}

	return server.EnableRedirection200JSONResponse{
		Instance: instance,
		Enabled:  enabled,
	}, nil
}
