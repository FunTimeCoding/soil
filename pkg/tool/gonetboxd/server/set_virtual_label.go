package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) SetVirtualLabel(
	_ context.Context,
	r server.SetVirtualLabelRequestObject,
) (server.SetVirtualLabelResponseObject, error) {
	result, e := s.service.SetVirtualLabel(r.Name, r.Key, r.Body.Value)

	if not_found.Is(e) {
		return server.SetVirtualLabel404JSONResponse(*clientError(e)), nil
	}

	if validation.Is(e) {
		return server.SetVirtualLabel400JSONResponse(*clientError(e)), nil
	}

	if forbidden.Is(e) {
		return server.SetVirtualLabel403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.SetVirtualLabel500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.SetVirtualLabel200JSONResponse(*convert.Label(result)), nil
}
