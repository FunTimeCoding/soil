package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/generated/server"
)

func (s *Server) ListTables(
	c context.Context,
	r server.ListTablesRequestObject,
) (server.ListTablesResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.ListTables400JSONResponse(*clientError(e)), nil
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
