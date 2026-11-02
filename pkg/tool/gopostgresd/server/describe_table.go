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

func (s *Server) DescribeTable(
	c context.Context,
	r server.DescribeTableRequestObject,
) (server.DescribeTableResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if not_found.Is(e) {
		return server.DescribeTable404JSONResponse(*clientError(e)), nil
	}

	if not_selected.Is(e) {
		return server.DescribeTable400JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.DescribeTable500JSONResponse(*s.answerWith(e, "")), nil
	}

	var schema string

	if r.Params.Schema != nil {
		schema = *r.Params.Schema
	}

	rows, e := s.service.DescribeTable(c, instance, schema, r.Table)

	if validation.Is(e) {
		return server.DescribeTable400JSONResponse(*clientError(e)), nil
	}

	if forbidden.Is(e) {
		return server.DescribeTable403JSONResponse(*clientError(e)), nil
	}

	if conflict.Is(e) {
		return server.DescribeTable409JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.DescribeTable503JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.DescribeTable500JSONResponse(*s.answerWith(e, "")), nil
	}

	return server.DescribeTable200JSONResponse{
		Columns: rows.Columns,
		Rows:    rows.Rows,
	}, nil
}
