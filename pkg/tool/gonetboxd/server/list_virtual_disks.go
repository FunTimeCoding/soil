package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListVirtualDisks(
	_ context.Context,
	r server.ListVirtualDisksRequestObject,
) (server.ListVirtualDisksResponseObject, error) {
	disks, e := s.service.VirtualDisks(r.Name)

	if not_found.Is(e) {
		return server.ListVirtualDisks404JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListVirtualDisks500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListVirtualDisks200JSONResponse(convert.Disks(disks)), nil
}
