package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListCables(
	x context.Context,
	_ server.ListCablesRequestObject,
) (server.ListCablesResponseObject, error) {
	cables, e := s.service.Cables()

	if unavailable.Is(e) {
		return server.ListCables503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListCables500JSONResponse(*s.captureDetail(x, e)), nil
	}

	return server.ListCables200JSONResponse(convert.Cables(cables)), nil
}
