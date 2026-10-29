package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/types/context_change"
)

func (s *Server) UpdateConfigurationContext(
	_ context.Context,
	r server.UpdateConfigurationContextRequestObject,
) (server.UpdateConfigurationContextResponseObject, error) {
	c := context_change.New()
	c.Weight = r.Body.Weight

	if r.Body.Payload != nil {
		c.Payload = *r.Body.Payload
	}

	if r.Body.Assignment != nil {
		c.Assignment = convert.Assignment(r.Body.Assignment)
	}

	result, e := s.service.UpdateConfigurationContext(r.Name, c)

	if not_found.Is(e) {
		return server.UpdateConfigurationContext404JSONResponse(*clientError(e)),
			nil
	}

	if validation.Is(e) {
		return server.UpdateConfigurationContext400JSONResponse(*clientError(e)), nil
	}

	if forbidden.Is(e) {
		return server.UpdateConfigurationContext403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.UpdateConfigurationContext500JSONResponse(
				*s.captureDetail(e),
			),
			nil
	}

	return server.UpdateConfigurationContext200JSONResponse(
			*convert.ConfigurationContext(result),
		),
		nil
}
