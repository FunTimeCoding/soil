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

func (s *Server) Query(
	c context.Context,
	r server.QueryRequestObject,
) (server.QueryResponseObject, error) {
	instance, e := s.resolveInstance(r.Body.Instance)

	if not_found.Is(e) {
		return server.Query404JSONResponse(*clientError(e)), nil
	}

	if not_selected.Is(e) {
		return server.Query400JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.Query500JSONResponse(*s.answerWith(e, "")), nil
	}

	rows, e := s.service.Query(c, instance, r.Body.Sql)

	if validation.Is(e) {
		return server.Query400JSONResponse(*clientError(e)), nil
	}

	if forbidden.Is(e) {
		return server.Query403JSONResponse(*clientError(e)), nil
	}

	if conflict.Is(e) {
		return server.Query409JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.Query503JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.Query500JSONResponse(*s.answerWith(e, "")), nil
	}

	return server.Query200JSONResponse{
		Columns: rows.Columns,
		Rows:    rows.Rows,
	}, nil
}
