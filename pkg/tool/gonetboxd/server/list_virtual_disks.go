package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListVirtualDisks(
	_ context.Context,
	r server.ListVirtualDisksRequestObject,
) (server.ListVirtualDisksResponseObject, error) {
	m, e := s.client.VirtualMachineByName(r.Name)

	if e != nil {
		return server.ListVirtualDisks500JSONResponse(*s.captureDetail(e)), nil
	}

	disks, f := s.client.VirtualMachineDisks(m)

	if f != nil {
		return server.ListVirtualDisks500JSONResponse(*s.captureDetail(f)), nil
	}

	return server.ListVirtualDisks200JSONResponse(convert.Disks(disks)), nil
}
