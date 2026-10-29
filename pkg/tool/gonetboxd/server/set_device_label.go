package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) SetDeviceLabel(
	_ context.Context,
	r server.SetDeviceLabelRequestObject,
) (server.SetDeviceLabelResponseObject, error) {
	result, e := s.service.SetDeviceLabel(r.Name, r.Key, r.Body.Value)

	if not_found.Is(e) {
		return server.SetDeviceLabel404JSONResponse(*clientError(e)), nil
	}

	if validation.Is(e) {
		return server.SetDeviceLabel400JSONResponse(*clientError(e)), nil
	}

	if forbidden.Is(e) {
		return server.SetDeviceLabel403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.SetDeviceLabel500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.SetDeviceLabel200JSONResponse(*convert.Label(result)), nil
}
