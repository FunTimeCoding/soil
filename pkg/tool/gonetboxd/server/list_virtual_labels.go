package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListVirtualLabels(
	_ context.Context,
	r server.ListVirtualLabelsRequestObject,
) (server.ListVirtualLabelsResponseObject, error) {
	m, e := s.client.VirtualMachineByName(r.Name)

	if e != nil {
		return server.ListVirtualLabels500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	labels, e := s.store.Labels(
		constant.VirtualMachineAddress,
		m.Identifier,
	)

	if e != nil {
		return server.ListVirtualLabels500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.ListVirtualLabels200JSONResponse(
		convert.Labels(labels),
	), nil
}
