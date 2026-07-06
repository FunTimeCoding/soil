package server

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) GetContainer(
	_ context.Context,
	r server.GetContainerRequestObject,
) (server.GetContainerResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.GetContainer400JSONResponse(*clientError(e)), nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.GetContainer500JSONResponse(*s.captureDetail(e)), nil
	}

	node := ""

	if r.Params.Node != nil {
		node = *r.Params.Node
	}

	ct, e := s.service.GetContainer(c, int(r.Identifier), node)

	if e != nil {
		if errors.Is(e, not_found.Sentinel) {
			return server.GetContainer404JSONResponse{Error: e.Error()}, nil
		}

		return server.GetContainer500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.GetContainer200JSONResponse(
		*convert.ContainerDetail(ct),
	), nil
}
