package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) ListContainers(
	_ context.Context,
	r server.ListContainersRequestObject,
) (server.ListContainersResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if not_found.Is(e) {
		return server.ListContainers404JSONResponse(*clientError(e)), nil
	}

	if not_selected.Is(e) {
		return server.ListContainers400JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListContainers500JSONResponse(
			*s.captureFail(e, e.Error()),
		), nil
	}

	node := ""

	if r.Params.Node != nil {
		node = *r.Params.Node
	}

	containers, e := s.service.ListContainers(instance, node)

	if e != nil {
		return server.ListContainers500JSONResponse(*s.captureDetail(e)), nil
	}

	pointers := convert.Containers(containers)
	result := make(server.ListContainers200JSONResponse, len(pointers))

	for i, ct := range pointers {
		result[i] = *ct
	}

	return result, nil
}
