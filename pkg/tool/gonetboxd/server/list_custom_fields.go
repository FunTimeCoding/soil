package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListCustomFields(
	_ context.Context,
	_ server.ListCustomFieldsRequestObject,
) (server.ListCustomFieldsResponseObject, error) {
	result, e := s.service.CustomFields()

	if forbidden.Is(e) {
		return server.ListCustomFields403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListCustomFields500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListCustomFields200JSONResponse(convert.CustomFields(result)),
		nil
}
