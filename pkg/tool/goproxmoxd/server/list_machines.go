package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) ListMachines(
	_ context.Context,
	r server.ListMachinesRequestObject,
) (server.ListMachinesResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if not_found.Is(e) {
		return server.ListMachines404JSONResponse(*clientError(e)), nil
	}

	if not_selected.Is(e) {
		return server.ListMachines400JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListMachines500JSONResponse(*s.captureFail(e, e.Error())), nil
	}

	node := ""

	if r.Params.Node != nil {
		node = *r.Params.Node
	}

	machines, e := s.service.ListMachines(instance, node)

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
