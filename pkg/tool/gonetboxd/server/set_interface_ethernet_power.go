package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) SetInterfaceEthernetPower(
	x context.Context,
	r server.SetInterfaceEthernetPowerRequestObject,
) (server.SetInterfaceEthernetPowerResponseObject, error) {
	mode := ""

	if r.Body.Mode != nil {
		mode = *r.Body.Mode
	}

	kind := ""

	if r.Body.Type != nil {
		kind = *r.Body.Type
	}

	i, e := s.service.SetInterfaceEthernetPower(r.Name, r.Interface, mode, kind)

	if not_found.Is(e) {
		return server.SetInterfaceEthernetPower404JSONResponse(
			*clientError(e),
		), nil
	}

	if validation.Is(e) {
		return server.SetInterfaceEthernetPower400JSONResponse(
			*clientError(e),
		), nil
	}

	if unavailable.Is(e) {
		return server.SetInterfaceEthernetPower503JSONResponse{
			Error: e.Error(),
		}, nil
	}

	if e != nil {
		return server.SetInterfaceEthernetPower500JSONResponse(
			*s.captureDetail(x, e),
		), nil
	}

	return server.SetInterfaceEthernetPower200JSONResponse(
		*convert.Interface(i),
	), nil
}
