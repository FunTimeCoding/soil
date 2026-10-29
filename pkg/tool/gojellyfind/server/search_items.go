package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/server"
	"strings"
)

func (s *Server) SearchItems(
	_ context.Context,
	r server.SearchItemsRequestObject,
) (server.SearchItemsResponseObject, error) {
	var types []string

	if r.Params.Types != nil {
		for _, t := range strings.Split(*r.Params.Types, ",") {
			types = append(types, strings.TrimSpace(t))
		}
	}

	page := 0
	perPage := 0

	if r.Params.Page != nil {
		page = *r.Params.Page
	}

	if r.Params.PerPage != nil {
		perPage = *r.Params.PerPage
	}

	result, e := s.service.SearchItems(r.Params.Q, types, page, perPage)

	if e != nil {
		return server.SearchItems500JSONResponse(*s.answer(e)), nil
	}

	rows := make([]server.Item, 0, len(result.Items))

	for _, v := range result.Items {
		rows = append(rows, *convertItem(v))
	}

	return server.SearchItems200JSONResponse(
		server.SearchResult{
			Total: result.Total,
			Page:  result.Page,
			Items: rows,
		},
	), nil
}
