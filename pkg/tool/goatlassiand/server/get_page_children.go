package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) GetPageChildren(
	_ context.Context,
	r server.GetPageChildrenRequestObject,
) (server.GetPageChildrenResponseObject, error) {
	result, e := s.confluence.ChildPagesByIdentifier(r.Identifier)

	if e != nil {
		return server.GetPageChildren500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.GetPageChildren200JSONResponse(
		convert.ConfluencePagesFromPages(result),
	), nil
}
