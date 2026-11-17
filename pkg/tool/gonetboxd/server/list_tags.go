package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListTags(
	x context.Context,
	_ server.ListTagsRequestObject,
) (server.ListTagsResponseObject, error) {
	tags, e := s.service.Tags()

	if unavailable.Is(e) {
		return server.ListTags503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListTags500JSONResponse(*s.captureDetail(x, e)), nil
	}

	return server.ListTags200JSONResponse(convert.Tags(tags)), nil
}
