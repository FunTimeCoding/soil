package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/server"
)

func (s *Server) ListIndexes(
	c context.Context,
	r server.ListIndexesRequestObject,
) (server.ListIndexesResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if not_found.Is(e) {
		return server.ListIndexes404JSONResponse(*clientError(e)), nil
	}

	if not_selected.Is(e) {
		return server.ListIndexes400JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListIndexes500JSONResponse(*s.captureFail(e, e.Error())), nil
	}

	schema := "public"

	if r.Params.Schema != nil {
		schema = *r.Params.Schema
	}

	rows, e := s.service.ListIndexes(c, instance, schema, r.Params.Table)

	if e != nil {
		return server.ListIndexes500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.ListIndexes200JSONResponse{Rows: toRows(rows)}, nil
}
