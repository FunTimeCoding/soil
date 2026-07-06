package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/server"
)

func (s *Server) GetTags(
	_ context.Context,
	_ server.GetTagsRequestObject,
) (server.GetTagsResponseObject, error) {
	tags := s.service.ListSourceTypes()
	result := make(server.GetTags200JSONResponse, len(tags))

	for i, t := range tags {
		result[i] = &server.SourceTypeTag{
			Collection: new(string(t.Collection)),
			PathPrefix: new(string(t.PathPrefix)),
			SourceType: new(string(t.SourceType)),
		}
	}

	return result, nil
}
