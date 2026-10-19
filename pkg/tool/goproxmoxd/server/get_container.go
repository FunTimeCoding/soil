package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) GetContainer(
	_ context.Context,
	r server.GetContainerRequestObject,
) (server.GetContainerResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if not_found.Is(e) {
		return server.GetContainer404JSONResponse(*clientError(e)), nil
	}

	if not_selected.Is(e) {
		return server.GetContainer400JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.GetContainer500JSONResponse(*s.captureFail(e, e.Error())), nil
	}

	node := ""

	if r.Params.Node != nil {
		node = *r.Params.Node
	}

	ct, e := s.service.GetContainer(instance, int(r.Identifier), node)

	if e != nil {
		if not_found.Is(e) {
			return server.GetContainer404JSONResponse{Error: e.Error()}, nil
		}

		return server.GetContainer500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.GetContainer200JSONResponse(*convert.ContainerDetail(ct)), nil
}
