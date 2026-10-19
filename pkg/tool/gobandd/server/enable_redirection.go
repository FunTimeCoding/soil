package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/server"
)

func (s *Server) EnableRedirection(
	_ context.Context,
	r server.EnableRedirectionRequestObject,
) (server.EnableRedirectionResponseObject, error) {
	instance, e := s.resolveInstance(r.Body.Instance)

	if not_found.Is(e) {
		return server.EnableRedirection404JSONResponse(*clientError(e)), nil
	}

	if not_selected.Is(e) {
		return server.EnableRedirection400JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.EnableRedirection500JSONResponse(
			*s.captureFail(e, e.Error()),
		), nil
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
