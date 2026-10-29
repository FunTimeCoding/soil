package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateVirtualDisk(
	_ context.Context,
	r server.CreateVirtualDiskRequestObject,
) (server.CreateVirtualDiskResponseObject, error) {
	result, e := s.service.CreateVirtualDisk(r.Name, r.Body.Name, r.Body.Size)

	if not_found.Is(e) {
		return server.CreateVirtualDisk404JSONResponse(*clientError(e)), nil
	}

	if validation.Is(e) {
		return server.CreateVirtualDisk400JSONResponse(*clientError(e)), nil
	}

	if forbidden.Is(e) {
		return server.CreateVirtualDisk403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.CreateVirtualDisk500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.CreateVirtualDisk201JSONResponse(*convert.Disk(result)), nil
}
