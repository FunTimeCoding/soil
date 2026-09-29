package server

import (
	"context"
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

	v, e := s.client.ListsPage(page)

	if e != nil {
		return server.GetLists500JSONResponse(*s.captureDetail(e)), nil
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
