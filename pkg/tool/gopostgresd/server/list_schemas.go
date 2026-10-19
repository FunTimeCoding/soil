package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
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
		return server.ListSchemas500JSONResponse(*s.captureFail(e, e.Error())), nil
	}

	rows, e := s.service.ListSchemas(c, instance)

	if e != nil {
		return server.ListSchemas500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.ListSchemas200JSONResponse{Rows: toRows(rows)}, nil
}
