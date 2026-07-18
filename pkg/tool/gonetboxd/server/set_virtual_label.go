package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) SetVirtualLabel(
	_ context.Context,
	r server.SetVirtualLabelRequestObject,
) (server.SetVirtualLabelResponseObject, error) {
	m, e := s.client.VirtualMachineByName(r.Name)

	if e != nil {
		return server.SetVirtualLabel500JSONResponse(*s.captureDetail(e)), nil
	}

	result, e := s.store.SetLabel(
		constant.VirtualMachineAddress,
		m.Identifier,
		r.Key,
		r.Body.Value,
	)

	if e != nil {
		return server.SetVirtualLabel500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.SetVirtualLabel200JSONResponse(*convert.Label(result)), nil
}
