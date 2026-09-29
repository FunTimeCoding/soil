package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
)

func (s *Server) GetLinks(
	_ context.Context,
	r server.GetLinksRequestObject,
) (server.GetLinksResponseObject, error) {
	page := 1

	if r.Params.Page != nil {
		page = int(*r.Params.Page)
	}

	if r.Params.List != nil {
		links, e := s.client.LinksByList(int(*r.Params.List))

		if e != nil {
			return server.GetLinks500JSONResponse(*s.captureDetail(e)), nil
		}

		result := make([]server.Link, len(links))

		for i, l := range links {
			result[i] = server.Link{
				Identifier: int32(l.Identifier),
				Name:       l.Title,
				Link:       l.Link,
			}
		}

		total := len(result)

		return server.GetLinks200JSONResponse(
			server.LinkPage{Total: &total, Links: &result},
		), nil
	}

	v, e := s.client.LinksPage(page)

	if e != nil {
		return server.GetLinks500JSONResponse(*s.captureDetail(e)), nil
	}

	result := make([]server.Link, len(v.Items))

	for i, l := range v.Items {
		result[i] = server.Link{
			Identifier: int32(l.Identifier),
			Name:       l.Title,
			Link:       l.Link,
		}
	}

	total := v.Total

	return server.GetLinks200JSONResponse(
		server.LinkPage{Total: &total, Links: &result},
	), nil
}
