package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
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

	result := make([]*server.Tag, 0, len(tags))

	for _, t := range tags {
		result = append(
			result,
			&server.Tag{Identifier: t.Identifier, Name: t.Name},
		)
	}

	return server.ListTags200JSONResponse(result), nil
}
