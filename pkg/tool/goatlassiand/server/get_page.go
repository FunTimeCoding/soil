package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) GetPage(
	_ context.Context,
	r server.GetPageRequestObject,
) (server.GetPageResponseObject, error) {
	result, e := s.confluence.Page(r.Identifier)

	if e != nil {
		return server.GetPage500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.GetPage200JSONResponse(
		*convert.ConfluencePageDetail(result),
	), nil
}
