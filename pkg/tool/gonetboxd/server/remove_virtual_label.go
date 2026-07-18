package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) RemoveVirtualLabel(
	_ context.Context,
	r server.RemoveVirtualLabelRequestObject,
) (server.RemoveVirtualLabelResponseObject, error) {
	m, e := s.client.VirtualMachineByName(r.Name)

	if e != nil {
		return server.RemoveVirtualLabel500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	if e := s.store.RemoveLabel(
		constant.VirtualMachineAddress,
		m.Identifier,
		r.Key,
	); e != nil {
		return server.RemoveVirtualLabel500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.RemoveVirtualLabel204Response{}, nil
}
