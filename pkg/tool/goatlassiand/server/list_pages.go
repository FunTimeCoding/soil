package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) ListPages(
	_ context.Context,
	r server.ListPagesRequestObject,
) (server.ListPagesResponseObject, error) {
	status := ""

	if r.Params.Status != nil {
		status = *r.Params.Status
	}

	result, e := s.confluence.PagesBySpace(r.Params.Space, status)

	if e != nil {
		return server.ListPages500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListPages200JSONResponse(
		convert.ConfluencePagesFromPages(result),
	), nil
}
