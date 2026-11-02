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

func (s *Server) ListNodes(
	_ context.Context,
	r server.ListNodesRequestObject,
) (server.ListNodesResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if not_found.Is(e) {
		return server.ListNodes404JSONResponse(*clientError(e)), nil
	}

	if not_selected.Is(e) {
		return server.ListNodes400JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListNodes500JSONResponse(*s.answer(e)), nil
	}

	nodes, e := s.service.ListNodes(instance)

	switch {
	case validation.Is(e):
		return server.ListNodes400JSONResponse(*clientError(e)), nil
	case not_found.Is(e):
		return server.ListNodes404JSONResponse(*clientError(e)), nil
	case forbidden.Is(e):
		return server.ListNodes403JSONResponse(*clientError(e)), nil
	case unavailable.Is(e):
		return server.ListNodes503JSONResponse(*clientError(e)), nil
	case e != nil:
		return server.ListNodes500JSONResponse(*s.answer(e)), nil
	}

	pointers := convert.Nodes(nodes)
	result := make(server.ListNodes200JSONResponse, len(pointers))

	for i, n := range pointers {
		result[i] = *n
	}

	return result, nil
}
