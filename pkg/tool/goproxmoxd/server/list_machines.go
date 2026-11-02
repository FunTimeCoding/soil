package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
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
		return server.ListMachines500JSONResponse(*s.answer(e)), nil
	}

	node := ""

	if r.Params.Node != nil {
		node = *r.Params.Node
	}

	machines, e := s.service.ListMachines(instance, node)

	switch {
	case validation.Is(e):
		return server.ListMachines400JSONResponse(*clientError(e)), nil
	case not_found.Is(e):
		return server.ListMachines404JSONResponse(*clientError(e)), nil
	case forbidden.Is(e):
		return server.ListMachines403JSONResponse(*clientError(e)), nil
	case unavailable.Is(e):
		return server.ListMachines503JSONResponse(*clientError(e)), nil
	case e != nil:
		return server.ListMachines500JSONResponse(*s.answer(e)), nil
	}

	pointers := convert.Machines(machines)
	result := make(server.ListMachines200JSONResponse, len(pointers))

	for i, m := range pointers {
		result[i] = *m
	}

	return result, nil
}
