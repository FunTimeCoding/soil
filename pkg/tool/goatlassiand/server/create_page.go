package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) CreatePage(
	_ context.Context,
	r server.CreatePageRequestObject,
) (server.CreatePageResponseObject, error) {
	result, e := s.confluence.CreatePage(
		r.Body.SpaceIdentifier,
		r.Body.ParentIdentifier,
		r.Body.Title,
		r.Body.Body,
	)

	if e != nil {
		return server.CreatePage500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.CreatePage201JSONResponse(
		*convert.ConfluencePageDetail(result),
	), nil
}
