package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/server"
)

func (s *Server) CreateView(
	_ context.Context,
	r server.CreateViewRequestObject,
) (server.CreateViewResponseObject, error) {
	var syntax string

	if r.Body.Syntax != nil {
		syntax = *r.Body.Syntax
	}

	v, e := s.client.CreateView(r.Body.Title, r.Body.Content, syntax)
	errors.PanicOnError(e)

	return server.CreateView200JSONResponse(*convertView(v)), nil
}
