package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/types/edit"
)

func (s *Server) EditView(
	_ context.Context,
	r server.EditViewRequestObject,
) (server.EditViewResponseObject, error) {
	v, e := s.service.EditView(
		edit.New(
			r.Id,
			r.Body.OldString,
			r.Body.NewString,
			r.Body.ReplaceAll != nil && *r.Body.ReplaceAll,
		),
	)

	if validation.Is(e) {
		return server.EditView400JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.EditView404JSONResponse{Error: e.Error()}, nil
	}

	if conflict.Is(e) {
		return server.EditView409JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.EditView500JSONResponse(*s.answer(e)), nil
	}

	return server.EditView200JSONResponse(*convertView(v)), nil
}
