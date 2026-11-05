package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListConfigurationContexts(
	x context.Context,
	_ server.ListConfigurationContextsRequestObject,
) (server.ListConfigurationContextsResponseObject, error) {
	result, e := s.service.ConfigurationContexts()

	if unavailable.Is(e) {
		return server.ListConfigurationContexts503JSONResponse{Error: e.Error()},
			nil
	}

	if e != nil {
		return server.ListConfigurationContexts500JSONResponse(
				*s.captureDetail(x, e),
			),
			nil
	}

	return server.ListConfigurationContexts200JSONResponse(
			convert.ConfigurationContexts(result),
		),
		nil
}
