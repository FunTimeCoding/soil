package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) ListMachines(
	_ context.Context,
	r server.ListMachinesRequestObject,
) (server.ListMachinesResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.ListMachines400JSONResponse(*clientError(e)), nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.ListMachines500JSONResponse(*s.captureDetail(e)), nil
	}

	node := ""

	if r.Params.Node != nil {
		node = *r.Params.Node
	}

	machines, e := s.service.ListMachines(c, node)

	if e != nil {
		return server.ListMachines500JSONResponse(*s.captureDetail(e)), nil
	}

	pointers := convert.Machines(machines)
	result := make(server.ListMachines200JSONResponse, len(pointers))

	for i, m := range pointers {
		result[i] = *m
	}

	return result, nil
}
