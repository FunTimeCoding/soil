package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/server"
)

func (s *Server) ListSchemas(
	c context.Context,
	r server.ListSchemasRequestObject,
) (server.ListSchemasResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if not_found.Is(e) {
		return server.ListSchemas404JSONResponse(*clientError(e)), nil
	}

	if not_selected.Is(e) {
		return server.ListSchemas400JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListSchemas500JSONResponse(*s.answerWith(e, "")), nil
	}

	rows, e := s.service.ListSchemas(c, instance)

	if validation.Is(e) {
		return server.ListSchemas400JSONResponse(*clientError(e)), nil
	}

	if forbidden.Is(e) {
		return server.ListSchemas403JSONResponse(*clientError(e)), nil
	}

	if conflict.Is(e) {
		return server.ListSchemas409JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.ListSchemas503JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListSchemas500JSONResponse(*s.answerWith(e, "")), nil
	}

	return server.ListSchemas200JSONResponse{
		Columns: rows.Columns,
		Rows:    rows.Rows,
	}, nil
}
