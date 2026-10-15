package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateConfigurationContext(
	_ context.Context,
	r server.CreateConfigurationContextRequestObject,
) (server.CreateConfigurationContextResponseObject, error) {
	weight := constant.ConfigurationContextWeight

	if r.Body.Weight != nil {
		weight = *r.Body.Weight
	}

	result, e := s.client.CreateConfigurationContext(
		r.Body.Name,
		weight,
		r.Body.Payload,
		convert.Assignment(r.Body.Assignment),
	)

	if not_found.Is(e) {
		return server.CreateConfigurationContext404JSONResponse(
			*clientError(e),
		), nil
	}

	if e != nil {
		return server.CreateConfigurationContext500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.CreateConfigurationContext201JSONResponse(
		*convert.ConfigurationContext(result),
	), nil
}
