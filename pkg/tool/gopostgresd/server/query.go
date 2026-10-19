package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
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
		return server.Query500JSONResponse(*s.captureFail(e, e.Error())), nil
	}

	rows, e := s.service.Query(c, instance, r.Body.Sql)

	if e != nil {
		return server.Query500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.Query200JSONResponse{Rows: toRows(rows)}, nil
}
