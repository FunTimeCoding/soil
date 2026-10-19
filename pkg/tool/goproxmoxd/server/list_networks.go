package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) ListNetworks(
	_ context.Context,
	r server.ListNetworksRequestObject,
) (server.ListNetworksResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if not_found.Is(e) {
		return server.ListNetworks404JSONResponse(*clientError(e)), nil
	}

	if not_selected.Is(e) {
		return server.ListNetworks400JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListNetworks500JSONResponse(*s.captureFail(e, e.Error())), nil
	}

	networks, e := s.service.ListNetworks(instance, r.Name)

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
