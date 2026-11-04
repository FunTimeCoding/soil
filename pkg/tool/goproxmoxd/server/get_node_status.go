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

func (s *Server) GetNodeStatus(
	x context.Context,
	r server.GetNodeStatusRequestObject,
) (server.GetNodeStatusResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if not_found.Is(e) {
		return server.GetNodeStatus404JSONResponse(*clientError(e)), nil
	}

	if not_selected.Is(e) {
		return server.GetNodeStatus400JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.GetNodeStatus500JSONResponse(*s.answer(x, e)), nil
	}

	result, e := s.service.GetNodeStatus(instance, r.Name)

	switch {
	case validation.Is(e):
		return server.GetNodeStatus400JSONResponse(*clientError(e)), nil
	case not_found.Is(e):
		return server.GetNodeStatus404JSONResponse(*clientError(e)), nil
	case forbidden.Is(e):
		return server.GetNodeStatus403JSONResponse(*clientError(e)), nil
	case unavailable.Is(e):
		return server.GetNodeStatus503JSONResponse(*clientError(e)), nil
	case e != nil:
		return server.GetNodeStatus500JSONResponse(*s.answer(x, e)), nil
	}

	return server.GetNodeStatus200JSONResponse(*convert.NodeStatus(result)), nil
}
