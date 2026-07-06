package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) ListNodes(
	_ context.Context,
	r server.ListNodesRequestObject,
) (server.ListNodesResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.ListNodes400JSONResponse(*clientError(e)), nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.ListNodes500JSONResponse(*s.captureDetail(e)), nil
	}

	nodes, e := s.service.ListNodes(c)

	if e != nil {
		return server.ListNodes500JSONResponse(*s.captureDetail(e)), nil
	}

	pointers := convert.Nodes(nodes)
	result := make(server.ListNodes200JSONResponse, len(pointers))

	for i, n := range pointers {
		result[i] = *n
	}

	return result, nil
}
