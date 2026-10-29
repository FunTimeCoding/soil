package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/server"
)

func (s *Server) SaveView(
	_ context.Context,
	r server.SaveViewRequestObject,
) (server.SaveViewResponseObject, error) {
	var path string

	if r.Body.FilePath != nil {
		path = *r.Body.FilePath
	}

	e := s.service.SaveView(r.Id, path)

	if validation.Is(e) {
		return server.SaveView400JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.SaveView404JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.SaveView500JSONResponse(*s.answer(e)), nil
	}

	return server.SaveView204Response{}, nil
}
