package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListCustomFields(
	x context.Context,
	_ server.ListCustomFieldsRequestObject,
) (server.ListCustomFieldsResponseObject, error) {
	result, e := s.service.CustomFields()

	if unavailable.Is(e) {
		return server.ListCustomFields503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListCustomFields500JSONResponse(*s.captureDetail(x, e)), nil
	}

	return server.ListCustomFields200JSONResponse(convert.CustomFields(result)),
		nil
}
