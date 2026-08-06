package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateVirtualDisk(
	_ context.Context,
	r server.CreateVirtualDiskRequestObject,
) (server.CreateVirtualDiskResponseObject, error) {
	result, e := s.client.CreateVirtualDisk(
		r.Name,
		r.Body.Name,
		r.Body.Size,
	)

	if e != nil {
		return server.CreateVirtualDisk500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.CreateVirtualDisk201JSONResponse(*convert.Disk(result)), nil
}
