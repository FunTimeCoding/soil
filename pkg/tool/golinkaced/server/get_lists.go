package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
)

func (s *Server) GetLists(
	_ context.Context,
	r server.GetListsRequestObject,
) (server.GetListsResponseObject, error) {
	page := 1

	if r.Params.Page != nil {
		page = int(*r.Params.Page)
	}

	v, e := s.service.ListsPage(page)

	if forbidden.Is(e) {
		return server.GetLists403JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.GetLists503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetLists500JSONResponse(*s.answer(e)), nil
	}

	result := make([]server.List, len(v.Items))

	for i, l := range v.Items {
		result[i] = server.List{
			Identifier: int32(l.Identifier),
			Name:       l.Name,
		}
	}

	total := v.Total

	return server.GetLists200JSONResponse(
		server.ListPage{Total: &total, Lists: &result},
	), nil
}
