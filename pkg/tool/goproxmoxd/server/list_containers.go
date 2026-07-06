package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) ListContainers(
	_ context.Context,
	r server.ListContainersRequestObject,
) (server.ListContainersResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.ListContainers400JSONResponse(*clientError(e)), nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.ListContainers500JSONResponse(*s.captureDetail(e)), nil
	}

	node := ""

	if r.Params.Node != nil {
		node = *r.Params.Node
	}

	containers, e := s.service.ListContainers(c, node)

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
