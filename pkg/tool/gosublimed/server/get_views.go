package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/server"
)

func (s *Server) GetViews(
	_ context.Context,
	_ server.GetViewsRequestObject,
) (server.GetViewsResponseObject, error) {
	views, e := s.client.Views()
	errors.PanicOnError(e)
	result := make(server.GetViews200JSONResponse, 0, len(views))

	for _, v := range views {
		result = append(result, *convertView(v))
	}

	return result, nil
}
