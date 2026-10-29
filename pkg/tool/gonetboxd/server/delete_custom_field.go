package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) DeleteCustomField(
	_ context.Context,
	r server.DeleteCustomFieldRequestObject,
) (server.DeleteCustomFieldResponseObject, error) {
	e := s.service.DeleteCustomField(r.Name)

	if not_found.Is(e) {
		return server.DeleteCustomField404JSONResponse(*clientError(e)), nil
	}

	if validation.Is(e) {
		return server.DeleteCustomField400JSONResponse(*clientError(e)), nil
	}

	if forbidden.Is(e) {
		return server.DeleteCustomField403JSONResponse(*clientError(e)), nil
	}

	if conflict.Is(e) {
		return server.DeleteCustomField409JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.DeleteCustomField500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.DeleteCustomField204Response{}, nil
}
