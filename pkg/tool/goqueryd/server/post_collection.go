package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/server"
)

func (s *Server) PostCollection(
	_ context.Context,
	r server.PostCollectionRequestObject,
) (server.PostCollectionResponseObject, error) {
	pattern := ""

	if r.Body.Pattern != nil {
		pattern = *r.Body.Pattern
	}

	s.service.AddCollection(r.Body.Name, r.Body.Path, pattern)

	return server.PostCollection200Response{}, nil
}
