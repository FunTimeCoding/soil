package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/server"
)

func (s *Server) GetItem(
	_ context.Context,
	r server.GetItemRequestObject,
) (server.GetItemResponseObject, error) {
	v, e := s.service.Item(r.Identifier)

	if not_found.Is(e) {
		return server.GetItem404JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetItem500JSONResponse(*s.answer(e)), nil
	}

	return server.GetItem200JSONResponse(*convertItem(v)), nil
}
