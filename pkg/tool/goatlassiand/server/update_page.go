package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) UpdatePage(
	_ context.Context,
	r server.UpdatePageRequestObject,
) (server.UpdatePageResponseObject, error) {
	message := ""

	if r.Body.Message != nil {
		message = *r.Body.Message
	}

	result, e := s.confluence.UpdatePage(
		r.Identifier,
		r.Body.Title,
		r.Body.Body,
		message,
	)

	if e != nil {
		return server.UpdatePage500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.UpdatePage200JSONResponse(
		*convert.ConfluencePageDetail(result),
	), nil
}
