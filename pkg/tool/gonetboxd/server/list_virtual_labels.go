package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListVirtualLabels(
	_ context.Context,
	r server.ListVirtualLabelsRequestObject,
) (server.ListVirtualLabelsResponseObject, error) {
	labels, e := s.service.VirtualLabels(r.Name)

	if not_found.Is(e) {
		return server.ListVirtualLabels404JSONResponse(*clientError(e)), nil
	}

	if forbidden.Is(e) {
		return server.ListVirtualLabels403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListVirtualLabels500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListVirtualLabels200JSONResponse(convert.Labels(labels)), nil
}
