package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListConfigurationContexts(
	_ context.Context,
	_ server.ListConfigurationContextsRequestObject,
) (server.ListConfigurationContextsResponseObject, error) {
	result, e := s.service.ConfigurationContexts()

	if forbidden.Is(e) {
		return server.ListConfigurationContexts403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListConfigurationContexts500JSONResponse(
				*s.captureDetail(e),
			),
			nil
	}

	return server.ListConfigurationContexts200JSONResponse(
			convert.ConfigurationContexts(result),
		),
		nil
}
