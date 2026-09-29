package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
)

func (s *Server) Search(
	_ context.Context,
	r server.SearchRequestObject,
) (server.SearchResponseObject, error) {
	entityType := ""

	if r.Params.Type != nil {
		entityType = *r.Params.Type
	}

	v, e := s.service.Search(r.Params.Query, entityType)

	if e != nil {
		return server.Search500JSONResponse(*s.captureDetail(e)), nil
	}

	result := server.SearchResult{}

	if v.Links != nil {
		links := make([]server.Link, len(v.Links))

		for i, l := range v.Links {
			links[i] = server.Link{
				Identifier: int32(l.Identifier),
				Name:       l.Title,
				Link:       l.Link,
			}
		}

		result.Links = &links
	}

	if v.Lists != nil {
		lists := make([]server.List, len(v.Lists))

		for i, l := range v.Lists {
			lists[i] = server.List{
				Identifier: int32(l.Identifier),
				Name:       l.Name,
			}
		}

		result.Lists = &lists
	}

	if v.Tags != nil {
		tags := make([]server.Tag, len(v.Tags))

		for i, t := range v.Tags {
			tags[i] = server.Tag{
				Identifier: int32(t.Identifier),
				Name:       t.Name,
			}
		}

		result.Tags = &tags
	}

	return server.Search200JSONResponse(result), nil
}
