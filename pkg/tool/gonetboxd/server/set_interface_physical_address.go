package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) SetInterfacePhysicalAddress(
	_ context.Context,
	r server.SetInterfacePhysicalAddressRequestObject,
) (server.SetInterfacePhysicalAddressResponseObject, error) {
	updated, e := s.service.SetInterfacePhysicalAddress(
		r.Name,
		r.Interface,
		r.Body.Address,
	)

	if e != nil {
		return server.SetInterfacePhysicalAddress500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.SetInterfacePhysicalAddress200JSONResponse(
		*convert.InterfacePhysicalAddress(updated),
	), nil
}
