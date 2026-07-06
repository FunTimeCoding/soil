package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateTag(
	_ context.Context,
	r server.CreateTagRequestObject,
) (server.CreateTagResponseObject, error) {
	t, e := s.client.CreateTag(r.Body.Name)

	if e != nil {
		return server.CreateTag500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.CreateTag201JSONResponse{
		Identifier: t.Identifier,
		Name:       t.Name,
	}, nil
}
