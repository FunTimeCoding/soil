package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/server"
)

func (s *Server) GetViews(
	x context.Context,
	_ server.GetViewsRequestObject,
) (server.GetViewsResponseObject, error) {
	views, e := s.service.Views()

	if unavailable.Is(e) {
		return server.GetViews503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetViews500JSONResponse(*s.answer(x, e)), nil
	}

	result := make(server.GetViews200JSONResponse, 0, len(views))

	for _, v := range views {
		result = append(result, *convertView(v))
	}

	return result, nil
}
