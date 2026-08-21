package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListCables(
	_ context.Context,
	_ server.ListCablesRequestObject,
) (server.ListCablesResponseObject, error) {
	cables, e := s.client.Cables()

	if e != nil {
		return server.ListCables500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListCables200JSONResponse(convert.Cables(cables)), nil
}
