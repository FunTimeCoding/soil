package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) RemoveVirtualLabel(
	_ context.Context,
	r server.RemoveVirtualLabelRequestObject,
) (server.RemoveVirtualLabelResponseObject, error) {
	e := s.service.RemoveVirtualLabel(r.Name, r.Key)

	if not_found.Is(e) {
		return server.RemoveVirtualLabel404JSONResponse(*clientError(e)), nil
	}

	if validation.Is(e) {
		return server.RemoveVirtualLabel400JSONResponse(*clientError(e)), nil
	}

	if forbidden.Is(e) {
		return server.RemoveVirtualLabel403JSONResponse(*clientError(e)), nil
	}

	if conflict.Is(e) {
		return server.RemoveVirtualLabel409JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.RemoveVirtualLabel500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.RemoveVirtualLabel204Response{}, nil
}
