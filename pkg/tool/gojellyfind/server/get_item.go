package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/server"
)

func (s *Server) GetItem(
	_ context.Context,
	r server.GetItemRequestObject,
) (server.GetItemResponseObject, error) {
	v, e := s.client.Item(r.Identifier)

	if e != nil {
		return server.GetItem500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.GetItem200JSONResponse(*convertItem(v)), nil
}
