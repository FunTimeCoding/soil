package server

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) ListContainerSnapshots(
	_ context.Context,
	r server.ListContainerSnapshotsRequestObject,
) (server.ListContainerSnapshotsResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.ListContainerSnapshots400JSONResponse(*clientError(e)), nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.ListContainerSnapshots500JSONResponse(*s.captureDetail(e)), nil
	}

	node := ""

	if r.Params.Node != nil {
		node = *r.Params.Node
	}

	snapshots, e := s.service.ListContainerSnapshots(
		c,
		int(r.Identifier),
		node,
	)

	if e != nil {
		if errors.Is(e, not_found.Sentinel) {
			return server.ListContainerSnapshots404JSONResponse{Error: e.Error()}, nil
		}

		return server.ListContainerSnapshots500JSONResponse(*s.captureDetail(e)), nil
	}

	converted := convert.ContainerSnapshots(snapshots)
	result := make(
		server.ListContainerSnapshots200JSONResponse,
		len(converted),
	)

	for i, v := range converted {
		result[i] = *v
	}

	return result, nil
}
