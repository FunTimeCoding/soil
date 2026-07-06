package server

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) ListMachineSnapshots(
	_ context.Context,
	r server.ListMachineSnapshotsRequestObject,
) (server.ListMachineSnapshotsResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.ListMachineSnapshots400JSONResponse(*clientError(e)), nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.ListMachineSnapshots500JSONResponse(*s.captureDetail(e)), nil
	}

	node := ""

	if r.Params.Node != nil {
		node = *r.Params.Node
	}

	snapshots, e := s.service.ListMachineSnapshots(c, int(r.Identifier), node)

	if e != nil {
		if errors.Is(e, not_found.Sentinel) {
			return server.ListMachineSnapshots404JSONResponse{Error: e.Error()}, nil
		}

		return server.ListMachineSnapshots500JSONResponse(*s.captureDetail(e)), nil
	}

	converted := convert.MachineSnapshots(snapshots)
	result := make(server.ListMachineSnapshots200JSONResponse, len(converted))

	for i, v := range converted {
		result[i] = *v
	}

	return result, nil
}
