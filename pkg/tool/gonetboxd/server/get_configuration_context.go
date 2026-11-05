package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) GetConfigurationContext(
	x context.Context,
	r server.GetConfigurationContextRequestObject,
) (server.GetConfigurationContextResponseObject, error) {
	result, e := s.service.ConfigurationContextByName(r.Name)

	if not_found.Is(e) {
		return server.GetConfigurationContext404JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.GetConfigurationContext503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetConfigurationContext500JSONResponse(
				*s.captureDetail(x, e),
			),
			nil
	}

	return server.GetConfigurationContext200JSONResponse(
			*convert.ConfigurationContext(result),
		),
		nil
}
