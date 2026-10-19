package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
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
		return server.ListTables500JSONResponse(*s.captureFail(e, e.Error())), nil
	}

	schema := "public"

	if r.Params.Schema != nil {
		schema = *r.Params.Schema
	}

	rows, e := s.service.ListTables(c, instance, schema)

	if e != nil {
		return server.ListTables500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.ListTables200JSONResponse{Rows: toRows(rows)}, nil
}
