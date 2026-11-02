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
		return server.GetContainer500JSONResponse(*s.answer(e)), nil
	}

	node := ""

	if r.Params.Node != nil {
		node = *r.Params.Node
	}

	ct, e := s.service.GetContainer(instance, int(r.Identifier), node)

	switch {
	case validation.Is(e):
		return server.GetContainer400JSONResponse(*clientError(e)), nil
	case not_found.Is(e):
		return server.GetContainer404JSONResponse(*clientError(e)), nil
	case forbidden.Is(e):
		return server.GetContainer403JSONResponse(*clientError(e)), nil
	case unavailable.Is(e):
		return server.GetContainer503JSONResponse(*clientError(e)), nil
	case e != nil:
		return server.GetContainer500JSONResponse(*s.answer(e)), nil
	}

	return server.GetContainer200JSONResponse(*convert.ContainerDetail(ct)), nil
}
