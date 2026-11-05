package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/server"
)

func (s *Server) GetItem(
	x context.Context,
	r server.GetItemRequestObject,
) (server.GetItemResponseObject, error) {
	v, e := s.service.Item(r.Identifier)

	if not_found.Is(e) {
		return server.GetItem404JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.GetItem503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetItem500JSONResponse(*s.answer(x, e)), nil
	}

	return server.GetItem200JSONResponse(*convertItem(v)), nil
}
