package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/convert"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/server"
)

func (s *Server) GetTags(
	_ context.Context,
	_ server.GetTagsRequestObject,
) (server.GetTagsResponseObject, error) {
	result, e := s.habitica.Tags()

	if e != nil {
		return server.GetTags500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.GetTags200JSONResponse(convert.Tags(result)), nil
}
