package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) ListNetworks(
	_ context.Context,
	r server.ListNetworksRequestObject,
) (server.ListNetworksResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.ListNetworks400JSONResponse(*clientError(e)), nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.ListNetworks500JSONResponse(*s.captureDetail(e)), nil
	}

	networks, e := s.service.ListNetworks(c, r.Name)

	if e != nil {
		return server.ListNetworks500JSONResponse(*s.captureDetail(e)), nil
	}

	pointers := convert.Networks(networks)
	result := make(server.ListNetworks200JSONResponse, len(pointers))

	for i, n := range pointers {
		result[i] = *n
	}

	return result, nil
}
