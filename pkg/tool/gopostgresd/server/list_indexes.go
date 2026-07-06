package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/server"
)

func (s *Server) ListIndexes(
	c context.Context,
	r server.ListIndexesRequestObject,
) (server.ListIndexesResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.ListIndexes400JSONResponse(*clientError(e)), nil
	}

	schema := "public"

	if r.Params.Schema != nil {
		schema = *r.Params.Schema
	}

	rows, e := s.service.ListIndexes(
		c,
		instance,
		schema,
		r.Params.Table,
	)

	if e != nil {
		return server.ListIndexes500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.ListIndexes200JSONResponse{Rows: toRows(rows)}, nil
}
