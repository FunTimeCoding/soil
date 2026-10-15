package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) GetConfigurationContext(
	_ context.Context,
	r server.GetConfigurationContextRequestObject,
) (server.GetConfigurationContextResponseObject, error) {
	result, e := s.client.ConfigurationContextByName(r.Name)

	if not_found.Is(e) {
		return server.GetConfigurationContext404JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.GetConfigurationContext500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.GetConfigurationContext200JSONResponse(
		*convert.ConfigurationContext(result),
	), nil
}
