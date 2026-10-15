package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) DeleteConfigurationContext(
	_ context.Context,
	r server.DeleteConfigurationContextRequestObject,
) (server.DeleteConfigurationContextResponseObject, error) {
	e := s.client.DeleteConfigurationContext(r.Name)

	if not_found.Is(e) {
		return server.DeleteConfigurationContext404JSONResponse(
			*clientError(e),
		), nil
	}

	if e != nil {
		return server.DeleteConfigurationContext500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.DeleteConfigurationContext204Response{}, nil
}
