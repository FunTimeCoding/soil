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

func (s *Server) ListContainerSnapshots(
	_ context.Context,
	r server.ListContainerSnapshotsRequestObject,
) (server.ListContainerSnapshotsResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if not_found.Is(e) {
		return server.ListContainerSnapshots404JSONResponse(*clientError(e)), nil
	}

	if not_selected.Is(e) {
		return server.ListContainerSnapshots400JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListContainerSnapshots500JSONResponse(*s.answer(e)), nil
	}

	node := ""

	if r.Params.Node != nil {
		node = *r.Params.Node
	}

	snapshots, e := s.service.ListContainerSnapshots(
		instance,
		int(r.Identifier),
		node,
	)

	switch {
	case validation.Is(e):
		return server.ListContainerSnapshots400JSONResponse(*clientError(e)), nil
	case not_found.Is(e):
		return server.ListContainerSnapshots404JSONResponse(*clientError(e)), nil
	case forbidden.Is(e):
		return server.ListContainerSnapshots403JSONResponse(*clientError(e)), nil
	case unavailable.Is(e):
		return server.ListContainerSnapshots503JSONResponse(*clientError(e)), nil
	case e != nil:
		return server.ListContainerSnapshots500JSONResponse(*s.answer(e)), nil
	}

	converted := convert.ContainerSnapshots(snapshots)
	result := make(server.ListContainerSnapshots200JSONResponse, len(converted))

	for i, v := range converted {
		result[i] = *v
	}

	return result, nil
}
