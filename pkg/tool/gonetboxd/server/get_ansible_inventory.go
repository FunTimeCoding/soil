package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) GetAnsibleInventory(
	x context.Context,
	r server.GetAnsibleInventoryRequestObject,
) (server.GetAnsibleInventoryResponseObject, error) {
	field := constant.AnsibleField

	if r.Params.Field != nil {
		field = *r.Params.Field
	}

	result, e := s.service.AnsibleInventory(field)

	if unavailable.Is(e) {
		return server.GetAnsibleInventory503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetAnsibleInventory500JSONResponse(*s.captureDetail(x, e)), nil
	}

	return server.GetAnsibleInventory200TextResponse(result.Markup()), nil
}
