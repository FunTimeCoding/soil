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

func (s *Server) ListMachineSnapshots(
	x context.Context,
	r server.ListMachineSnapshotsRequestObject,
) (server.ListMachineSnapshotsResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if not_found.Is(e) {
		return server.ListMachineSnapshots404JSONResponse(*clientError(e)), nil
	}

	if not_selected.Is(e) {
		return server.ListMachineSnapshots400JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListMachineSnapshots500JSONResponse(*s.answer(x, e)), nil
	}

	node := ""

	if r.Params.Node != nil {
		node = *r.Params.Node
	}

	snapshots, e := s.service.ListMachineSnapshots(
		instance,
		int(r.Identifier),
		node,
	)

	switch {
	case validation.Is(e):
		return server.ListMachineSnapshots400JSONResponse(*clientError(e)), nil
	case not_found.Is(e):
		return server.ListMachineSnapshots404JSONResponse(*clientError(e)), nil
	case forbidden.Is(e):
		return server.ListMachineSnapshots403JSONResponse(*clientError(e)), nil
	case unavailable.Is(e):
		return server.ListMachineSnapshots503JSONResponse(*clientError(e)), nil
	case e != nil:
		return server.ListMachineSnapshots500JSONResponse(*s.answer(x, e)), nil
	}

	converted := convert.MachineSnapshots(snapshots)
	result := make(server.ListMachineSnapshots200JSONResponse, len(converted))

	for i, v := range converted {
		result[i] = *v
	}

	return result, nil
}
