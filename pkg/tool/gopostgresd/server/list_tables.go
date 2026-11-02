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

func (s *Server) ListTables(
	c context.Context,
	r server.ListTablesRequestObject,
) (server.ListTablesResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if not_found.Is(e) {
		return server.ListTables404JSONResponse(*clientError(e)), nil
	}

	if not_selected.Is(e) {
		return server.ListTables400JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListTables500JSONResponse(*s.answerWith(e, "")), nil
	}

	var schema string

	if r.Params.Schema != nil {
		schema = *r.Params.Schema
	}

	rows, e := s.service.ListTables(c, instance, schema)

	if validation.Is(e) {
		return server.ListTables400JSONResponse(*clientError(e)), nil
	}

	if forbidden.Is(e) {
		return server.ListTables403JSONResponse(*clientError(e)), nil
	}

	if conflict.Is(e) {
		return server.ListTables409JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.ListTables503JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListTables500JSONResponse(*s.answerWith(e, "")), nil
	}

	return server.ListTables200JSONResponse{
		Columns: rows.Columns,
		Rows:    rows.Rows,
	}, nil
}
