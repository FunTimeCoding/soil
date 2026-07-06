package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListTags(
	_ context.Context,
	_ server.ListTagsRequestObject,
) (server.ListTagsResponseObject, error) {
	tags, e := s.client.Tags()

	if e != nil {
		return server.ListTags500JSONResponse(*s.captureDetail(e)), nil
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
