package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/server"
)

func (s *Server) PostTag(
	_ context.Context,
	r server.PostTagRequestObject,
) (server.PostTagResponseObject, error) {
	collection := ""

	if r.Body.Collection != nil {
		collection = *r.Body.Collection
	}

	s.service.SetSourceType(collection, r.Body.Path, r.Body.SourceType)

	return server.PostTag200Response{}, nil
}
