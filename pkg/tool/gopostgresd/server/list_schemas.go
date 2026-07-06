package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/server"
)

func (s *Server) ListSchemas(
	c context.Context,
	r server.ListSchemasRequestObject,
) (server.ListSchemasResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.ListSchemas400JSONResponse(*clientError(e)), nil
	}

	rows, e := s.service.ListSchemas(c, instance)

	if e != nil {
		return server.ListSchemas500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.ListSchemas200JSONResponse{Rows: toRows(rows)}, nil
}
