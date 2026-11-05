package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/server"
)

func (s *Server) CreateView(
	x context.Context,
	r server.CreateViewRequestObject,
) (server.CreateViewResponseObject, error) {
	var syntax string

	if r.Body.Syntax != nil {
		syntax = *r.Body.Syntax
	}

	v, e := s.service.CreateView(r.Body.Title, r.Body.Content, syntax)

	if unavailable.Is(e) {
		return server.CreateView503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.CreateView500JSONResponse(*s.answer(x, e)), nil
	}

	return server.CreateView200JSONResponse(*convertView(v)), nil
}
