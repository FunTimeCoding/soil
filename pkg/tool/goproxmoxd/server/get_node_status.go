package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) GetNodeStatus(
	_ context.Context,
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
		return server.GetNodeStatus500JSONResponse(*s.captureFail(e, e.Error())), nil
	}

	result, e := s.service.GetNodeStatus(instance, r.Name)

	if e != nil {
		return server.GetNodeStatus500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.GetNodeStatus200JSONResponse(*convert.NodeStatus(result)), nil
}
