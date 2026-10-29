package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
)

func (s *Server) GetTags(
	_ context.Context,
	r server.GetTagsRequestObject,
) (server.GetTagsResponseObject, error) {
	page := 1

	if r.Params.Page != nil {
		page = int(*r.Params.Page)
	}

	v, e := s.service.TagsPage(page)

	if forbidden.Is(e) {
		return server.GetTags403JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.GetTags503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetTags500JSONResponse(*s.answer(e)), nil
	}

	result := make([]server.Tag, len(v.Items))

	for i, t := range v.Items {
		result[i] = server.Tag{
			Identifier: int32(t.Identifier),
			Name:       t.Name,
		}
	}

	total := v.Total

	return server.GetTags200JSONResponse(
		server.TagPage{Total: &total, Tags: &result},
	), nil
}
