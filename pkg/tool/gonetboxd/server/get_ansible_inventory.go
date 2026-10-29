package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) GetAnsibleInventory(
	_ context.Context,
	r server.GetAnsibleInventoryRequestObject,
) (server.GetAnsibleInventoryResponseObject, error) {
	field := constant.AnsibleField

	if r.Params.Field != nil {
		field = *r.Params.Field
	}

	result, e := s.service.AnsibleInventory(field)

	if forbidden.Is(e) {
		return server.GetAnsibleInventory403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.GetAnsibleInventory500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.GetAnsibleInventory200TextResponse(result.Markup()), nil
}
