package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/server"
)

func (s *Server) GetContext(
	_ context.Context,
	_ server.GetContextRequestObject,
) (server.GetContextResponseObject, error) {
	entries := s.service.ListContexts()
	result := make(server.GetContext200JSONResponse, len(entries))

	for i, e := range entries {
		result[i] = &server.ContextEntry{
			Collection:  e.Collection,
			PathPrefix:  e.PathPrefix,
			Description: e.Description,
		}
	}

	return result, nil
}
