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

func (s *Server) Explain(
	c context.Context,
	r server.ExplainRequestObject,
) (server.ExplainResponseObject, error) {
	instance, e := s.resolveInstance(r.Body.Instance)

	if not_found.Is(e) {
		return server.Explain404JSONResponse(*clientError(e)), nil
	}

	if not_selected.Is(e) {
		return server.Explain400JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.Explain500JSONResponse(*s.answerWith(e, "")), nil
	}

	rows, e := s.service.Explain(
		c,
		instance,
		r.Body.Sql,
		r.Body.Analyze != nil && *r.Body.Analyze,
	)

	if validation.Is(e) {
		return server.Explain400JSONResponse(*clientError(e)), nil
	}

	if forbidden.Is(e) {
		return server.Explain403JSONResponse(*clientError(e)), nil
	}

	if conflict.Is(e) {
		return server.Explain409JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.Explain503JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.Explain500JSONResponse(*s.answerWith(e, "")), nil
	}

	return server.Explain200JSONResponse{
		Columns: rows.Columns,
		Rows:    rows.Rows,
	}, nil
}
