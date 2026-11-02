package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
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

	if validation.Is(f) {
		return server.EnableRedirection400JSONResponse(*clientError(f)), nil
	}

	if forbidden.Is(f) {
		return server.EnableRedirection403JSONResponse(*clientError(f)), nil
	}

	if unavailable.Is(f) {
		return server.EnableRedirection503JSONResponse(*clientError(f)), nil
	}

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
